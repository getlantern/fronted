package fronted

import (
	"bytes"
	"compress/gzip"
	"context"
	"crypto/x509"
	"embed"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/goccy/go-yaml"
	tls "github.com/refraction-networking/utls"

	"github.com/getlantern/keepcurrent"
	"github.com/getlantern/ops"

	"github.com/alitto/pond/v2"
)

const (
	defaultMaxAllowedCachedAge = 24 * time.Hour
	defaultMaxCacheSize        = 1000
	defaultCacheSaveInterval   = 5 * time.Second
	maxTries                   = 6
)

var (
	log                      = slog.Default().With(slog.String("module", "fronted"))
	defaultFrontedProviderID = "cloudfront"
)

// fronted identifies working IP address/domain pairings for domain fronting and is
// an implementation of http.RoundTripper for the convenience of callers.
type fronted struct {
	certPool            atomic.Value
	fronts              *threadSafeFronts
	maxAllowedCachedAge time.Duration
	maxCacheSize        int
	cacheFile           string
	cacheSaveInterval   time.Duration
	cacheDirty          chan interface{}
	cacheClosed         chan interface{}
	closeCacheOnce      sync.Once
	defaultProviderID   string
	providers           map[string]*Provider
	clientHelloID       tls.ClientHelloID

	providersMu   sync.RWMutex
	frontsMu      sync.RWMutex
	stopCh        chan interface{}
	crawlOnce     sync.Once
	stopped       atomic.Bool
	countryCode   string
	httpClient    *http.Client
	configURL     string
	frontsCh      chan Front
	panicListener func(string)
}

// Interface for sending HTTP traffic over domain fronting.
type Fronted interface {
	NewConnectedRoundTripper(ctx context.Context, addr string) (http.RoundTripper, error)

	// onNewFrontsConfig updates the set of domain fronts to try from a YAML configuration.
	onNewFrontsConfig(yml []byte)

	// onNewFronts updates the set of domain fronts to try.
	onNewFronts(pool *x509.CertPool, providers map[string]*Provider)

	// Close closes any resources, such as goroutines that are testing fronts.
	Close()
}

//go:generate ./download_config.sh
//go:embed fronted.yaml.gz
var embedFS embed.FS

// Option is a functional option type that allows us to configure the fronted client.
type Option func(*fronted)

// NewFronted creates a new Fronted instance with the given cache file.
// At this point it does not have the actual IPs, domains, etc of the fronts to try.
// defaultProviderID is used when a front without a provider is encountered (eg in a cache file)
func NewFronted(options ...Option) Fronted {
	log.Debug("Creating new fronted")

	f := &fronted{
		certPool:            atomic.Value{},
		fronts:              newThreadSafeFronts(0),
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
		cacheSaveInterval:   defaultCacheSaveInterval,
		cacheDirty:          make(chan any, 1),
		cacheClosed:         make(chan any),
		providers:           make(map[string]*Provider),
		// We can and should update this as new ClientHellos become available in utls.
		clientHelloID:     tls.HelloChrome_131,
		stopCh:            make(chan any, 10),
		defaultProviderID: defaultFrontedProviderID,
		httpClient:        http.DefaultClient,
		configURL:         "",
		frontsCh:          make(chan Front, 4000),
	}

	for _, opt := range options {
		opt(f)
	}
	if f.cacheFile == "" {
		f.cacheFile = defaultCacheFilePath()
	}

	f.initCaching(f.cacheFile)
	f.readFrontsFromEmbeddedConfig()
	f.keepCurrent()

	return f
}

// WithHTTPClient sets the HTTP client to use for fetching the fronted configuration. For example, the client
// could be censorship-resistant in some way.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(f *fronted) {
		f.httpClient = httpClient
	}
}

// WithCacheFile sets the file to use for caching domains that have successfully connected.
func WithCacheFile(file string) Option {
	return func(f *fronted) {
		f.cacheFile = file
	}
}

// WithCountryCode sets the country code to use for fronting, which is particularly relevant for the
// SNI to use when connecting to the fronting domain.
func WithCountryCode(cc string) Option {
	return func(f *fronted) {
		f.countryCode = cc
	}
}

