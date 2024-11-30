package fronted

import (
	"context"
	"crypto/x509"
	"fmt"
	"net/http"
	"time"

	tls "github.com/refraction-networking/utls"

	"github.com/getlantern/eventual/v2"
)

// Create an interface for the fronting context
type Fronted interface {
	UpdateConfig(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string)
	NewRoundTripper(timeout time.Duration) (http.RoundTripper, error)
	Close()
}

var defaultContext = newFrontingContext("default")

// Make sure that the default context is a Fronting
var _ Fronted = defaultContext

// Configure sets the masquerades to use, the trusted root CAs, and the
// cache file for caching masquerades to set up direct domain fronting
// in the default context.
//
// defaultProviderID is used when a masquerade without a provider is
// encountered (eg in a cache file)
func NewFronted(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string, cacheFile string) (Fronted, error) {
	if err := defaultContext.configure(pool, providers, defaultProviderID, cacheFile); err != nil {
		return nil, log.Errorf("Error configuring fronting %s context: %s!!", defaultContext.name, err)
	}
	return defaultContext, nil
}

func newFrontingContext(name string) *frontingContext {
	return &frontingContext{
		name:             name,
		instance:         eventual.NewValue(),
		connectingFronts: newConnectingFronts(),
	}
}

type frontingContext struct {
	name             string
	instance         eventual.Value
	fronted          *fronted
	connectingFronts *connectingFronts
}

// UpdateConfig updates the configuration of the fronting context
func (fctx *frontingContext) UpdateConfig(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string) {
	fctx.fronted.updateConfig(pool, providers, defaultProviderID)
}

// configure sets the masquerades to use, the trusted root CAs, and the
// cache file for caching masquerades to set up direct domain fronting.
// defaultProviderID is used when a masquerade without a provider is
// encountered (eg in a cache file)
func (fctx *frontingContext) configure(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string, cacheFile string) error {
	return fctx.configureWithHello(pool, providers, defaultProviderID, cacheFile, tls.ClientHelloID{})
}

func (fctx *frontingContext) configureWithHello(pool *x509.CertPool, providers map[string]*Provider, defaultProviderID string, cacheFile string, clientHelloID tls.ClientHelloID) error {
	log.Debugf("Configuring fronted %s context", fctx.name)

	if len(providers) == 0 {
		return fmt.Errorf("no fronted providers for %s context", fctx.name)
	}

	if _existing, err := fctx.instance.Get(eventual.DontWait); err != nil {
		log.Debugf("No existing instance for %s context: %s", fctx.name, err)
	} else if _existing != nil {
		existing := _existing.(*fronted)
		log.Debugf("Closing cache from existing instance for %s context", fctx.name)
		existing.closeCache()
	}

	var err error
	if fctx.fronted, err = newFronted(pool, providers, defaultProviderID, cacheFile, clientHelloID, func(f *fronted) {
		log.Debug("Setting fronted instance")
		fctx.instance.Set(f)
	}, fctx.connectingFronts); err != nil {
		return err
	}
	return nil
}

// NewFronted creates a new http.RoundTripper that does direct domain fronting.
// If the context isn't configured within the given timeout, this method
// returns nil, false.
func (fctx *frontingContext) NewRoundTripper(timeout time.Duration) (http.RoundTripper, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	instance, err := fctx.instance.Get(ctx)
	if err != nil {
		return nil, log.Errorf("No DirectHttpClient available within %v for context %s with error %v", timeout, fctx.name, err)
	} else {
		log.Debugf("DirectHttpClient available for context %s after %v with duration %v", fctx.name, time.Since(start), timeout)
	}
	return instance.(http.RoundTripper), nil
}

// Close closes any existing cache file in the default contexxt.
func (fctx *frontingContext) Close() {
	_existing, err := fctx.instance.Get(eventual.DontWait)
	if err != nil {
		log.Errorf("Error getting existing instance for %s context: %s", fctx.name, err)
		return
	}
	if _existing != nil {
		existing := _existing.(*fronted)
		log.Debugf("Closing cache from existing instance in %s context", fctx.name)
		existing.closeCache()
	}
}
