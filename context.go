package fronted

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/getlantern/eventual"
	tls "github.com/refraction-networking/utls"
)

var (
	DefaultContext = NewFrontingContext("default")
)

// ErrorTimeout is returned when an operation times out.
type ErrorTimeout struct {
	msg string
}

func (err ErrorTimeout) Error() string {
	return err.msg
}

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

// NewDirect creates a new http.RoundTripper that does direct domain fronting.
// The default context must be configured to create a RoundTripper.
func NewDirect(ctx context.Context, opts DirectOptions) (http.RoundTripper, error) {
	return DefaultContext.NewDirect(ctx, opts)
}

func NewFrontingContext(name string) *FrontingContext {
	return &FrontingContext{
		name, map[string]*masqueradeCache{}, sync.Mutex{}, eventual.NewValue(), eventual.NewValue()}
}

type FrontingContext struct {
	name       string
	caches     map[string]*masqueradeCache
	cachesLock sync.Mutex

	// Set by Configure.
	providers         eventual.Value // map[string]*Provider
	defaultProviderID eventual.Value // string
}

// Configure sets the masquerades to use. The defaultProviderID is used when a
// masquerade without a provider is encountered (e.g. in a cache file).
func (fctx *FrontingContext) Configure(providers map[string]*Provider, defaultProviderID string) error {
	log.Tracef("Configuring fronted %s context", fctx.name)

	// Sanity check inputs as NewDirect expects valid values.
	if providers == nil || len(providers) == 0 {
		return errors.New("providers are required")
	}
	if defaultProviderID == "" {
		return errors.New("default provider ID is required")
	}
	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}
	if size == 0 {
		return errors.New("no masquerades in providers")
	}

	fctx.providers.Set(providers)
	fctx.defaultProviderID.Set(defaultProviderID)
	return nil
}

// NewDirect creates a new http.RoundTripper that does direct domain fronting.
// The fronting context must be configured to create a RoundTripper.
func (fctx *FrontingContext) NewDirect(ctx context.Context, opts DirectOptions) (http.RoundTripper, error) {
	// Note: eventual.Value.Get(-1) will wait forever. If no deadline is set, this is what we want.
	timeout := time.Duration(-1)
	if deadline, ok := ctx.Deadline(); ok {
		timeout = time.Until(deadline)
	}
	providersCh, defaultProviderIDCh := make(chan interface{}), make(chan interface{})
	go func() { v, _ := fctx.providers.Get(timeout); providersCh <- v }()
	go func() { v, _ := fctx.defaultProviderID.Get(timeout); defaultProviderIDCh <- v }()

	var (
		providers         map[string]*Provider
		defaultProviderID string
	)
	for providers == nil || defaultProviderID == "" {
		select {
		case _providers := <-providersCh:
			providers = _providers.(map[string]*Provider)
		case _defaultProviderID := <-defaultProviderIDCh:
			defaultProviderID = _defaultProviderID.(string)
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}

	var (
		cache    *masqueradeCache
		newCache bool
		err      error
	)
	if opts.CacheFile != "" {
		cache, newCache, err = fctx.getCache(opts.CacheFile)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize cache file: %w", err)
		}
	}

	d, err := newDirect(ctx, providers, defaultProviderID, cache, opts)
	if err != nil && newCache {
		fctx.closeCache(opts.CacheFile)
	}
	return d, err
}

// Close the context and any associated resources. RoundTrippers created via NewDirect will continue
// to operate, but will no longer cache masquerades.
//
// Always returns nil.
func (fctx *FrontingContext) Close() error {
	// Note: we return an error in the signature to (a) implement io.Closer and
	// (b) allow us to return an error in the future if the need arises.
	// If we do start returning non-nil errors, the doc should be updated.
	fctx.cachesLock.Lock()
	defer fctx.cachesLock.Unlock()
	for _, c := range fctx.caches {
		c.close()
	}
	fctx.caches = map[string]*masqueradeCache{}
	return nil
}

func (fctx *FrontingContext) getCache(filename string) (c *masqueradeCache, isNew bool, err error) {
	fctx.cachesLock.Lock()
	defer fctx.cachesLock.Unlock()
	if c, ok := fctx.caches[filename]; ok {
		return c, false, nil
	}
	c, err = newMasqueradeCache(
		filename, defaultMaxCacheSize, defaultMaxAllowedCachedAge, defaultCacheSaveInterval)
	if err != nil {
		return nil, false, err
	}
	fctx.caches[filename] = c
	return c, true, nil
}

func (fctx *FrontingContext) closeCache(filename string) {
	fctx.cachesLock.Lock()
	fctx.caches[filename].close()
	delete(fctx.caches, filename)
	fctx.cachesLock.Unlock()
}
