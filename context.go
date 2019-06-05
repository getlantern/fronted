package fronted

import "crypto/x509"
import "fmt"
import "net/http"
import "time"
import "sync"

import "github.com/getlantern/eventual"

var (
	contextsMx       sync.Mutex
	contexts         = make(map[string]*frontingContext)
	defaultContextID = "default"
)

// Configure sets the masquerades to use, the trusted root CAs, and the
// cache file for caching masquerades to set up direct domain fronting
// in the default context.
//
// defaultProviderID is used when a masquerade without a provider is
// encountered (eg in a cache file)
func Configure(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string, cacheFile string) {
	ConfigureContext(defaultContextID, pool, providers, defaultProviderID, cacheFile)
}

// ConfigureContext establishes a new independent fronting
// configuration for each id.  Masquerades are chosen, vetted and
// used separately from the default configuration and other
// fronting contexts.
func ConfigureContext(id string, pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string, cacheFile string) {
	fctx := getOrCreateContext(id)
	if err := fctx.Configure(pool, providers, defaultProviderID, cacheFile); err != nil {
		log.Errorf("Error configuring fronting %s context: %s!!", id, err)
	}
}

// NewDirect creates a new http.RoundTripper that does direct domain fronting
// using the default context. If it can't obtain a working masquerade within
// the given timeout, it will return nil/false.
func NewDirect(timeout time.Duration) (http.RoundTripper, bool) {
	return NewDirectContext(defaultContextID, timeout)
}

func NewDirectContext(id string, timeout time.Duration) (http.RoundTripper, bool) {
	return getOrCreateContext(id).NewDirect(timeout)
}

// CloseCache closes any existing cache file.
func CloseCache() {
	contextsMx.Lock()
	ids := make([]string, 0, len(contexts))
	for id := range contexts {
		ids = append(ids, id)
	}
	contextsMx.Unlock()

	for _, id := range ids {
		CloseCacheContext(id)
	}
}

func CloseCacheContext(id string) {
	getOrCreateContext(id).CloseCache()
}

func getOrCreateContext(id string) *frontingContext {
	contextsMx.Lock()
	fctx := contexts[id]
	if fctx == nil {
		fctx = &frontingContext{
			id:       id,
			instance: eventual.NewValue(),
		}
		contexts[id] = fctx
	}
	contextsMx.Unlock()
	return fctx
}

type frontingContext struct {
	id       string
	instance eventual.Value
}

// Configure sets the masquerades to use, the trusted root CAs, and the
// cache file for caching masquerades to set up direct domain fronting.
// defaultProviderID is used when a masquerade without a provider is
// encountered (eg in a cache file)
func (fctx *frontingContext) Configure(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string, cacheFile string) error {
	log.Tracef("Configuring fronted %s context", fctx.id)

	if providers == nil || len(providers) == 0 {
		return fmt.Errorf("No fronted providers for %s context.", fctx.id)
	}

	_existing, ok := fctx.instance.Get(0)
	if ok && _existing != nil {
		existing := _existing.(*direct)
		log.Debugf("Closing cache from existing instance for %s context", fctx.id)
		existing.closeCache()
	}

	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}

	if size == 0 {
		return fmt.Errorf("No masquerades for %s context.", fctx.id)
	}

	d := &direct{
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
		d.providers[k] = NewProvider(p.HostAliases, p.TestURL, p.Masquerades, p.Validator, p.PassthroughDomains)
	}

	numberToVet := numberToVetInitially
	if cacheFile != "" {
		numberToVet -= d.initCaching(cacheFile)
	}

	d.loadCandidates(d.providers)
	if numberToVet > 0 {
		d.vet(numberToVet)
	} else {
		log.Debugf("Not vetting any masquerades for %s context because we have enough cached ones", fctx.id)
		d.signalReady()
	}
	fctx.instance.Set(d)
	return nil
}

// NewDirect creates a new http.RoundTripper that does direct domain fronting.
// If it can't obtain a working masquerade within the given timeout, it will
// return nil/false.
func (fctx *frontingContext) NewDirect(timeout time.Duration) (http.RoundTripper, bool) {
	start := time.Now()
	instance, ok := fctx.instance.Get(timeout)
	if !ok {
		log.Errorf("No DirectHttpClient available within %v for context %s", timeout, fctx.id)
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
func (fctx *frontingContext) CloseCache() {
	_existing, ok := fctx.instance.Get(0)
	if ok && _existing != nil {
		existing := _existing.(*direct)
		log.Debugf("Closing cache from existing instance in %s context", fctx.id)
		existing.closeCache()
	}
}
