package fronted

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/getlantern/netx"
	"github.com/getlantern/ops"
	"github.com/getlantern/tlsdialer/v3"
	tls "github.com/refraction-networking/utls"
)

const (
	NumWorkers = 10 // number of worker goroutines for verifying
)

var (
	defaultValidator = NewStatusCodeValidator([]int{403})
)

// CA represents a certificate authority
type CA struct {
	CommonName string
	Cert       string // PEM-encoded
}

// Masquerade contains the data for a single masquerade host, including
// the domain and the root CA.
type Masquerade struct {
	// Domain: the domain to use for domain fronting
	Domain string

	// IpAddress: pre-resolved ip address to use instead of Domain (if
	// available)
	IpAddress string

	// SNI: the SNI to use for this masquerade
	SNI string

	// VerifyHostname is used for checking if the certificate for a given hostname is valid.
	// This is used for verifying if the peer certificate for the hostnames that are being fronted are valid.
	VerifyHostname *string
}

// Create a masquerade interface for easier testing.
type Front interface {
	dial(rootCAs *x509.CertPool, clientHelloID tls.ClientHelloID) (net.Conn, error)

	// Accessor for the domain of the masquerade
	getDomain() string

	// Accessor for the IP address of the masquerade
	getIpAddress() string

	markSucceeded()

	markFailed()

	lastSucceeded() time.Time

	setLastSucceeded(time.Time)

	verifyWithPost(net.Conn, string) bool

	getProviderID() string

	isSucceeding() bool

	markWithResult(bool) bool

	markCacheDirty()
}

type front struct {
	Masquerade
	// lastSucceeded: the most recent time at which this Masquerade succeeded
	LastSucceeded time.Time
	// id of DirectProvider that this masquerade is provided by
	ProviderID string
	mx         sync.RWMutex
	cacheDirty chan interface{}
}

func newFront(m *Masquerade, providerID string, cacheDirty chan interface{}) Front {
	return &front{
		Masquerade:    *m,
		ProviderID:    providerID,
		LastSucceeded: time.Time{},
		cacheDirty:    cacheDirty,
	}
}
func (fr *front) dial(rootCAs *x509.CertPool, clientHelloID tls.ClientHelloID) (net.Conn, error) {
	tlsConfig := &tls.Config{
		ServerName: fr.Domain,
		RootCAs:    rootCAs,
	}
	dialTimeout := 5 * time.Second
	addr := fr.IpAddress
	var sendServerNameExtension bool
	if fr.SNI != "" {
		sendServerNameExtension = true
		tlsConfig.ServerName = fr.SNI
		tlsConfig.InsecureSkipVerify = true
		tlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
			var verifyHostname string
			if fr.VerifyHostname != nil {
				verifyHostname = *fr.VerifyHostname
			}
			return verifyPeerCertificate(rawCerts, rootCAs, verifyHostname)
		}
	}
	dialer := &tlsdialer.Dialer{
		DoDial:         netx.DialTimeout,
		Timeout:        dialTimeout,
		SendServerName: sendServerNameExtension,
		Config:         tlsConfig,
		ClientHelloID:  clientHelloID,
	}
	_, _, err := net.SplitHostPort(addr)
	if err != nil {
		// If there is no port, we default to 443
		addr = net.JoinHostPort(addr, "443")
	}
	return dialer.Dial("tcp", addr)
}

// verifyWithPost does a post with invalid data to verify domain-fronting works
func (fr *front) verifyWithPost(conn net.Conn, testURL string) bool {
	client := &http.Client{
		Transport: connectedConnHTTPTransport(conn, true),
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
		log.Info("Error vetting masquerade", "error", err, "method", method, "url", u)
		return false
	}
	if resp.Body != nil {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}
	if resp.StatusCode != expectedStatus {
		op.Set("response_status", resp.StatusCode)
		op.Set("expected_status", expectedStatus)
		err := fmt.Errorf("Unexpected response status vetting masquerade, expected %d got %d: %v", expectedStatus, resp.StatusCode, resp.Status)
		op.FailIf(err)
		log.Info("Unexpected response status vetting masquerade", "expected", expectedStatus, "statusCode", resp.StatusCode, "status", resp.Status)
		return false
	}
	return true
}

