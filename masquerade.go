package fronted

import (
	"strings"
	"time"
)

const (
	NumWorkers = 10 // number of worker goroutines for verifying
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
	// Url used to vet masquerades for this provider
	TestURL     string
	Masquerades []*Masquerade
}

// Create a DirectProvider with the given details
func NewProvider(hosts map[string]string, testURL string, masquerades []*Masquerade) *Provider {
	d := &Provider{
		HostAliases: make(map[string]string),
		TestURL:     testURL,
		Masquerades: make([]*Masquerade, 0, len(masquerades)),
	}
	for k, v := range hosts {
		d.HostAliases[strings.ToLower(k)] = v
	}
	for _, m := range masquerades {
		d.Masquerades = append(d.Masquerades, &Masquerade{Domain: m.Domain, IpAddress: m.IpAddress})
	}
	return d
}

func (p *Provider) Lookup(hostname string) string {
	return p.HostAliases[strings.ToLower(hostname)]
}
