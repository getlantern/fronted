package fronted

import (
	"bytes"
	"crypto/tls"
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

	"github.com/getlantern/eventual"
	"github.com/getlantern/golog"
	"github.com/getlantern/idletiming"
	"github.com/getlantern/netx"
	"github.com/getlantern/tlsdialer"
)

const (
	numberToVetInitially       = 10
	defaultMaxAllowedCachedAge = 24 * time.Hour
	defaultMaxCacheSize        = 1000
	defaultCacheSaveInterval   = 5 * time.Second
	maxTries                   = 6
)

var (
	log       = golog.LoggerFor("fronted")
	_instance = eventual.NewValue()

	// Shared client session cache for all connections
	clientSessionCache = tls.NewLRUClientSessionCache(1000)
)

// direct is an implementation of http.RoundTripper
type direct struct {
	tlsConfigsMutex     sync.Mutex
	tlsConfigs          map[string]*tls.Config
	certPool            *x509.CertPool
	candidates          chan masquerade
	masquerades         chan masquerade
	maxAllowedCachedAge time.Duration
	maxCacheSize        int
	cacheSaveInterval   time.Duration
	toCache             chan masquerade
	defaultProviderID   string
	providers           map[string]*Provider
	ready               chan struct{}
	readyOnce           sync.Once
}

// Configure sets the masquerades to use, the trusted root CAs, and the
// cache file for caching masquerades to set up direct domain fronting.
// defaultProviderID is used when a masquerade without a provider is
// encountered (eg in a cache file)
func Configure(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string, cacheFile string) {
	log.Trace("Configuring fronted")

	if providers == nil || len(providers) == 0 {
		log.Errorf("No fronted providers!!")
		return
	}

	CloseCache()

	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}

	if size == 0 {
		log.Errorf("No masquerades!!")
		return
	}

	d := &direct{
		tlsConfigs:          make(map[string]*tls.Config),
		certPool:            pool,
		candidates:          make(chan masquerade, size),
		masquerades:         make(chan masquerade, size),
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
		cacheSaveInterval:   defaultCacheSaveInterval,
		toCache:             make(chan masquerade, defaultMaxCacheSize),
		defaultProviderID:   defaultProviderID,
		providers:           make(map[string]*Provider),
		ready:               make(chan struct{}),
	}

	// copy providers
	for k, p := range providers {
		d.providers[k] = NewProvider(p.HostAliases, p.TestURL, p.Masquerades)
	}

	numberToVet := numberToVetInitially
	if cacheFile != "" {
		numberToVet -= d.initCaching(cacheFile)
	}

	d.loadCandidates(d.providers)
	if numberToVet > 0 {
		d.vet(numberToVet)
	} else {
		log.Debug("Not vetting any masquerades because we have enough cached ones")
		d.signalReady()
	}
	_instance.Set(d)
}

