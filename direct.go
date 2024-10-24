package fronted

import (
	"bytes"
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	tls "github.com/refraction-networking/utls"

	"github.com/getlantern/golog"
	"github.com/getlantern/idletiming"
	"github.com/getlantern/ops"
)

const (
	defaultMaxAllowedCachedAge = 24 * time.Hour
	defaultMaxCacheSize        = 1000
	defaultCacheSaveInterval   = 5 * time.Second
	maxTries                   = 6
)

var (
	log = golog.LoggerFor("fronted")
)

// direct is an implementation of http.RoundTripper
type direct struct {
	certPool            *x509.CertPool
	masquerades         sortedMasquerades
	maxAllowedCachedAge time.Duration
	maxCacheSize        int
	cacheSaveInterval   time.Duration
	cacheDirty          chan interface{}
	cacheClosed         chan interface{}
	closeCacheOnce      sync.Once
	defaultProviderID   string
	providers           map[string]*Provider
	clientHelloID       tls.ClientHelloID
}

func (d *direct) loadCandidates(initial map[string]*Provider) {
	log.Debugf("Loading candidates for %d providers", len(initial))
	defer log.Debug("Finished loading candidates")

	for key, p := range initial {
		arr := p.Masquerades
		size := len(arr)
		log.Debugf("Adding %d candidates for %v", size, key)

		// make a shuffled copy of arr
		// ('inside-out' Fisher-Yates)
		sh := make([]*Masquerade, size)
		for i := 0; i < size; i++ {
			j := rand.IntN(i + 1) // 0 <= j <= i
			sh[i] = sh[j]
			sh[j] = arr[i]
		}

		for _, c := range sh {
			log.Trace("Adding candidate")
			d.masquerades = append(d.masquerades, &masquerade{Masquerade: *c, ProviderID: key})
		}
	}
}

func (d *direct) providerFor(m MasqueradeInterface) *Provider {
	pid := m.getProviderID()
	if pid == "" {
		pid = d.defaultProviderID
	}
	return d.providers[pid]
}

// Vet vets the specified Masquerade, verifying certificate using the given CertPool.
// This is used in genconfig.
func Vet(m *Masquerade, pool *x509.CertPool, testURL string) bool {
	d := &direct{
		certPool:            pool,
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
	}
	masq := &masquerade{Masquerade: *m}
	conn, _, err := d.doDial(masq)
	if err != nil {
		return false
	}
	defer conn.Close()
	return masq.postCheck(conn, testURL)
}

// findWorkingMasquerades finds working masquerades by vetting them in batches and in
// parallel. Speed is of the essence here, as without working masquerades, users will
// be unable to fetch proxy configurations, particularly in the case of a first time
// user who does not have proxies cached on disk.
func (d *direct) findWorkingMasquerades() {
	// vet masquerades in batches
	const batchSize int = 25
	var successful atomic.Uint32

	// We loop through all of them until we have 4 successful ones.
	for i := 0; i < len(d.masquerades) && successful.Load() < 4; i += batchSize {
		d.vetBatch(i, batchSize, &successful)
	}
}

func (d *direct) vetBatch(start, batchSize int, successful *atomic.Uint32) {
	var wg sync.WaitGroup
	masqueradeSize := len(d.masquerades)
	for j := start; j < start+batchSize && j < masqueradeSize; j++ {
		wg.Add(1)
		go func(m MasqueradeInterface) {
			defer wg.Done()
			if d.vetMasquerade(m) {
				successful.Add(1)
			}
		}(d.masquerades[j])
	}
	wg.Wait()
}

func (d *direct) vetMasquerade(m MasqueradeInterface) bool {
	conn, masqueradeGood, err := d.dialMasquerade(m)
	if err != nil {
		log.Errorf("unexpected error vetting masquerades: %v", err)
		return false
	}
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	provider := d.providerFor(m)
	if provider == nil {
		log.Debugf("Skipping masquerade with disabled/unknown provider id '%s' not in %v",
			m.getProviderID(), d.providers)
		return false
	}
	if !masqueradeGood(m.postCheck(conn, provider.TestURL)) {
		log.Debugf("Unsuccessful vetting with POST request, discarding masquerade")
		return false
	}

	log.Debugf("Finished vetting one masquerade %v", m)
	return true
}

// Do continually retries a given request until it succeeds because some
// fronting providers will return a 403 for some domains.
func (d *direct) RoundTrip(req *http.Request) (*http.Response, error) {
	res, _, err := d.RoundTripHijack(req)
	return res, err
}