// WithConfigURL sets the URL from which to continually fetch updated domain fronting configurations.
func WithConfigURL(configURL string) Option {
	return func(f *fronted) {
		f.configURL = configURL
	}
}

// WithPanicListener sets a listener for panics that occur in the fronted.
func WithPanicListener(panicListener func(string)) Option {
	return func(f *fronted) {
		f.panicListener = panicListener
	}
}

// SetLogger sets the logger to use by fronted.
func SetLogger(logger *slog.Logger) {
	log = logger.With(slog.String("module", "fronted"))
}

func (f *fronted) paniced(msg string, r any, stack []byte) {
	var pcs [1]uintptr
	// skip [runtime.Callers, this function]
	runtime.Callers(2, pcs[:])
	fs := runtime.CallersFrames(pcs[:])
	fr, _ := fs.Next()

	log.Error(
		msg,
		slog.Any("panic(r)", r),
		slog.String("function", fr.Function),
		slog.String("stack", string(stack)),
	)
	if f.panicListener != nil {
		f.panicListener(fmt.Sprintf("%s. panic(r)=%v", msg, r))
	}
}

func defaultCacheFilePath() string {
	if dir, err := os.UserConfigDir(); err != nil {
		log.Warn("Unable to get user config dir", "error", err)
		return mkdirall(os.TempDir(), "domainfronting", "fronted_cache.json")
	} else {
		return mkdirall(dir, "domainfronting", "fronted_cache.json")
	}
}

func mkdirall(base, path, fileName string) string {
	path = filepath.Join(base, path)
	if err := os.MkdirAll(path, 0o700); err != nil {
		log.Warn("Unable to create directory for fronted cache", "path", path, "error", err)
	}
	return filepath.Join(path, fileName)
}

// keepCurrent fetches the fronted configuration from the given URL and keeps it up
// to date by fetching it periodically.
func (f *fronted) keepCurrent() {
	if f.configURL == "" {
		log.Debug("No config URL provided -- not updating fronting configuration")
		return
	}

	log.Debug("Updating fronted configuration", "url", f.configURL)
	source := keepcurrent.FromWebWithClient(f.configURL, f.httpClient)
	chDB := make(chan []byte)
	dest := keepcurrent.ToChannel(chDB)

	runner := keepcurrent.NewWithValidator(
		f.validator(),
		source,
		dest,
	)

	go func() {
		// Recover from panics and log them
		defer func() {
			if r := recover(); r != nil {
				f.paniced("PANIC waiting for fronts", r, debug.Stack())
			}
		}()
		for data := range chDB {
			log.Debug("Received new fronted configuration")
			f.onNewFrontsConfig(data)
		}
	}()

	runner.Start(12 * time.Hour)
}

func (f *fronted) validator() func([]byte) error {
	return func(data []byte) error {
		_, _, err := processYaml(data)
		if err != nil {
			log.Error("Failed to validate fronted configuration", "error", err)
			return err
		}
		return nil
	}
}

func (f *fronted) readFrontsFromEmbeddedConfig() {
	yml, err := embedFS.ReadFile("fronted.yaml.gz")
	if err != nil {
		log.Debug("Failed to read smart dialer config", "error", err)
		return
	}
	f.onNewFrontsConfig(yml)
}

func (f *fronted) onNewFrontsConfig(gzippedYaml []byte) {
	pool, providers, err := processYaml(gzippedYaml)
	if err != nil {
		log.Error("Failed to process fronted config", "error", err)
		return
	}
	f.onNewFronts(pool, providers)
}

// onNewFronts sets the domain fronts to use, the trusted root CAs and the fronting providers
// (such as Akamai, Cloudfront, etc)
func (f *fronted) onNewFronts(pool *x509.CertPool, providers map[string]*Provider) {
	// Make copies just to avoid any concurrency issues with access that may be happening on the
	// caller side.
	log.Debug("Updating fronted configuration")
	if len(providers) == 0 {
		log.Error("No providers configured")
		return
	}
	providersCopy := copyProviders(providers, f.countryCode)
	f.addProviders(providersCopy)

	log.Debug("Loading candidates for providers", "numProviders", len(providersCopy))
	fronts := loadFronts(providersCopy, f.cacheDirty)
	log.Debug("Finished loading candidates")

	f.fronts.addFronts(fronts...)
	f.certPool.Store(pool)

	// The goroutine for finding working fronts runs forever, so only start it once.
	f.crawlOnce.Do(func() {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					f.paniced("PANIC finding working fronts", r, debug.Stack())
				}
			}()
			f.findWorkingFronts()
		}()
	})
}

