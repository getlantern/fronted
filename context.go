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

// Configure sets the masquerades to use, the trusted root CAs, and the
// cache file for caching masquerades to set up direct domain fronting
// in the default context.
//
// defaultProviderID is used when a masquerade without a provider is
// encountered (e.g. in a cache file).
func Configure(providers map[string]*Provider, defaultProviderID string, opts ConfigureOptions) {
	if err := DefaultContext.Configure(providers, defaultProviderID, opts); err != nil {
		log.Errorf("Error configuring fronting %s context: %s!!", DefaultContext.name, err)
	}
}

// NewDirect creates a new http.RoundTripper that does direct domain fronting
// using the default context. If it can't obtain a working masquerade within
// the given timeout, it will return nil/false.
func NewDirect(timeout time.Duration) (http.RoundTripper, bool) {
	return DefaultContext.NewDirect(timeout)
}

// CloseCache closes any existing cache file in the default context
func CloseCache() {
	DefaultContext.CloseCache()
}

func NewFrontingContext(name string) *FrontingContext {
	return &FrontingContext{
		name:     name,
		instance: eventual.NewValue(),
	}
}

type FrontingContext struct {
	name     string
	instance eventual.Value
}

// Configure sets the masquerades to use, the trusted root CAs, and the
// cache file for caching masquerades to set up direct domain fronting.
//
// defaultProviderID is used when a masquerade without a provider is
// encountered (e.g. in a cache file).
func (fctx *FrontingContext) Configure(providers map[string]*Provider, defaultProviderID string, opts ConfigureOptions) error {
	log.Tracef("Configuring fronted %s context", fctx.name)

	if providers == nil || len(providers) == 0 {
		return fmt.Errorf("no fronted providers for %s context", fctx.name)
	}

	_existing, ok := fctx.instance.Get(0)
	if ok && _existing != nil {
		existing := _existing.(*direct)
		log.Debugf("Closing cache from existing instance for %s context", fctx.name)
		existing.closeCache()
	}

	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}
	if size == 0 {
		return fmt.Errorf("no masquerades for %s context", fctx.name)
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
		numberToVet -= d.initCaching(opts.CacheFile)
	}

	d.loadCandidates(d.providers)
	if numberToVet > 0 {
		d.vet(numberToVet)
	} else {
		log.Debugf("Not vetting any masquerades for %s context because we have enough cached ones", fctx.name)
		d.signalReady()
	}
	fctx.instance.Set(d)
	return nil
}

// NewDirect creates a new http.RoundTripper that does direct domain fronting.
// If it can't obtain a working masquerade within the given timeout, it will
// return nil/false.
func (fctx *FrontingContext) NewDirect(timeout time.Duration) (http.RoundTripper, bool) {
	start := time.Now()
	instance, ok := fctx.instance.Get(timeout)
	if !ok {
		log.Errorf("No DirectHttpClient available within %v for context %s", timeout, fctx.name)
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

// CloseCache closes any existing cache file in the default contexxt.
func (fctx *FrontingContext) CloseCache() {
	_existing, ok := fctx.instance.Get(0)
	if ok && _existing != nil {
		existing := _existing.(*direct)
		log.Debugf("Closing cache from existing instance in %s context", fctx.name)
		existing.closeCache()
	}
}
