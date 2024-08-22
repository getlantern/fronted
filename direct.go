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
	"time"

	tls "github.com/refraction-networking/utls"

	"github.com/getlantern/golog"
	"github.com/getlantern/idletiming"
	"github.com/getlantern/netx"
	"github.com/getlantern/ops"
	"github.com/getlantern/tlsdialer/v3"
)

const (
	numberToVetInitially       = 10
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

func (d *direct) providerFor(m *masquerade) *Provider {
	pid := m.ProviderID
	if pid == "" {
		pid = d.defaultProviderID
	}
	return d.providers[pid]
}

// Vet vets the specified Masquerade, verifying certificate using the given CertPool
func Vet(m *Masquerade, pool *x509.CertPool, testURL string) bool {
	return vet(m, pool, testURL)
}

func vet(m *Masquerade, pool *x509.CertPool, testURL string) bool {
	op := ops.Begin("vet_masquerade")
	defer op.End()
	op.Set("masquerade_domain", m.Domain)
	op.Set("masquerade_ip", m.IpAddress)

	d := &direct{
		certPool:            pool,
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
	}
	conn, _, err := d.doDial(m)
	if err != nil {
		op.FailIf(err)
		return false
	}
	defer conn.Close()
	return postCheck(conn, testURL)
}

func (d *direct) vet(numberToVet int) {
	log.Debugf("Vetting %d initial candidates in series", numberToVet)
	for i := 0; i < numberToVet; i++ {
		d.vetOne()
	}
}

func (d *direct) vetOne() {
	// We're just testing the ability to connect here, destination site doesn't
	// really matter
	log.Trace("Vetting one")
	unvettedMasquerades := make([]*masquerade, 0, len(d.masquerades))
	for _, m := range d.masquerades {
		if m.lastSucceeded().IsZero() {
			unvettedMasquerades = append(unvettedMasquerades, m)
		}
	}

	// Don't take more than 10 seconds to dial a masquerade for vetting
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, m, masqueradeGood, err := d.dialWith(ctx, unvettedMasquerades)
	if err != nil {
		log.Errorf("unexpected error vetting masquerades: %v", err)
		return
	}
	defer conn.Close()

	provider := d.providerFor(m)
	if provider == nil {
		log.Tracef("Skipping masquerade with disabled/unknown provider id '%s'", m.ProviderID)
		return
	}

	if !masqueradeGood(postCheck(conn, provider.TestURL)) {
		log.Tracef("Unsuccessful vetting with POST request, discarding masquerade")
		return
	}

	log.Trace("Finished vetting one")
}

// postCheck does a post with invalid data to verify domain-fronting works
func postCheck(conn net.Conn, testURL string) bool {
	client := &http.Client{
		Transport: frontedHTTPTransport(conn, true),
	}
	return doCheck(client, http.MethodPost, http.StatusAccepted, testURL)
}

func doCheck(client *http.Client, method string, expectedStatus int, u string) bool {
	op := ops.Begin("check_masquerade")
	defer op.End()

	isPost := method == http.MethodPost
	var requestBody io.Reader
	if isPost {
		requestBody = strings.NewReader("a")
	}
	req, _ := http.NewRequest(method, u, requestBody)
	if isPost {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := client.Do(req)
	if err != nil {
		op.FailIf(err)
		log.Debugf("Unsuccessful vetting with %v request, discarding masquerade: %v", method, err)
		return false
	}
	if resp.Body != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	if resp.StatusCode != expectedStatus {
		op.Set("response_status", resp.StatusCode)
		op.Set("expected_status", expectedStatus)
		msg := fmt.Sprintf("Unexpected response status vetting masquerade, expected %d got %d: %v", expectedStatus, resp.StatusCode, resp.Status)
		op.FailIf(fmt.Errorf(msg))
		log.Debug(msg)
		return false
	}
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

		conn, m, masqueradeGood, err := d.dial(req.Context())
		if err != nil {
			// unable to find good masquerade, fail
			op.FailIf(err)
			return nil, nil, err
		}
		provider := d.providerFor(m)
		if provider == nil {
			log.Debugf("Skipping masquerade with disabled/unknown provider '%s'", m.ProviderID)
			masqueradeGood(false)
			continue
		}
		frontedHost := provider.Lookup(originHost)
		if frontedHost == "" {
			// this error is not the masquerade's fault in particular
			// so it is returned as good.
			conn.Close()
			masqueradeGood(true)
			err := fmt.Errorf("no domain fronting mapping for '%s'. Please add it to provider_map.yaml or equivalent for %s", m.ProviderID, originHost)
			op.FailIf(err)
			return nil, nil, err
		}
		log.Tracef("Translated origin %s -> %s for provider %s...", originHost, frontedHost, m.ProviderID)

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

// Dial dials out using a masquerade. If the available masquerade fails, it
// retries with others until it either succeeds or exhausts the available
// masquerades. If successful, it returns a connection to the masquerade,
// the selected masquerade, and a function that the caller can use to
// tell us whether the masquerade is good or not (i.e. if masquerade was good,
// keep it).
func (d *direct) dial(ctx context.Context) (net.Conn, *masquerade, func(bool) bool, error) {
	conn, m, masqueradeGood, err := d.dialWith(ctx, d.masquerades)
	return conn, m, masqueradeGood, err
}

func (d *direct) dialWith(ctx context.Context, masquerades sortedMasquerades) (net.Conn, *masquerade, func(bool) bool, error) {
	// never take more than a minute trying to find a dialer
	ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	defer cancel()

	masqueradesToTry := masquerades.sortedCopy()
dialLoop:
	for _, m := range masqueradesToTry {
		// check to see if we've timed out
		select {
		case <-ctx.Done():
			break dialLoop
		default:
			// okay
		}

		log.Tracef("Dialing to %v", m)

		// We do the full TLS connection here because in practice the domains at a given IP
		// address can change frequently on CDNs, so the certificate may not match what
		// we expect.
		conn, retriable, err := d.doDial(&m.Masquerade)
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
			log.Trace("Returning connection")
			return conn, m, masqueradeGood, err
		} else if !retriable {
			log.Debugf("Dropping masquerade: non retryable error: %v", err)
			masqueradeGood(false)
		}
	}

	return nil, nil, nil, errors.New("could not dial any masquerade?")
}

func (d *direct) doDial(m *Masquerade) (conn net.Conn, retriable bool, err error) {
	op := ops.Begin("dial_masquerade")
	defer op.End()
	op.Set("masquerade_domain", m.Domain)
	op.Set("masquerade_ip", m.IpAddress)

	conn, err = d.dialServerWith(m)
	if err != nil {
		op.FailIf(err)
		log.Tracef("Could not dial to %v, %v", m.IpAddress, err)
		// Don't re-add this candidate if it's any certificate error, as that
		// will just keep failing and will waste connections. We can't access the underlying
		// error at this point so just look for "certificate" and "handshake".
		if strings.Contains(err.Error(), "certificate") || strings.Contains(err.Error(), "handshake") {
			log.Debugf("Not re-adding candidate that failed on error '%v'", err.Error())
			retriable = false
		} else {
			log.Tracef("Unexpected error dialing, keeping masquerade: %v", err)
			retriable = true
		}
	} else {
		log.Debugf("Got successful connection to: %v", m)
		idleTimeout := 70 * time.Second

		log.Debugf("Wrapping connection in idletiming connection: %v", m)
		conn = idletiming.Conn(conn, idleTimeout, func() {
			log.Tracef("Connection to %v idle for %v, closed", conn.RemoteAddr(), idleTimeout)
		})
	}
	return
}

func (d *direct) dialServerWith(m *Masquerade) (net.Conn, error) {
	op := ops.Begin("dial_server_with")
	defer op.End()

	op.Set("masquerade_domain", m.Domain)
	op.Set("masquerade_ip", m.IpAddress)

	tlsConfig := d.frontingTLSConfig(m)
	dialTimeout := 10 * time.Second
	addr := m.IpAddress
	var sendServerNameExtension bool

	if m.SNI != "" {
		sendServerNameExtension = true

		op.Set("arbitrary_sni", m.SNI)
		tlsConfig.ServerName = m.SNI
		tlsConfig.InsecureSkipVerify = true
		tlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
			log.Tracef("verifying peer certificate for masquerade domain [%s] and SNI [%s]", m.Domain, m.SNI)
			return verifyPeerCertificate(rawCerts, d.certPool, m.Domain, m.SNI)
		}

	}

	_, _, err := net.SplitHostPort(addr)
	if err != nil {
		addr = net.JoinHostPort(addr, "443")
	}

	dialer := &tlsdialer.Dialer{
		DoDial:         netx.DialTimeout,
		Timeout:        dialTimeout,
		SendServerName: sendServerNameExtension,
		Config:         tlsConfig,
		ClientHelloID:  d.clientHelloID,
	}
	conn, err := dialer.Dial("tcp", addr)
	if err != nil && m != nil {
		err = fmt.Errorf("unable to dial masquerade %s: %s", m.Domain, err)
		op.FailIf(err)
	}
	return conn, err
}

func verifyPeerCertificate(rawCerts [][]byte, roots *x509.CertPool, domain string, sni string) error {
	if len(rawCerts) == 0 {
		return fmt.Errorf("no certificates presented")
	}
	cert, err := x509.ParseCertificate(rawCerts[0])
	if err != nil {
		return fmt.Errorf("unable to parse certificate: %w", err)
	}

	masqueradeOpts := x509.VerifyOptions{
		Roots:         roots,
		CurrentTime:   time.Now(),
		DNSName:       domain,
		Intermediates: x509.NewCertPool(),
	}

	sniOpts := x509.VerifyOptions{
		Roots:         roots,
		CurrentTime:   time.Now(),
		DNSName:       sni,
		Intermediates: x509.NewCertPool(),
	}

	for i := range rawCerts {
		if i == 0 {
			continue
		}
		crt, err := x509.ParseCertificate(rawCerts[i])
		if err != nil {
			return fmt.Errorf("unable to parse intermediate certificate: %w", err)
		}
		masqueradeOpts.Intermediates.AddCert(crt)
		sniOpts.Intermediates.AddCert(crt)
	}

	_, sniErr := cert.Verify(sniOpts)
	_, masqueradeErr := cert.Verify(masqueradeOpts)
	if masqueradeErr != nil && sniErr != nil {
		return fmt.Errorf("certificate verification failed for masquerade and SNI: [%w],[%w]", masqueradeErr, sniErr)
	}

	return nil
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
