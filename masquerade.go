package fronted

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
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

type masquerade struct {
	Masquerade
	// lastSucceeded: the most recent time at which this Masquerade succeeded
	LastSucceeded time.Time
	// id of DirectProvider that this masquerade is provided by
	ProviderID string
	mx         sync.RWMutex
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
type sortedMasquerades []*masquerade

func (m sortedMasquerades) Len() int      { return len(m) }
func (m sortedMasquerades) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m sortedMasquerades) Less(i, j int) bool {
	if m[i].lastSucceeded().After(m[j].lastSucceeded()) {
		return true
	} else if m[j].lastSucceeded().After(m[i].lastSucceeded()) {
		return false
	} else {
		return m[i].IpAddress < m[j].IpAddress
	}
}

func (m sortedMasquerades) sortedCopy() sortedMasquerades {
	c := make(sortedMasquerades, len(m))
	copy(c, m)
	sort.Sort(c)
	return c
}