// Do continually retries a given request until it succeeds because some
// fronting providers will return a 403 for some domains.  Also return the
// underlying net.Conn established.
func (d *direct) RoundTripHijack(req *http.Request) (*http.Response, net.Conn, error) {
	op := ops.Begin("fronted_roundtrip")
	defer op.End()

	isIdempotent := req.Method != http.MethodPost && req.Method != http.MethodPatch
	op.Set("is_idempotent", isIdempotent)

	originHost := req.URL.Hostname()
	op.Set("origin_host", originHost)

	var body []byte
	var err error
	if isIdempotent && req.Body != nil {
		// store body in-memory to be able to replay it if necessary
		body, err = io.ReadAll(req.Body)
		if err != nil {
			err := fmt.Errorf("unable to read request body: %v", err)
			op.FailIf(err)
			return nil, nil, err
		}
	}

	getBody := func() io.ReadCloser {
		if req.Body == nil {
			return nil
		}

		if !isIdempotent {
			return req.Body
		}
		return io.NopCloser(bytes.NewReader(body))
	}

	tries := 1
	if isIdempotent {
		tries = maxTries
	}

	for i := 0; i < tries; i++ {
		if i > 0 {
			log.Debugf("Retrying domain-fronted request, pass %d", i)
		}

		conn, m, masqueradeGood, err := d.dialAll(req.Context())
		if err != nil {
			// unable to find good masquerade, fail
			op.FailIf(err)
			return nil, nil, err
		}
		provider := d.providerFor(m)
		if provider == nil {
			log.Debugf("Skipping masquerade with disabled/unknown provider '%s'", m.getProviderID())
			masqueradeGood(false)
			continue
		}
		frontedHost := provider.Lookup(originHost)
		if frontedHost == "" {
			// this error is not the masquerade's fault in particular
			// so it is returned as good.
			conn.Close()
			masqueradeGood(true)
			err := fmt.Errorf("no domain fronting mapping for '%s'. Please add it to provider_map.yaml or equivalent for %s",
				m.getProviderID(), originHost)
			op.FailIf(err)
			return nil, nil, err
		}
		log.Debugf("Translated origin %s -> %s for provider %s...", originHost, frontedHost, m.getProviderID())

		reqi, err := cloneRequestWith(req, frontedHost, getBody())
		if err != nil {
			return nil, nil, op.FailIf(log.Errorf("Failed to copy http request with origin translated to %v?: %v", frontedHost, err))
		}

		// don't clobber/confuse Connection header on Upgrade requests.
		disableKeepAlives := true
		if strings.EqualFold(reqi.Header.Get("Connection"), "upgrade") {
			disableKeepAlives = false
		}

		tr := frontedHTTPTransport(conn, disableKeepAlives)
		resp, err := tr.RoundTrip(reqi)
		if err != nil {
			log.Debugf("Could not complete request: %v", err)
			masqueradeGood(false)
			continue
		}

		err = provider.ValidateResponse(resp)
		if err != nil {
			log.Debugf("Could not complete request: %v", err)
			resp.Body.Close()
			masqueradeGood(false)
			continue
		}

		masqueradeGood(true)
		return resp, conn, nil
	}

	return nil, nil, op.FailIf(errors.New("could not complete request even with retries"))
}

// Dial dials out using all available masquerades until one succeeds.
func (d *direct) dialAll(ctx context.Context) (net.Conn, MasqueradeInterface, func(bool) bool, error) {
	conn, m, masqueradeGood, err := d.dialAllWith(ctx, d.masquerades)
	return conn, m, masqueradeGood, err
}

func (d *direct) dialAllWith(ctx context.Context, masquerades sortedMasquerades) (net.Conn, MasqueradeInterface, func(bool) bool, error) {
	// never take more than a minute trying to find a dialer
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	masqueradesToTry := masquerades.sortedCopy()
	totalMasquerades := len(masqueradesToTry)
dialLoop:
	for _, m := range masqueradesToTry {
		select {
		case <-ctx.Done():
			log.Debugf("Timed out dialing to %v with %v total masquerades", m, totalMasquerades)
			break dialLoop
		default:
			// okay
		}
		conn, masqueradeGood, err := d.dialMasquerade(m)
		if err == nil {
			return conn, m, masqueradeGood, nil
		}
	}

	return nil, nil, nil, log.Errorf("could not dial any masquerade? tried %v", totalMasquerades)
}

func (d *direct) dialMasquerade(m MasqueradeInterface) (net.Conn, func(bool) bool, error) {
	// check to see if we've timed out

	log.Tracef("Dialing to %v", m)

	// We do the full TLS connection here because in practice the domains at a given IP
	// address can change frequently on CDNs, so the certificate may not match what
	// we expect.
	conn, retriable, err := d.doDial(m)
	masqueradeGood := func(good bool) bool {
		if good {
			m.markSucceeded()
		} else {
			m.markFailed()
		}
		d.markCacheDirty()
		return good
	}
	if err == nil {
		log.Debug("Returning connection")
		return conn, masqueradeGood, err
	} else if !retriable {
		log.Debugf("Dropping masquerade: non retryable error: %v", err)
		masqueradeGood(false)
	}
	return conn, masqueradeGood, err
}