// getDomain implements MasqueradeInterface.
func (fr *front) getDomain() string {
	return fr.Domain
}

// getIpAddress implements MasqueradeInterface.
func (fr *front) getIpAddress() string {
	return fr.IpAddress
}

// getProviderID implements MasqueradeInterface.
func (fr *front) getProviderID() string {
	return fr.ProviderID
}

// MarshalJSON marshals masquerade into json
func (fr *front) MarshalJSON() ([]byte, error) {
	fr.mx.RLock()
	defer fr.mx.RUnlock()
	// Type alias for masquerade so that we don't infinitely recurse when marshaling the struct
	type alias front
	return json.Marshal((*alias)(fr))
}

func (fr *front) lastSucceeded() time.Time {
	fr.mx.RLock()
	defer fr.mx.RUnlock()
	return fr.LastSucceeded
}

func (fr *front) setLastSucceeded(t time.Time) {
	fr.mx.Lock()
	defer fr.mx.Unlock()
	fr.LastSucceeded = t
}

func (fr *front) markSucceeded() {
	fr.mx.Lock()
	defer fr.mx.Unlock()
	fr.LastSucceeded = time.Now()
}

func (fr *front) markFailed() {
	fr.mx.Lock()
	defer fr.mx.Unlock()
	fr.LastSucceeded = time.Time{}
}

func (fr *front) isSucceeding() bool {
	fr.mx.RLock()
	defer fr.mx.RUnlock()
	return fr.LastSucceeded.After(time.Time{})
}

// Make sure that the masquerade struct implements the MasqueradeInterface
var _ Front = (*front)(nil)

// A Direct fronting provider configuration.
type Provider struct {
	// Specific hostname mappings used for this provider.
	// remaps certain requests to provider specific host names.
	HostAliases map[string]string

	// Allow unaliased pass-through of hostnames
	// matching these patterns.
	// eg "*.cloudfront.net" for cloudfront provider
	// would permit all .cloudfront.net domains to
	// pass through without alias. Only suffix
	// patterns and exact matches are supported.
	PassthroughPatterns []string

	// Url used to vet masquerades for this provider
	TestURL     string
	Masquerades []*Masquerade

	// VerifyHostname is used for checking if the certificate for a given hostname is valid.
	// This attribute is only being defined here so it can be sent to the masquerade struct later.
	VerifyHostname *string

	// FrontingSNIs is a map of country code the the SNI config to use for that country.
	FrontingSNIs map[string]*SNIConfig
}

type SNIConfig struct {
	UseArbitrarySNIs bool
	ArbitrarySNIs    []string
}

// Create a Provider with the given details
func NewProvider(hosts map[string]string, testURL string, masquerades []*Masquerade, passthrough []string, frontingSNIs map[string]*SNIConfig, verifyHostname *string, countryCode string) *Provider {
	p := &Provider{
		HostAliases:         make(map[string]string),
		TestURL:             testURL,
		Masquerades:         make([]*Masquerade, 0, len(masquerades)),
		PassthroughPatterns: make([]string, 0, len(passthrough)),
		VerifyHostname:      verifyHostname,
		FrontingSNIs:        frontingSNIs,
	}
	for k, v := range hosts {
		p.HostAliases[strings.ToLower(k)] = v
	}

	var config *SNIConfig
	if countryCode != "" {
		var ok bool
		config, ok = frontingSNIs[countryCode]
		if !ok {
			config = frontingSNIs["default"]
		}
	}
	for _, m := range masquerades {
		sni := generateSNI(config, m)
		p.Masquerades = append(p.Masquerades, &Masquerade{Domain: m.Domain, IpAddress: m.IpAddress, SNI: sni, VerifyHostname: verifyHostname})
	}
	p.PassthroughPatterns = append(p.PassthroughPatterns, passthrough...)
	return p
}