// findWorkingFronts finds working domain fronts by testing them using a worker pool. Speed
// is of the essence here, as without working fronts, users will
// be unable to fetch proxy configurations, particularly in the case of a first time
// user who does not have proxies cached on disk.
func (f *fronted) findWorkingFronts() {
	// Keep looping through all fronts making sure we have working ones.
	for {
		// Continually loop through the fronts until we have 4 working ones.
		// This is important, for example, when the user goes offline and all fronts start failing.
		// We want to just keep trying in that case so that we find working fronts as soon as they
		// come back online.
		if !f.hasEnoughWorkingFronts() {
			// Note that trying all fronts takes awhile, as it only completes when we either
			// have enough working fronts, or we've tried all of them.
			log.Debug("findWorkingFronts::Trying all fronts")
			f.tryAllFronts()
			log.Debug("findWorkingFronts::Tried all fronts")

			// Sleep to avoid spinning infinitely in the case where we don't even know of fronts
			// to try, for example.
			time.Sleep(1 * time.Second)
		} else {
			log.Debug("findWorkingFronts::Enough working fronts...sleeping")
			select {
			case <-f.stopCh:
				log.Debug("findWorkingFronts::Stopping parallel dialing")
				return
			case <-time.After(time.Duration(randRange(6, 12)) * time.Second):
				// Run again after a random time between 0 and 12 seconds
			}
		}
	}
}

// onConnected adds a working front to the channel of working fronts.
func (f *fronted) onConnected(fr Front) {
	f.frontsCh <- fr
}

func (f *fronted) tryAllFronts() {
	// Find working fronts using a worker pool of goroutines.
	pool := pond.NewPool(10)

	// Submit all fronts to the worker pool.
	for i := range f.fronts.frontSize() {
		m := f.fronts.frontAt(i)
		pool.Submit(func() {
			if f.isStopped() {
				return
			}
			if f.hasEnoughWorkingFronts() {
				// We have enough working fronts, so no need to continue.
				// log.Debug("Enough working fronts...ignoring task")
				return
			}
			working := f.vetFront(m)
			if working {
				f.onConnected(m)
			} else {
				m.markFailed()
			}
		})
	}

	// Stop the pool and wait for all submitted tasks to complete
	pool.StopAndWait()
}

func (f *fronted) hasEnoughWorkingFronts() bool {
	return len(f.frontsCh) >= 2
}

func (f *fronted) vetFront(fr Front) bool {
	conn, err := f.dialFront(fr)
	if err != nil {
		log.Debug("unexpected error vetting masquerades", "error", err)
		return false
	}
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	provider := f.providerFor(fr)
	if provider == nil {
		log.Debug("Skipping masquerade with disabled/unknown provider",
			slog.String("providerID", fr.getProviderID()), slog.Any("providers", f.providers))
		return false
	}
	if !fr.markWithResult(fr.verifyWithPost(conn, provider.TestURL)) {
		log.Debug("Unsuccessful vetting with POST request, discarding masquerade")
		return false
	}

	return true
}

func (f *fronted) NewConnectedRoundTripper(ctx context.Context, addr string) (http.RoundTripper, error) {
	for range 6 {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		// Add a case for the stop channel being called
		case <-f.stopCh:
			return nil, errors.New("fronted stopped")
		case fr := <-f.frontsCh:
			// The front may have stopped succeeding since we last checked,
			// so only return it if it's still succeeding.
			if !fr.isSucceeding() {
				continue
			}
			provider := f.providerFor(fr)
			if provider == nil {
				log.Debug("Skipping masquerade with disabled/unknown provider", "providerID", fr.getProviderID())
				fr.markWithResult(false)
				continue
			}

			conn, err := f.dialFront(fr)
			if err != nil {
				log.Debug("failed to dial masquerade", slog.Any("masquerade", fr), slog.String("error", err.Error()))
				fr.markWithResult(false)
				continue
			}
			fr.markWithResult(true)
			// Add the front back to the channel.
			f.frontsCh <- fr

			return newConnectedRoundTripper(fr, conn, provider), err
		}
	}
	return nil, fmt.Errorf("could not connect to any front")
}