func (d *direct) loadCandidates(initial map[string]*Provider) {
	log.Debug("Loading candidates")
	for key, p := range initial {
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
		tlsConfigs:          make(map[string]*tls.Config),
		certPool:            pool,
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
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
		Transport: httpTransport(conn, nil),
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

// NewDirect creates a new http.RoundTripper that does direct domain fronting.
// If it can't obtain a working masquerade within the given timeout, it will
// return nil/false.
func NewDirect(timeout time.Duration) (http.RoundTripper, bool) {
	start := time.Now()
	instance, ok := _instance.Get(timeout)
	if !ok {
		log.Errorf("No DirectHttpClient available within %v", timeout)
		return nil, false
	}
	remaining := timeout - time.Since(start)

	// Wait to be signalled that at least one masquerade has been vetted...
	select {
	case <-instance.(*direct).ready:
		return instance.(http.RoundTripper), true
	case <-time.After(remaining):
		log.Errorf("No DirectHttpClient available within %v", timeout)
		return nil, false
	}
}

// Do continually retries a given request until it succeeds because some
// fronting providers will return a 403 for some domains.
func (d *direct) RoundTrip(req *http.Request) (*http.Response, error) {
	isIdempotent := req.Method != http.MethodPost && req.Method != http.MethodPatch

	originHost := req.URL.Host

	var body []byte
	var err error
	if isIdempotent && req.Body != nil {
		// store body in-memory to be able to replay it if necessary
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("Unable to read request body: %v", err)
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
			return nil, err
		}
		provider := d.providerFor(m)
		if provider == nil {
			log.Debugf("Skipping masquerade with disabled/unknown provider '%s'", m.ProviderID)
			masqueradeGood(false)
			continue
		}
		frontedHost := provider.Lookup(originHost)
		if frontedHost == "" {
			log.Tracef("Not translating unknown origin %s...", originHost)
			frontedHost = originHost
		} else {
			log.Tracef("Translated origin %s -> %s for provider %s...", originHost, frontedHost, m.ProviderID)
		}

		reqi, err := cloneRequestWith(req, frontedHost, getBody())
		if err != nil {
			log.Errorf("Failed to copy http request?")
			masqueradeGood(true)
			continue
		}

		tr := httpTransport(conn, clientSessionCache)
		resp, err := tr.RoundTrip(reqi)
		if err != nil {
			log.Debugf("Could not complete request %v", err)
			masqueradeGood(false)
			continue
		}

		if resp.StatusCode == http.StatusForbidden {
			log.Debugf("Could not complete request due to response status: %v", resp.Status)
			resp.Body.Close()
			masqueradeGood(false)
			continue
		}

		masqueradeGood(true)
		return resp, nil
	}

	return nil, errors.New("Could not complete request even with retries")
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
			copy(vs, v)
			r.Header[k] = v
		}
	}
	return r, nil
}

// Dial dials out using a masquerade. If the available masquerade fails, it
// retries with others until it either succeeds or exhausts the available
// masquerades. If successful, it returns a function that the caller can use to
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
					select {
					case d.toCache <- m:
						// ok
					default:
						// cache writing has fallen behind, drop masquerade
						log.Debug("Dropping masquerade: cache writing is behind")
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
	tlsConfig := d.tlsConfig(m)
	dialTimeout := 10 * time.Second
	sendServerNameExtension := false
	addr := m.IpAddress
	if strings.IndexByte(addr, ':') == -1 {
		addr = addr + ":443"
	}

	conn, err := tlsdialer.DialTimeout(
		netx.DialTimeout,
		dialTimeout,
		"tcp",
		addr,
		sendServerNameExtension, // SNI or no
		tlsConfig)

	if err != nil && m != nil {
		err = fmt.Errorf("Unable to dial masquerade %s: %s", m.Domain, err)
	}
	return conn, err
}

// tlsConfig builds a tls.Config for dialing the upstream host. Constructed
// tls.Configs are cached on a per-masquerade basis to enable client session
// caching and reduce the amount of PEM certificate parsing.
func (d *direct) tlsConfig(m *Masquerade) *tls.Config {
	d.tlsConfigsMutex.Lock()
	defer d.tlsConfigsMutex.Unlock()

	tlsConfig := d.tlsConfigs[m.Domain]
	if tlsConfig == nil {
		tlsConfig = &tls.Config{
			ClientSessionCache: tls.NewLRUClientSessionCache(1000),
			InsecureSkipVerify: false,
			ServerName:         m.Domain,
			RootCAs:            d.certPool,
		}
		d.tlsConfigs[m.Domain] = tlsConfig
	}

	return tlsConfig
}

func httpTransport(conn net.Conn, clientSessionCache tls.ClientSessionCache) http.RoundTripper {
	return &directTransport{
		Transport: http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return conn, nil
			},
			TLSHandshakeTimeout: 40 * time.Second,
			DisableKeepAlives:   true,
			TLSClientConfig: &tls.Config{
				ClientSessionCache: clientSessionCache,
			},
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