// generateSNI generates a SNI for the given domain and ip address
func generateSNI(config *SNIConfig, m *Masquerade) string {
	if config != nil && m != nil && config.UseArbitrarySNIs && len(config.ArbitrarySNIs) > 0 {
		// Ensure that we use a consistent SNI for a given combination of IP address and SNI set
		hash := sha256.New()
		hash.Write([]byte(m.IpAddress))
		checksum := int(hash.Sum(nil)[0])
		// making sure checksum is positive
		if checksum < 0 {
			checksum = -checksum
		}
		return config.ArbitrarySNIs[checksum%len(config.ArbitrarySNIs)]
	}
	return ""
}

// Lookup the host alias for the given hostname for this provider
func (p *Provider) Lookup(hostname string) string {
	// only consider the host porition if given a port as well.
	if h, _, err := net.SplitHostPort(hostname); err == nil {
		hostname = h
	}
	hostname = strings.ToLower(hostname)
	if alias := p.HostAliases[hostname]; alias != "" {
		return alias
	}

	for _, pt := range p.PassthroughPatterns {
		pt = strings.ToLower(pt)
		if strings.HasPrefix(pt, "*.") && strings.HasSuffix(hostname, pt[1:]) {
			return hostname
		} else if pt == hostname {
			return hostname
		}
	}

	return ""
}

// Validate a fronted response.  Returns an error if the
// response failed to reach the origin, eg if the request
// was rejected by the provider.
func (p *Provider) ValidateResponse(res *http.Response) error {
	return defaultValidator(res)
}

// A validator for fronted responses.  Returns an error if the
// response failed to reach the origin, eg if the request
// was rejected by the provider.
type ResponseValidator func(*http.Response) error

// Create a new ResponseValidator that rejects any response with
// a given list of http status codes.
func NewStatusCodeValidator(reject []int) ResponseValidator {
	bad := make(map[int]bool)
	for _, code := range reject {
		bad[code] = true
	}
	return func(res *http.Response) error {
		if bad[res.StatusCode] {
			return fmt.Errorf("response status %d: %v", res.StatusCode, res.Status)
		}
		return nil
	}
}

type threadSafeFronts struct {
	fronts sortedFronts
	mx     sync.RWMutex
}

func newThreadSafeFronts(size int) *threadSafeFronts {
	return &threadSafeFronts{
		fronts: make(sortedFronts, 0, size),
		mx:     sync.RWMutex{},
	}
}

func (tsf *threadSafeFronts) sortedCopy() sortedFronts {
	tsf.mx.RLock()
	defer tsf.mx.RUnlock()
	c := make(sortedFronts, len(tsf.fronts))
	copy(c, tsf.fronts)
	sort.Sort(c)
	return c
}

func (tsf *threadSafeFronts) addFronts(newFronts ...Front) {
	tsf.mx.Lock()
	defer tsf.mx.Unlock()
	tsf.fronts = append(tsf.fronts, newFronts...)
}

func (tsf *threadSafeFronts) frontSize() int {
	tsf.mx.RLock()
	defer tsf.mx.RUnlock()
	return len(tsf.fronts)
}

func (tsf *threadSafeFronts) frontAt(i int) Front {
	tsf.mx.RLock()
	defer tsf.mx.RUnlock()
	return tsf.fronts[i]
}

// slice of masquerade sorted by last vetted time
type sortedFronts []Front

func (m sortedFronts) Len() int      { return len(m) }
func (m sortedFronts) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m sortedFronts) Less(i, j int) bool {
	if m[i].lastSucceeded().After(m[j].lastSucceeded()) {
		return true
	} else if m[j].lastSucceeded().After(m[i].lastSucceeded()) {
		return false
	} else {
		return m[i].getIpAddress() < m[j].getIpAddress()
	}
}

func (fr *front) markCacheDirty() {
	select {
	case fr.cacheDirty <- nil:
		// okay
	default:
		// already dirty
	}
}

func (fr *front) markWithResult(good bool) bool {
	if good {
		fr.markSucceeded()
	} else {
		fr.markFailed()
	}
	fr.markCacheDirty()
	return good
}