func (f *fronted) dialFront(fr Front) (net.Conn, error) {
	log.Debug("Dialing front", slog.Any("front", fr))

	// We do the full TLS connection here because in practice the domains at a given IP
	// address can change frequently on CDNs, so the certificate may not match what
	// we expect.
	conn, retriable, err := f.doDial(fr)
	if err == nil {
		return conn, err
	} else if !retriable {
		log.Debug("Dropping masquerade: not retriable", "error", err)
		fr.markWithResult(false)
	}
	return conn, err
}

func (f *fronted) doDial(fr Front) (net.Conn, bool, error) {
	op := ops.Begin("dial_masquerade")
	defer op.End()
	op.Set("masquerade_domain", fr.getDomain())
	op.Set("masquerade_ip", fr.getIpAddress())
	op.Set("masquerade_provider", fr.getProviderID())

	var conn net.Conn
	var err error
	retriable := false
	// A nil cert pool will just use the system's root CAs.
	pool, typeCorrect := f.certPool.Load().(*x509.CertPool)
	if !typeCorrect || pool == nil {
		pool, err = x509.SystemCertPool()
		if err != nil {
			return nil, retriable, fmt.Errorf("failed to load system cert pool: %w", err)
		}
	}
	conn, err = fr.dial(pool, f.clientHelloID)
	if err != nil {
		if !isNetworkUnreachable(err) {
			op.FailIf(err)
		}
		log.Debug("failed to dial address", slog.String("address", fr.getIpAddress()), slog.String("error", err.Error()))
		// Don't re-add this candidate if it's any certificate error, as that
		// will just keep failing and will waste connections. We can't access the underlying
		// error at this point so just look for "certificate" and "handshake".
		if strings.Contains(err.Error(), "certificate") || strings.Contains(err.Error(), "handshake") {
			log.Debug("Not re-adding candidate", "error", err, "masquerade", fr)
			retriable = false
		} else {
			log.Debug("Unexpected error dialing, keeping masquerade", "error", err, "masquerade", fr)
			retriable = true
		}
	}
	return conn, retriable, err
}

func isNetworkUnreachable(err error) bool {
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		// The following error verifications look for errors that generally happen at Linux/Unix devices
		if errors.Is(opErr.Err, syscall.ENETUNREACH) || errors.Is(opErr.Err, syscall.EHOSTUNREACH) {
			return true
		}

		// The string verification errors use a broader approach with errors from windows and also linux/unix devices
		errMsg := opErr.Err.Error()
		if strings.Contains(errMsg, "network is unreachable") ||
			strings.Contains(errMsg, "no route to host") ||
			strings.Contains(errMsg, "unreachable network") ||
			strings.Contains(errMsg, "unreachable host") {
			return true
		}
	}
	return false
}

func verifyPeerCertificate(rawCerts [][]byte, roots *x509.CertPool, domain string) error {
	if len(rawCerts) == 0 {
		return fmt.Errorf("no certificates presented")
	}
	cert, err := x509.ParseCertificate(rawCerts[0])
	if err != nil {
		return fmt.Errorf("unable to parse certificate: %w", err)
	}

	opts := []x509.VerifyOptions{generateVerifyOptions(roots, domain)}
	for i := range rawCerts {
		if i == 0 {
			continue
		}
		crt, err := x509.ParseCertificate(rawCerts[i])
		if err != nil {
			return fmt.Errorf("unable to parse intermediate certificate: %w", err)
		}

		for _, opt := range opts {
			opt.Intermediates.AddCert(crt)
		}
	}

	var verificationErrors error
	for _, opt := range opts {
		_, err := cert.Verify(opt)
		if err != nil {
			verificationErrors = errors.Join(verificationErrors, err)
		}
	}

	if verificationErrors != nil {
		return fmt.Errorf("certificate verification failed: %w", verificationErrors)
	}

	return nil
}

