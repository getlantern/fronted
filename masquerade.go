package fronted

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"errors"
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
type MasqueradeInterface interface {
	dial(rootCAs *x509.CertPool, clientHelloID tls.ClientHelloID) (net.Conn, error)

	// Accessor for the domain of the masquerade
	getDomain() string

	//Accessor for the IP address of the masquerade
	getIpAddress() string

	markSucceeded()

	markFailed()

	lastSucceeded() time.Time

	setLastSucceeded(time.Time)

	postCheck(net.Conn, string) bool

	getProviderID() string
}

type masquerade struct {
	Masquerade
	// lastSucceeded: the most recent time at which this Masquerade succeeded
	LastSucceeded time.Time
	// id of DirectProvider that this masquerade is provided by
	ProviderID string
	mx         sync.RWMutex
}

func (m *masquerade) dial(rootCAs *x509.CertPool, clientHelloID tls.ClientHelloID) (net.Conn, error) {
	tlsConfig := &tls.Config{
		ServerName: m.Domain,
		RootCAs:    rootCAs,
	}
	dialTimeout := 5 * time.Second
	addr := m.IpAddress
	var sendServerNameExtension bool
	if m.SNI != "" {
		sendServerNameExtension = true
		tlsConfig.ServerName = m.SNI
		tlsConfig.InsecureSkipVerify = true
		tlsConfig.VerifyPeerCertificate = func(rawCerts [][]byte, _ [][]*x509.Certificate) error {
			var verifyHostname string
			if m.VerifyHostname != nil {
				verifyHostname = *m.VerifyHostname
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

// postCheck does a post with invalid data to verify domain-fronting works
func (m *masquerade) postCheck(conn net.Conn, testURL string) bool {
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
		op.FailIf(errors.New(msg))
		log.Debug(msg)
		return false
	}
	return true
}

// getDomain implements MasqueradeInterface.
func (m *masquerade) getDomain() string {
	return m.Domain
}

// getIpAddress implements MasqueradeInterface.
func (m *masquerade) getIpAddress() string {
	return m.IpAddress
}

// getProviderID implements MasqueradeInterface.
func (m *masquerade) getProviderID() string {
	return m.ProviderID
}

// MarshalJSON marshals masquerade into json
func (m *masquerade) MarshalJSON() ([]byte, error) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	// Type alias for masquerade so that we don't infinitely recurse when marshaling the struct
	type alias masquerade
	return json.Marshal((*alias)(m))
}

func (m *masquerade) lastSucceeded() time.Time {
	m.mx.RLock()
	defer m.mx.RUnlock()
	return m.LastSucceeded
}

func (m *masquerade) setLastSucceeded(t time.Time) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.LastSucceeded = t
}

func (m *masquerade) markSucceeded() {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.LastSucceeded = time.Now()
}

func (m *masquerade) markFailed() {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.LastSucceeded = time.Time{}
}

// Make sure that the masquerade struct implements the MasqueradeInterface
var _ MasqueradeInterface = (*masquerade)(nil)

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

	// SNIConfig has the configuration that sets if we should or not use arbitrary SNIs
	// and which SNIs to use.
	SNIConfig *SNIConfig

	// Optional response validator used to determine whether
	// fronting succeeded for this provider. If the validator
	// detects a failure for a given masquerade, it is discarded.
	// The default validator is used if nil.
	Validator ResponseValidator

	// VerifyHostname is used for checking if the certificate for a given hostname is valid.
	// This attribute is only being defined here so it can be sent to the masquerade struct later.
	VerifyHostname *string
}

type SNIConfig struct {
	UseArbitrarySNIs bool
	ArbitrarySNIs    []string
}

// Create a Provider with the given details
func NewProvider(hosts map[string]string, testURL string, masquerades []*Masquerade, validator ResponseValidator, passthrough []string, sniConfig *SNIConfig, verifyHostname *string) *Provider {
	d := &Provider{
		HostAliases:         make(map[string]string),
		TestURL:             testURL,
		Masquerades:         make([]*Masquerade, 0, len(masquerades)),
		Validator:           validator,
		PassthroughPatterns: make([]string, 0, len(passthrough)),
		SNIConfig:           sniConfig,
	}
	for k, v := range hosts {
		d.HostAliases[strings.ToLower(k)] = v
	}

	for _, m := range masquerades {
		sni := generateSNI(d.SNIConfig, m)
		d.Masquerades = append(d.Masquerades, &Masquerade{Domain: m.Domain, IpAddress: m.IpAddress, SNI: sni, VerifyHostname: verifyHostname})
	}
	d.PassthroughPatterns = append(d.PassthroughPatterns, passthrough...)
	return d
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
	if p.Validator != nil {
		return p.Validator(res)
	} else {
		return defaultValidator(res)
	}
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

// slice of masquerade sorted by last vetted time
type sortedMasquerades []MasqueradeInterface

func (m sortedMasquerades) Len() int      { return len(m) }
func (m sortedMasquerades) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m sortedMasquerades) Less(i, j int) bool {
	if m[i].lastSucceeded().After(m[j].lastSucceeded()) {
		return true
	} else if m[j].lastSucceeded().After(m[i].lastSucceeded()) {
		return false
	} else {
		return m[i].getIpAddress() < m[j].getIpAddress()
	}
}

func (m sortedMasquerades) sortedCopy() sortedMasquerades {
	c := make(sortedMasquerades, len(m))
	copy(c, m)
	sort.Sort(c)
	return c
}
