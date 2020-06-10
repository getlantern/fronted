package fronted

import (
	"bytes"
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/getlantern/golog"
	"github.com/getlantern/idletiming"
	"github.com/getlantern/netx"
	"github.com/getlantern/tlsdialer/v3"
	tls "github.com/refraction-networking/utls"
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

// DirectOptions defines optional paramaters for NewDirect and FrontingContext.NewDirect.
type DirectOptions struct {
	// CertPool sets the root CAs used to verify server certificates. If nil, the host's root CA set
	// will be used.
	CertPool *x509.CertPool

	// CacheFile, if provided, will be used to cache providers. Multiple calls to NewDirect may be
	// made with the same cache file. However, cache files should *not* be shared across contexts.
	CacheFile string

	// ClientHelloID, if provided, specifies the ID of a ClientHello to mimic. See
	// https://pkg.go.dev/github.com/refraction-networking/utls?tab=doc#pkg-variables
	ClientHelloID tls.ClientHelloID

	// DialTransport is used to establish the transport connection to the masquerade. This will
	// almost certainly be a TCP connection. If nil, getlantern/netx.DialContext will be used.
	DialTransport func(ctx context.Context, network, address string) (net.Conn, error)
}

// direct is an implementation of http.RoundTripper
type direct struct {
	certPool          *x509.CertPool
	candidates        chan masquerade
	masquerades       chan masquerade
	cache             *masqueradeCache
	defaultProviderID string
	providers         map[string]*Provider
	ready             chan struct{}
	readyOnce         sync.Once
	dialTransport     func(ctx context.Context, network, address string) (net.Conn, error)
	clientHelloID     tls.ClientHelloID
}

// Returns errorTimeout if the direct cannot be initialized in the provided timeout.
func newDirect(
	ctx context.Context, providers map[string]*Provider, defaultProviderID string,
	cache *masqueradeCache, opts DirectOptions) (*direct, error) {

	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}
	if opts.DialTransport == nil {
		opts.DialTransport = netx.DialContext
	}
	d := &direct{
		certPool:          opts.CertPool,
		candidates:        make(chan masquerade, size),
		masquerades:       make(chan masquerade, size),
		cache:             cache,
		defaultProviderID: defaultProviderID,
		providers:         make(map[string]*Provider),
		ready:             make(chan struct{}),
		dialTransport:     opts.DialTransport,
		clientHelloID:     opts.ClientHelloID,
	}
	// copy providers
	for k, p := range providers {
		d.providers[k] = NewProvider(
			p.HostAliases, p.TestURL, p.Masquerades, p.Validator, p.PassthroughPatterns)
	}
	numberToVet := numberToVetInitially
	pulledFromCache, err := d.initFromCache()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize from cache file: %w", err)
	}
	numberToVet -= pulledFromCache
	d.loadCandidates()
	if numberToVet > 0 {
		d.vet(numberToVet)
	} else {
		log.Debugf("Not vetting any masquerades because we have enough cached in %s", cache.filename)
		d.signalReady()
	}
	select {
	case <-d.ready:
		return d, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (d *direct) initFromCache() (used int, err error) {
	if d.cache == nil {
		return 0, nil
	}
	inCache, err := d.cache.read()
	if err != nil {
		return 0, fmt.Errorf("failed to read cache file: %w", err)
	}
	log.Debugf("Initializing from cache of %d masquerades", len(inCache))

	for _, m := range inCache {
		// Fill in default for masquerades lacking a provider ID.
		if m.ProviderID == "" {
			m.ProviderID = d.defaultProviderID
		}
		if _, ok := d.providers[m.ProviderID]; !ok {
			// Skip entries for providers that are not configured.
			log.Debugf("Skipping cached entry for unknown/disabled provider %s", m.ProviderID)
			continue
		}
		select {
		case d.masquerades <- m:
			used++
		default:
			// Channel is full, that's okay.
		}
	}
	return used, nil
}

func (d *direct) loadCandidates() {
	log.Debug("Loading candidates")
	for key, p := range d.providers {
		arr := p.Masquerades
		size := len(arr)
		log.Tracef("Adding %d candidates for %v", size, key)

		// make a shuffled copy of arr
		// ('inside-out' Fisher-Yates)
		sh := make([]*Masquerade, size)
		for i := 0; i < size; i++ {
			j := rand.Intn(i + 1) // 0 <= j <= i
			sh[i] = sh[j]
			sh[j] = arr[i]
		}

		for _, c := range sh {
			log.Trace("Adding candidate")
			// Note: we ensured in newDirect that the buffer for d.candidates is large enough to
			// hold all masquerades in d.providers.
			d.candidates <- masquerade{Masquerade: *c, ProviderID: key}
		}
	}
}

func (d *direct) signalReady() {
	d.readyOnce.Do(func() {
		close(d.ready)
	})
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
	d := &direct{
		certPool:      pool,
		dialTransport: netx.DialContext,
	}
	conn, _, err := d.doDial(m)
	if err != nil {
		return false
	}
	defer conn.Close()
	return postCheck(conn, testURL)
}

func (d *direct) vet(numberToVet int) {
	log.Tracef("Vetting %d initial candidates in parallel", numberToVet)
	for i := 0; i < numberToVet; i++ {
		go d.vetOneUntilGood()
	}
}

func (d *direct) vetOneUntilGood() {
	for {
		if !d.vetOne() {
			return
		}
	}
}

func (d *direct) vetOne() bool {
	// We're just testing the ability to connect here, destination site doesn't
	// really matter
	log.Trace("Vetting one")
	conn, m, masqueradeGood, masqueradesRemain, err := d.dialWith(d.candidates)
	if err != nil {
		return masqueradesRemain
	}
	defer conn.Close()

	provider := d.providerFor(m)
	if provider == nil {
		log.Tracef("Skipping masquerade with disabled/unknown provider id '%s'", m.ProviderID)
		return masqueradesRemain
	}

	if !masqueradeGood(postCheck(conn, provider.TestURL)) {
		log.Tracef("Unsuccessful vetting with POST request, discarding masquerade")
		return masqueradesRemain
	}

	log.Trace("Finished vetting one")
	// signal that at least one
	// masquerade has been vetted successfully.
	d.signalReady()
	return false
}

// postCheck does a post with invalid data to verify domain-fronting works
func postCheck(conn net.Conn, testURL string) bool {
	client := &http.Client{
		Transport: frontedHTTPTransport(conn, true),
	}
	return doCheck(client, http.MethodPost, http.StatusAccepted, testURL)
}

func doCheck(client *http.Client, method string, expectedStatus int, u string) bool {
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
		log.Debugf("Unsuccessful vetting with %v request, discarding masquerade: %v", method, err)
		return false
	}
	if resp.Body != nil {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
	if resp.StatusCode != expectedStatus {
		log.Debugf("Unexpected response status vetting masquerade, expected %d got %d: %v", expectedStatus, resp.StatusCode, resp.Status)
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
	isIdempotent := req.Method != http.MethodPost && req.Method != http.MethodPatch

	originHost := req.URL.Hostname()

	var body []byte
	var err error
	if isIdempotent && req.Body != nil {
		// store body in-memory to be able to replay it if necessary
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, nil, fmt.Errorf("Unable to read request body: %v", err)
		}
	}

	getBody := func() io.ReadCloser {
		if req.Body == nil {
			return nil
		}

		if !isIdempotent {
			return req.Body
		}
		return ioutil.NopCloser(bytes.NewReader(body))
	}

	tries := 1
	if isIdempotent {
		tries = maxTries
	}

	for i := 0; i < tries; i++ {
		if i > 0 {
			log.Debugf("Retrying domain-fronted request, pass %d", i)
		}

		conn, m, masqueradeGood, err := d.dial()
		if err != nil {
			// unable to find good masquerade, fail
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
			return nil, nil, fmt.Errorf("No alias for host %s", originHost)
		}
		log.Tracef("Translated origin %s -> %s for provider %s...", originHost, frontedHost, m.ProviderID)

		reqi, err := cloneRequestWith(req, frontedHost, getBody())
		if err != nil {
			log.Errorf("Failed to copy http request?")
			masqueradeGood(true)
			continue
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

	return nil, nil, errors.New("Could not complete request even with retries")
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
// keep it, else vet a new one).
func (d *direct) dial() (net.Conn, *masquerade, func(bool) bool, error) {
	conn, m, masqueradeGood, _, err := d.dialWith(d.masquerades)
	return conn, m, masqueradeGood, err
}

func (d *direct) dialWith(in chan masquerade) (net.Conn, *masquerade, func(bool) bool, bool, error) {
	retryLater := make([]masquerade, 0)
	defer func() {
		for _, m := range retryLater {
			// when network just recovered from offline, retryLater has more
			// elements than the capacity of the channel.
			select {
			case in <- m:
			default:
				log.Debug("Dropping masquerade: retry channel full")
			}
		}
	}()

	for {
		var m masquerade
		select {
		case m = <-in:
			log.Trace("Got vetted masquerade")
		default:
			log.Trace("No vetted masquerade found, falling back to unvetted candidate")
			select {
			case m = <-d.candidates:
				log.Trace("Got unvetted masquerade")
			default:
				return nil, nil, nil, false, errors.New("Could not dial any masquerade?")
			}
		}

		log.Tracef("Dialing to %v", m)

		// We do the full TLS connection here because in practice the domains at a given IP
		// address can change frequently on CDNs, so the certificate may not match what
		// we expect.
		conn, retriable, err := d.doDial(&m.Masquerade)
		if err == nil {
			log.Trace("Returning connection")
			masqueradeGood := func(good bool) bool {
				if good {
					m.LastVetted = time.Now()
					// Requeue the working connection to masquerades
					d.masquerades <- m
					if d.cache != nil {
						d.cache.write(m)
					}
				} else {
					go d.vetOneUntilGood()
				}

				return good
			}
			return conn, &m, masqueradeGood, true, err
		} else if retriable {
			retryLater = append(retryLater, m)
		} else {
			log.Debugf("Dropping masquerade: non retryable error: %v", err)
		}
	}
}

func (d *direct) doDial(m *Masquerade) (conn net.Conn, retriable bool, err error) {
	conn, err = d.dialServerWith(m)
	if err != nil {
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
	tlsConfig := d.frontingTLSConfig(m)
	dialTimeout := 10 * time.Second
	sendServerNameExtension := false
	addr := m.IpAddress

	_, _, err := net.SplitHostPort(addr)
	if err != nil {
		addr = net.JoinHostPort(addr, "443")
	}

	dialer := &tlsdialer.Dialer{
		DoDial: func(network, address string, timeout time.Duration) (net.Conn, error) {
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			return d.dialTransport(ctx, network, address)
		},
		Timeout:        dialTimeout,
		SendServerName: sendServerNameExtension,
		Config:         tlsConfig,
		ClientHelloID:  d.clientHelloID,
	}
	conn, err := dialer.Dial("tcp", addr)

	if err != nil && m != nil {
		err = fmt.Errorf("Unable to dial masquerade %s: %s", m.Domain, err)
	}
	return conn, err
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