func generateVerifyOptions(roots *x509.CertPool, domain string) x509.VerifyOptions {
	return x509.VerifyOptions{
		Roots:         roots,
		CurrentTime:   time.Now(),
		DNSName:       domain,
		Intermediates: x509.NewCertPool(),
	}
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

func (f *fronted) Close() {
	f.stopped.Store(true)
	f.closeCacheOnce.Do(func() {
		close(f.cacheClosed)
	})
	f.stopCh <- nil
}

func (f *fronted) isStopped() bool {
	return f.stopped.Load()
}

func copyProviders(providers map[string]*Provider, countryCode string) map[string]*Provider {
	providersCopy := make(map[string]*Provider, len(providers))
	for key, p := range providers {
		providersCopy[key] = NewProvider(p.HostAliases, p.TestURL, p.Masquerades, p.PassthroughPatterns, p.FrontingSNIs, p.VerifyHostname, countryCode)
	}
	return providersCopy
}

func loadFronts(providers map[string]*Provider, cacheDirty chan interface{}) []Front {
	// Preallocate the slice to avoid reallocation
	size := 0
	for _, p := range providers {
		size += len(p.Masquerades)
	}

	fronts := make([]Front, size)

	// Note that map iteration order is random, so the order of the providers is automatically randomized.
	index := 0
	for key, p := range providers {
		arr := p.Masquerades
		size := len(arr)

		// Shuffle the masquerades to avoid biasing the order in which they are tried
		// make a shuffled copy of arr
		// ('inside-out' Fisher-Yates)
		sh := make([]*Masquerade, size)
		for i := 0; i < size; i++ {
			j := rand.IntN(i + 1) // 0 <= j <= i
			sh[i] = sh[j]
			sh[j] = arr[i]
		}

		for _, c := range sh {
			fronts[index] = newFront(c, key, cacheDirty)
			index++
		}
	}
	return fronts
}

func (f *fronted) addProviders(providers map[string]*Provider) {
	// Add new providers to the existing providers map, overwriting any existing ones.
	f.providersMu.Lock()
	defer f.providersMu.Unlock()
	for key, p := range providers {
		f.providers[key] = p
	}
}

func (f *fronted) providerFor(m Front) *Provider {
	pid := m.getProviderID()
	if pid == "" {
		pid = f.defaultProviderID
	}
	return f.providers[pid]
}

// Vet vets the specified Masquerade, verifying certificate using the given CertPool.
// This is used in genconfig.
func Vet(m *Masquerade, pool *x509.CertPool, testURL string) bool {
	d := &fronted{
		certPool:            atomic.Value{},
		maxAllowedCachedAge: defaultMaxAllowedCachedAge,
		maxCacheSize:        defaultMaxCacheSize,
	}
	d.certPool.Store(pool)
	masq := &front{Masquerade: *m}
	conn, _, err := d.doDial(masq)
	if err != nil {
		return false
	}
	defer conn.Close()
	return masq.verifyWithPost(conn, testURL)
}

func processYaml(gzippedYaml []byte) (*x509.CertPool, map[string]*Provider, error) {
	r, gzipErr := gzip.NewReader(bytes.NewReader(gzippedYaml))
	if gzipErr != nil {
		// Wrap the error
		return nil, nil, fmt.Errorf("failed to create gzip reader: %w", gzipErr)
	}
	yml, err := io.ReadAll(r)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read gzipped file: %w", err)
	}
	path, err := yaml.PathString("$.providers")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create providers path: %w", err)
	}
	providers := make(map[string]*Provider)
	pathErr := path.Read(bytes.NewReader(yml), &providers)
	if pathErr != nil {
		return nil, nil, fmt.Errorf("failed to read providers: %w", pathErr)
	}

	trustedCAsPath, err := yaml.PathString("$.trustedcas")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create trusted CA path: %w", err)
	}
	var trustedCAs []*CA
	trustedCAsErr := trustedCAsPath.Read(bytes.NewReader(yml), &trustedCAs)
	if trustedCAsErr != nil {
		return nil, nil, fmt.Errorf("failed to read trusted CAs: %w", trustedCAsErr)
	}
	pool := x509.NewCertPool()
	for _, ca := range trustedCAs {
		pool.AppendCertsFromPEM([]byte(ca.Cert))
	}
	return pool, providers, nil
}
