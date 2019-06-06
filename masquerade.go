package fronted

import (
	"fmt"
	"net"
	"net/http"
	"strings"
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
}

type masquerade struct {
	Masquerade
	// lastVetted: the most recent time at which this Masquerade was vetted
	LastVetted time.Time
	// id of DirectProvider that this masquerade is provided by
	ProviderID string
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
	// Optional response validator used to determine whether
	// fronting succeeded for this provider. If the validator
	// detects a failure for a given masquerade, it is discarded.
	// The default validator is used if nil.
	Validator ResponseValidator
}

// Create a Provider with the given details
func NewProvider(hosts map[string]string, testURL string, masquerades []*Masquerade, validator ResponseValidator, passthrough []string) *Provider {
	d := &Provider{
		HostAliases:         make(map[string]string),
		TestURL:             testURL,
		Masquerades:         make([]*Masquerade, 0, len(masquerades)),
		Validator:           validator,
		PassthroughPatterns: make([]string, 0, len(passthrough)),
	}
	for k, v := range hosts {
		d.HostAliases[strings.ToLower(k)] = v
	}
	for _, m := range masquerades {
		d.Masquerades = append(d.Masquerades, &Masquerade{Domain: m.Domain, IpAddress: m.IpAddress})
	}
	for _, pt := range passthrough {
		d.PassthroughPatterns = append(d.PassthroughPatterns, pt)
	}
	return d
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
