package fronted

import (
	"bytes"
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
	"syscall"
	"time"

	tls "github.com/refraction-networking/utls"

	"github.com/getlantern/golog"
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

type workingFronts interface {
	onConnected(m MasqueradeInterface, connectTime time.Duration)
	onError(m MasqueradeInterface)
	workingFront() (MasqueradeInterface, error)
}

// fronted identifies working IP address/domain pairings for domain fronting and is
// an implementation of http.RoundTripper for the convenience of callers.
type fronted struct {
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
	workingFronts       workingFronts
	providersMu         sync.RWMutex
	masqueradesMu       sync.RWMutex
	frontedMu           sync.RWMutex
}

func newFronted(pool *x509.CertPool, providers map[string]*Provider,
	defaultProviderID, cacheFile string, clientHelloID tls.ClientHelloID,
	listener func(f *fronted),
	workingFronts workingFronts) (*fronted, error) {
	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}

	if size == 0 {
		return nil, fmt.Errorf("no masquerades found in providers")
	}

	// copy providers
	providersCopy := copyProviders(providers)

	f := &fronted{
		certPool:            pool,
		masquerades:         loadMasquerades(providersCopy, size),
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
		cacheSaveInterval:   defaultCacheSaveInterval,
		cacheDirty:          make(chan interface{}, 1),
		cacheClosed:         make(chan interface{}),
		defaultProviderID:   defaultProviderID,
		providers:           providersCopy,
		clientHelloID:       clientHelloID,
		workingFronts:       workingFronts,
	}

	if cacheFile != "" {
		f.initCaching(cacheFile)
	}
	f.findWorkingMasquerades(listener)

	return f, nil
}

func copyProviders(providers map[string]*Provider) map[string]*Provider {
	providersCopy := make(map[string]*Provider, len(providers))
	for key, p := range providers {
		providersCopy[key] = NewProvider(p.HostAliases, p.TestURL, p.Masquerades, p.Validator, p.PassthroughPatterns, p.SNIConfig, p.VerifyHostname)
		log.Debugf("Domain fronting provider is %v", providersCopy[key])
	}
	return providersCopy
}

func loadMasquerades(initial map[string]*Provider, size int) sortedMasquerades {
	log.Debugf("Loading candidates for %d providers", len(initial))
	defer log.Debug("Finished loading candidates")

	masquerades := make(sortedMasquerades, 0, size)
	for key, p := range initial {
		arr := p.Masquerades
		size := len(arr)

		// make a shuffled copy of arr
		// ('inside-out' Fisher-Yates)
		sh := make([]*Masquerade, size)
		for i := 0; i < size; i++ {
			j := rand.IntN(i + 1) // 0 <= j <= i
			sh[i] = sh[j]
			sh[j] = arr[i]
		}

		for _, c := range sh {
			masquerades = append(masquerades, &masquerade{Masquerade: *c, ProviderID: key})
		}
	}
	return masquerades
}

func (f *fronted) updateConfig(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string) {
	// Make copies just to avoid any concurrency issues with access that may be happening on the
	// caller side.
	providersCopy := copyProviders(providers)
	f.frontedMu.Lock()
	defer f.frontedMu.Unlock()
	f.addProviders(providersCopy)
	f.addMasquerades(loadMasquerades(providersCopy, len(providersCopy)))
	f.defaultProviderID = defaultProviderID
	f.certPool = pool
}

func (f *fronted) addProviders(providers map[string]*Provider) {
	// Add new providers to the existing providers map, overwriting any existing ones.
	f.providersMu.Lock()
	defer f.providersMu.Unlock()
	for key, p := range providers {
		f.providers[key] = p
	}
}

func (f *fronted) addMasquerades(masquerades sortedMasquerades) {
	// Add new masquerades to the existing masquerades slice, but add them at the beginning.
	f.masqueradesMu.Lock()
	defer f.masqueradesMu.Unlock()
	f.masquerades = append(masquerades, f.masquerades...)
}