func (d *direct) doDial(m MasqueradeInterface) (conn net.Conn, retriable bool, err error) {
	op := ops.Begin("dial_masquerade")
	defer op.End()
	op.Set("masquerade_domain", m.getDomain())
	op.Set("masquerade_ip", m.getIpAddress())

	conn, err = d.dialServerWith(m)
	if err != nil {
		op.FailIf(err)
		log.Debugf("Could not dial to %v, %v", m.getIpAddress(), err)
		// Don't re-add this candidate if it's any certificate error, as that
		// will just keep failing and will waste connections. We can't access the underlying
		// error at this point so just look for "certificate" and "handshake".
		if strings.Contains(err.Error(), "certificate") || strings.Contains(err.Error(), "handshake") {
			log.Debugf("Not re-adding candidate that failed on error '%v'", err.Error())
			retriable = false
		} else {
			log.Debugf("Unexpected error dialing, keeping masquerade: %v", err)
			retriable = true
		}
	} else {
		log.Debugf("Got successful connection to: %v", m)
		idleTimeout := 70 * time.Second

		log.Debugf("Wrapping connection in idletiming connection: %v", m)
		conn = idletiming.Conn(conn, idleTimeout, func() {
			log.Debugf("Connection to %v idle for %v, closed", conn.RemoteAddr(), idleTimeout)
		})
	}
	return
}

func (d *direct) dialServerWith(m MasqueradeInterface) (net.Conn, error) {
	op := ops.Begin("dial_server_with")
	defer op.End()

	op.Set("masquerade_domain", m.getDomain())
	op.Set("masquerade_ip", m.getIpAddress())

	conn, err := m.dial(d.certPool, d.clientHelloID)
	if err != nil && m != nil {
		err = fmt.Errorf("unable to dial masquerade %s: %s", m.getDomain(), err)
		op.FailIf(err)
	}
	return conn, err
}

func verifyPeerCertificate(rawCerts [][]byte, roots *x509.CertPool, domain string) error {
	if len(rawCerts) == 0 {
		return fmt.Errorf("no certificates presented")
	}
	cert, err := x509.ParseCertificate(rawCerts[0])
	if err != nil {
		return fmt.Errorf("unable to parse certificate: %w", err)
	}

	opts := []x509.VerifyOptions{generateVerifyOptions(roots, domain)}
	for i := range rawCerts {
		if i == 0 {
			continue
		}
		crt, err := x509.ParseCertificate(rawCerts[i])
		if err != nil {
			return fmt.Errorf("unable to parse intermediate certificate: %w", err)
		}

		for _, opt := range opts {
			opt.Intermediates.AddCert(crt)
		}
	}

	var verificationErrors error
	for _, opt := range opts {
		_, err := cert.Verify(opt)
		if err != nil {
			verificationErrors = errors.Join(verificationErrors, err)
		}
	}

	if verificationErrors != nil {
		return fmt.Errorf("certificate verification failed: %w", verificationErrors)
	}

	return nil
}

func generateVerifyOptions(roots *x509.CertPool, domain string) x509.VerifyOptions {
	return x509.VerifyOptions{
		Roots:         roots,
		CurrentTime:   time.Now(),
		DNSName:       domain,
		Intermediates: x509.NewCertPool(),
	}
}

// frontingTLSConfig builds a tls.Config for dialing the fronting domain. This is to establish the
// initial TCP connection to the CDN.
func (d *direct) frontingTLSConfig(m *Masquerade) *tls.Config {
	return &tls.Config{
		ServerName: m.Domain,
		RootCAs:    d.certPool,
	}
}

// frontedHTTPTransport is the transport to use to route to the actual fronted destination domain.
// This uses the pre-established connection to the CDN on the fronting domain.
func frontedHTTPTransport(conn net.Conn, disableKeepAlives bool) http.RoundTripper {
	return &directTransport{
		Transport: http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return conn, nil
			},
			TLSHandshakeTimeout: 40 * time.Second,
			DisableKeepAlives:   disableKeepAlives,
		},
	}
}

// directTransport is a wrapper struct enabling us to modify the protocol of outgoing
// requests to make them all HTTP instead of potentially HTTPS, which breaks our particular
// implemenation of direct domain fronting.
type directTransport struct {
	http.Transport
}

func (ddf *directTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	// The connection is already encrypted by domain fronting.  We need to rewrite URLs starting
	// with "https://" to "http://", lest we get an error for doubling up on TLS.

	// The RoundTrip interface requires that we not modify the memory in the request, so we just
	// create a copy.
	norm := new(http.Request)
	*norm = *req // includes shallow copies of maps, but okay
	norm.URL = new(url.URL)
	*norm.URL = *req.URL
	norm.URL.Scheme = "http"
	return ddf.Transport.RoundTrip(norm)
}

func cloneRequestWith(req *http.Request, frontedHost string, body io.ReadCloser) (*http.Request, error) {
	url := *req.URL
	url.Host = frontedHost
	r, err := http.NewRequest(req.Method, url.String(), body)
	if err != nil {
		return nil, err
	}

	for k, vs := range req.Header {
		if !strings.EqualFold(k, "Host") {
			v := make([]string, len(vs))
			copy(v, vs)
			r.Header[k] = v
		}
	}
	return r, nil
}
