package fronted

import (
	"context"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/getlantern/eventual"
	"github.com/getlantern/netx"
	tls "github.com/refraction-networking/utls"
)

var (
	DefaultContext = NewFrontingContext("default")
)

// ConfigureOptions is used in Configure and FrontingContext.Configure.
type ConfigureOptions struct {
	// CertPool sets the root CAs used to verify server certificates. If nil, the host's root CA set
	// will be used.
	CertPool *x509.CertPool

	// CacheFile, if provided, will be used to cache providers.
	CacheFile string

	// ClientHelloID, if provided, specifies the ID of a ClientHello to mimic. See
	// https://pkg.go.dev/github.com/refraction-networking/utls?tab=doc#pkg-variables
	ClientHelloID tls.ClientHelloID

	// DialTransport is used to establish the transport connection to the masquerade. This will
	// almost certainly be a TCP connection. If nil, getlantern/netx.DialContext will be used.
	DialTransport func(ctx context.Context, network, address string) (net.Conn, error)
}

// Configure sets the masquerades to use in the default context. The
// defaultProviderID is used when a masquerade without a provider is
// encountered (e.g. in a cache file).
func Configure(providers map[string]*Provider, defaultProviderID string) {
	if err := DefaultContext.Configure(providers, defaultProviderID); err != nil {
		log.Errorf("Error configuring fronting %s context: %s!!", DefaultContext.name, err)
	}
}

// NewDirect creates a new http.RoundTripper that does direct domain fronting
// using the default context. If it can't obtain a working masquerade within
// the given timeout, it will return nil/false.
func NewDirect(timeout time.Duration, opts DirectOptions) (http.RoundTripper, bool) {
	return DefaultContext.NewDirect(timeout, opts)
}

// CloseCache closes any existing cache file in the default context
func CloseCache() {
	DefaultContext.CloseCache()
}

func NewFrontingContext(name string) *FrontingContext {
	return &FrontingContext{
		name:              name,
		instance:          eventual.NewValue(),
		providers:         eventual.NewValue(),
		defaultProviderID: eventual.NewValue(),
	}
}

type FrontingContext struct {
	name     string
	instance eventual.Value // TODO: delete

	// Set by Configure.
	providers         eventual.Value // map[string]*Provider
	defaultProviderID eventual.Value // string
}

// Configure sets the masquerades to use. ThedefaultProviderID is used when a
// masquerade without a provider is encountered (e.g. in a cache file).
func (fctx *FrontingContext) Configure(providers map[string]*Provider, defaultProviderID string) error {
	log.Tracef("Configuring fronted %s context", fctx.name)

	if providers == nil || len(providers) == 0 {
		return fmt.Errorf("no fronted providers for %s context", fctx.name)
	}

	// TODO: handle cache closing
	// _existing, ok := fctx.instance.Get(0)
	// if ok && _existing != nil {
	// 	existing := _existing.(*direct)
	// 	log.Debugf("Closing cache from existing instance for %s context", fctx.name)
	// 	existing.closeCache()
	// }

	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}
	if size == 0 {
		return fmt.Errorf("no masquerades for %s context", fctx.name)
	}

	fctx.providers.Set(providers)
	fctx.defaultProviderID.Set(defaultProviderID)
	return nil
}

// DirectOptions defines optional paramaters for NewDirect and FrontingContext.NewDirect.
type DirectOptions struct {
	// CertPool sets the root CAs used to verify server certificates. If nil, the host's root CA set
	// will be used.
	CertPool *x509.CertPool

	// CacheFile, if provided, will be used to cache providers.
	CacheFile string

	// ClientHelloID, if provided, specifies the ID of a ClientHello to mimic. See
	// https://pkg.go.dev/github.com/refraction-networking/utls?tab=doc#pkg-variables
	ClientHelloID tls.ClientHelloID

	// DialTransport is used to establish the transport connection to the masquerade. This will
	// almost certainly be a TCP connection. If nil, getlantern/netx.DialContext will be used.
	DialTransport func(ctx context.Context, network, address string) (net.Conn, error)
}

// NewDirect creates a new http.RoundTripper that does direct domain fronting.
// If it can't obtain a working masquerade within the given timeout, it will
// return nil/false.
func (fctx *FrontingContext) NewDirect(timeout time.Duration, opts DirectOptions) (http.RoundTripper, bool) {
	start := time.Now()

	providersCh, defaultProviderIDCh := make(chan interface{}), make(chan interface{})
	go func() { v, _ := fctx.providers.Get(timeout); providersCh <- v }()
	go func() { v, _ := fctx.defaultProviderID.Get(timeout); defaultProviderIDCh <- v }()
	providersI, defaultProviderIDI := <-providersCh, <-defaultProviderIDCh
	if providersI == nil || defaultProviderIDI == nil {
		log.Errorf("configuration values not available within %v for context %s", timeout, fctx.name)
		return nil, false
	}
	providers := providersI.(map[string]*Provider)
	defaultProviderID := defaultProviderIDI.(string)

	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}
	if opts.DialTransport == nil {
		opts.DialTransport = netx.DialContext
	}
	d := &direct{
		certPool:            opts.CertPool,
		candidates:          make(chan masquerade, size),
		masquerades:         make(chan masquerade, size),
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
		cacheSaveInterval:   defaultCacheSaveInterval,
		toCache:             make(chan masquerade, defaultMaxCacheSize),
		defaultProviderID:   defaultProviderID,
		providers:           make(map[string]*Provider),
		ready:               make(chan struct{}),
		dialTransport:       opts.DialTransport,
		clientHelloID:       opts.ClientHelloID,
	}
	// copy providers
	for k, p := range providers {
		d.providers[k] = NewProvider(p.HostAliases, p.TestURL, p.Masquerades, p.Validator, p.PassthroughPatterns)
	}
	numberToVet := numberToVetInitially
	if opts.CacheFile != "" {
		// TODO: deal with shared cache file
		numberToVet -= d.initCaching(opts.CacheFile)
	}

	d.loadCandidates(d.providers)
	if numberToVet > 0 {
		d.vet(numberToVet)
	} else {
		log.Debugf("Not vetting any masquerades for %s context because we have enough cached ones", fctx.name)
		d.signalReady()
	}

	select {
	case <-d.ready:
		return d, true
	case <-time.After(timeout - time.Since(start)):
		return nil, false
	}
}

// CloseCache closes any existing cache file in the default contexxt.
func (fctx *FrontingContext) CloseCache() {
	// TODO: track all directs and close all caches here

	_existing, ok := fctx.instance.Get(0)
	if ok && _existing != nil {
		existing := _existing.(*direct)
		log.Debugf("Closing cache from existing instance in %s context", fctx.name)
		existing.closeCache()
	}
}