func (f *fronted) providerFor(m MasqueradeInterface) *Provider {
	pid := m.getProviderID()
	if pid == "" {
		pid = f.defaultProviderID
	}
	return f.providers[pid]
}

// Vet vets the specified Masquerade, verifying certificate using the given CertPool.
// This is used in genconfig.
func Vet(m *Masquerade, pool *x509.CertPool, testURL string) bool {
	d := &fronted{
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
func (f *fronted) findWorkingMasquerades(listener func(f *fronted)) {
	// vet masquerades in batches
	const batchSize int = 40
	var successful atomic.Uint32

	// We loop through all of them until we have 4 successful ones.
	for i := 0; i < f.masqueradeSize() && successful.Load() < 4; i += batchSize {
		f.vetBatch(i, batchSize, &successful, listener)
	}
}

func (f *fronted) masqueradeSize() int {
	f.masqueradesMu.Lock()
	defer f.masqueradesMu.Unlock()
	return len(f.masquerades)
}

func (f *fronted) masqueradeAt(i int) MasqueradeInterface {
	f.masqueradesMu.Lock()
	defer f.masqueradesMu.Unlock()
	return f.masquerades[i]
}

func (f *fronted) vetBatch(start, batchSize int, successful *atomic.Uint32, listener func(f *fronted)) {
	log.Debugf("Vetting masquerade batch %d-%d", start, start+batchSize)
	var wg sync.WaitGroup
	for i := start; i < start+batchSize && i < f.masqueradeSize(); i++ {
		wg.Add(1)
		go func(m MasqueradeInterface) {
			defer wg.Done()
			working, connectTime := f.vetMasquerade(m)
			if working {
				successful.Add(1)
				f.workingFronts.onConnected(m, connectTime)
				if listener != nil {
					go listener(f)
				}
			} else {
				f.workingFronts.onError(m)
			}
		}(f.masqueradeAt(i))
	}
	wg.Wait()
}

func (f *fronted) vetMasquerade(m MasqueradeInterface) (bool, time.Duration) {
	start := time.Now()
	conn, masqueradeGood, err := f.dialMasquerade(m)
	if err != nil {
		log.Debugf("unexpected error vetting masquerades: %v", err)
		return false, time.Since(start)
	}
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	provider := f.providerFor(m)
	if provider == nil {
		log.Debugf("Skipping masquerade with disabled/unknown provider id '%s' not in %v",
			m.getProviderID(), f.providers)
		return false, time.Since(start)
	}
	if !masqueradeGood(m.postCheck(conn, provider.TestURL)) {
		log.Debugf("Unsuccessful vetting with POST request, discarding masquerade")
		return false, time.Since(start)
	}

	log.Debugf("Successfully vetted one masquerade %v", m)
	return true, time.Since(start)
}

// RoundTrip loops through all available masquerades, sorted by the one that most recently
// connected, retrying several times on failures.
func (f *fronted) RoundTrip(req *http.Request) (*http.Response, error) {
	res, _, err := f.RoundTripHijack(req)
	return res, err
}

// RoundTripHijack loops through all available masquerades, sorted by the one that most
// recently connected, retrying several times on failures.
func (f *fronted) RoundTripHijack(req *http.Request) (*http.Response, net.Conn, error) {
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

	const tries = 6

	for i := 0; i < tries; i++ {
		if i > 0 {
			log.Debugf("Retrying domain-fronted request, pass %d", i)
		}

		m, err := f.workingFronts.workingFront()
		if err != nil {
			// For some reason we have no working fronts. Sleep for a bit and try again.
			time.Sleep(1 * time.Second)
			continue
		}

		conn, masqueradeGood, err := f.dialMasquerade(m)
		if err != nil {
			log.Debugf("Could not dial to %v: %v", m, err)
			f.workingFronts.onError(m)
			continue
		}

		resp, conn, err := f.request(req, conn, m, originHost, getBody, masqueradeGood)
		if err != nil {
			log.Debugf("Could not complete request: %v", err)
			f.workingFronts.onError(m)
		} else {
			return resp, conn, nil
		}
	}

	return nil, nil, op.FailIf(errors.New("could not complete request even with retries"))
}

func (f *fronted) request(req *http.Request, conn net.Conn, m MasqueradeInterface, originHost string, getBody func() io.ReadCloser, masqueradeGood func(bool) bool) (*http.Response, net.Conn, error) {
	op := ops.Begin("fronted_request")
	defer op.End()
	provider := f.providerFor(m)
	if provider == nil {
		log.Debugf("Skipping masquerade with disabled/unknown provider '%s'", m.getProviderID())
		masqueradeGood(false)
		return nil, nil, op.FailIf(log.Errorf("Skipping masquerade with disabled/unknown provider '%s'", m.getProviderID()))
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
	disableKeepAlives := true
	if strings.EqualFold(reqi.Header.Get("Connection"), "upgrade") {
		disableKeepAlives = false
	}

	tr := frontedHTTPTransport(conn, disableKeepAlives)
	resp, err := tr.RoundTrip(reqi)
	if err != nil {
		log.Debugf("Could not complete request: %v", err)
		masqueradeGood(false)
		return nil, nil, err
	}

	err = provider.ValidateResponse(resp)
	if err != nil {
		log.Debugf("Could not complete request: %v", err)
		resp.Body.Close()
		masqueradeGood(false)
		return nil, nil, err
	}

	masqueradeGood(true)
	return resp, conn, nil
}

func (f *fronted) dialMasquerade(m MasqueradeInterface) (net.Conn, func(bool) bool, error) {
	log.Tracef("Dialing to %v", m)

	// We do the full TLS connection here because in practice the domains at a given IP
	// address can change frequently on CDNs, so the certificate may not match what
	// we expect.
	conn, retriable, err := f.doDial(m)
	masqueradeGood := func(good bool) bool {
		if good {
			m.markSucceeded()
		} else {
			m.markFailed()
		}
		f.markCacheDirty()
		return good
	}
	if err == nil {
		log.Debugf("Returning connection for masquerade: %v", m)
		return conn, masqueradeGood, err
	} else if !retriable {
		log.Debugf("Dropping masquerade: non retryable error: %v", err)
		masqueradeGood(false)
	}
	return conn, masqueradeGood, err
}

func (f *fronted) doDial(m MasqueradeInterface) (net.Conn, bool, error) {
	op := ops.Begin("dial_masquerade")
	defer op.End()
	op.Set("masquerade_domain", m.getDomain())
	op.Set("masquerade_ip", m.getIpAddress())
	op.Set("masquerade_provider", m.getProviderID())

	var conn net.Conn
	var err error
	retriable := false
	start := time.Now()
	conn, err = m.dial(f.certPool, f.clientHelloID)
	if err != nil {
		if !isNetworkUnreachable(err) {
			op.FailIf(err)
		}
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
		log.Debugf("Got successful connection to: %+v in %v", m, time.Since(start))
	}
	return conn, retriable, err
}

func isNetworkUnreachable(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		// The following error verifications look for errors that generally happen at Linux/Unix devices
		if errors.Is(opErr.Err, syscall.ENETUNREACH) || errors.Is(opErr.Err, syscall.EHOSTUNREACH) {
			return true
		}

		// The string verification errors use a broader approach with errors from windows and also linux/unix devices
		errMsg := opErr.Err.Error()
		if strings.Contains(errMsg, "network is unreachable") ||
			strings.Contains(errMsg, "no route to host") ||
			strings.Contains(errMsg, "unreachable network") ||
			strings.Contains(errMsg, "unreachable host") {
			return true
		}
	}
	return false
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
			IdleConnTimeout:     70 * time.Second,
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
	defer func(op ops.Op) { op.End() }(ops.Begin("direct_transport_roundtrip"))

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
