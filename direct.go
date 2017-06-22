package fronted

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
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
	testURL                    = "http://d157vud77ygy87.cloudfront.net/measurements" // borda
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
	candidates          chan *Masquerade
	masquerades         chan *Masquerade
	maxAllowedCachedAge time.Duration
	maxCacheSize        int
	cacheSaveInterval   time.Duration
	toCache             chan *Masquerade
}

// Configure sets the masquerades to use, the trusted root CAs, and the
// cache file for caching masquerades to set up direct domain fronting.
func Configure(pool *x509.CertPool, masquerades map[string][]*Masquerade, cacheFile string) {
	log.Trace("Configuring fronted")
	if masquerades == nil || len(masquerades) == 0 {
		log.Errorf("No masquerades!!")
		return
	}

	CloseCache()

	// Make a copy of the masquerades to avoid data races.
	size := 0
	for _, v := range masquerades {
		size += len(v)
	}

	if size == 0 {
		log.Errorf("No masquerades!!")
		return
	}

	d := &direct{
		tlsConfigs:          make(map[string]*tls.Config),
		certPool:            pool,
		candidates:          make(chan *Masquerade, size),
		masquerades:         make(chan *Masquerade, size),
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
		cacheSaveInterval:   defaultCacheSaveInterval,
		toCache:             make(chan *Masquerade, defaultMaxCacheSize),
	}

	numberToVet := numberToVetInitially
	if cacheFile != "" {
		numberToVet -= d.initCaching(cacheFile)
	}

	d.loadCandidates(masquerades)
	if numberToVet > 0 {
		d.vet(numberToVet)
	} else {
		log.Debug("Not vetting any masquerades because we have enough cached ones")
	}
	_instance.Set(d)
}

func (d *direct) loadCandidates(initial map[string][]*Masquerade) {
	log.Debug("Loading candidates")
	for key, arr := range initial {
		size := len(arr)
		log.Tracef("Adding %d candidates for %v", size, key)
		for i := 0; i < size; i++ {
			// choose index uniformly in [i, n-1]
			r := i + rand.Intn(size-i)
			log.Trace("Adding candidate")
			d.candidates <- arr[r]
		}
	}
}

// Vet vets the specified Masquerade, verifying certificate using the given CertPool
func Vet(m *Masquerade, pool *x509.CertPool) bool {
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
	return postCheck(conn)
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
	conn, masqueradeGood, masqueradesRemain, err := d.dialWith(d.candidates)
	if err != nil {
		return masqueradesRemain
	}
	defer conn.Close()

	if !masqueradeGood(postCheck(conn)) {
		log.Tracef("Unsuccessful vetting with HEAD request, discarding masquerade")
		return masqueradesRemain
	}
	log.Trace("Finished vetting one")
	return false
}

// postCheck does a post with invalid data to verify domain-fronting works
func postCheck(conn net.Conn) bool {
	client := &http.Client{
		Transport: httpTransport(conn, nil),
	}
	return doPostCheck(client)
}

func doPostCheck(client *http.Client) bool {
	req, _ := http.NewRequest(http.MethodPost, testURL, strings.NewReader("{'junk': 'field'}"))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Tracef("Unsuccessful vetting with HEAD request, discarding masquerade")
		return false
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		log.Tracef("Unexpected response status vetting masquerade: %v, %v", resp.StatusCode, resp.Status)
		return false
	}
	return true
}

// NewDirect creates a new http.RoundTripper that does direct domain fronting.
func NewDirect(timeout time.Duration) http.RoundTripper {
	instance, ok := _instance.Get(timeout)
	if !ok {
		panic(fmt.Errorf("No DirectHttpClient available within %v", timeout))
	}
	return instance.(http.RoundTripper)
}

// Do continually retries a given request until it succeeds because some
// fronting providers will return a 403 for some domains.
func (d *direct) RoundTrip(req *http.Request) (*http.Response, error) {
	var body []byte
	var err error
	if req.Body != nil {
		// store body in-memory to be able to replay it if necessary
		body, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("Unable to read request body: %v", err)
		}
	}
	for i := 0; i < maxTries; i++ {
		if body != nil {
			req.Body = ioutil.NopCloser(bytes.NewReader(body))
		}
		conn, masqueradeGood, err := d.dial()
		if err != nil {
			// unable to find good masquerade, fail
			return nil, err
		}
		tr := httpTransport(conn, clientSessionCache)
		resp, err := tr.RoundTrip(req)
		if err != nil {
			log.Errorf("Could not complete request %v", err)
			masqueradeGood(false)
			continue
		}
		resp.Body.Close()
		if resp.StatusCode != http.StatusForbidden {
			masqueradeGood(true)
			return resp, nil
		}
		masqueradeGood(false)
	}

	return nil, errors.New("Could not complete request even with retries")
}

// Dial dials out using a masquerade. If the available masquerade fails, it
// retries with others until it either succeeds or exhausts the available
// masquerades. If successful, it returns a function that the caller can use to
// tell us whether the masquerade is good or not (i.e. if masquerade was good,
// keep it, else vet a new one).
func (d *direct) dial() (net.Conn, func(bool) bool, error) {
	conn, masqueradeGood, _, err := d.dialWith(d.masquerades)
	return conn, masqueradeGood, err
}

func (d *direct) dialWith(in chan *Masquerade) (net.Conn, func(bool) bool, bool, error) {
	retryLater := make([]*Masquerade, 0)
	defer func() {
		for _, m := range retryLater {
			in <- m
		}
	}()

	for {
		var m *Masquerade
		select {
		case m = <-in:
			log.Trace("Got vetted masquerade")
		default:
			log.Trace("No vetted masquerade found, falling back to unvetted candidate")
			select {
			case m = <-d.candidates:
				log.Trace("Got unvetted masquerade")
			default:
				return nil, nil, false, errors.New("Could not dial any masquerade?")
			}
		}

		log.Tracef("Dialing to %v", m)

		// We do the full TLS connection here because in practice the domains at a given IP
		// address can change frequently on CDNs, so the certificate may not match what
		// we expect.
		conn, retriable, err := d.doDial(m)
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
					}
				} else {
					go d.vetOneUntilGood()
				}

				return good
			}
			return conn, masqueradeGood, true, err
		} else if retriable {
			retryLater = append(retryLater, m)
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
			log.Tracef("Not re-adding candidate that failed on error '%v'", err.Error())
			retriable = false
		} else {
			log.Tracef("Unexpected error dialing, keeping masquerade: %v", err)
			retriable = true
		}
	} else {
		log.Tracef("Got successful connection to: %v", m)
		idleTimeout := 70 * time.Second

		log.Trace("Wrapping connecting in idletiming connection")
		conn = idletiming.Conn(conn, idleTimeout, func() {
			log.Tracef("Connection to %v idle for %v, closed", conn.RemoteAddr(), idleTimeout)
		})
	}
	return
}

func (d *direct) dialServerWith(masquerade *Masquerade) (net.Conn, error) {
	tlsConfig := d.tlsConfig(masquerade)
	dialTimeout := 10 * time.Second
	sendServerNameExtension := false

	conn, err := tlsdialer.DialTimeout(
		netx.DialTimeout,
		dialTimeout,
		"tcp",
		masquerade.IpAddress+":443",
		sendServerNameExtension, // SNI or no
		tlsConfig)

	if err != nil && masquerade != nil {
		err = fmt.Errorf("Unable to dial masquerade %s: %s", masquerade.Domain, err)
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
