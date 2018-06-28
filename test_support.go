package fronted

import (
	"crypto/x509"
	"testing"

	"github.com/getlantern/keyman"
)

var (
	testProviderID  = "cloudfront"
	pingTestURL     = "http://d157vud77ygy87.cloudfront.net/ping"
	getTestURL      = "http://d2wi0vwulmtn99.cloudfront.net/proxies.yaml.gz"
	testHosts       = map[string]string(nil)
	testMasquerades = DefaultCloudfrontMasquerades
)

// ConfigureForTest configures fronted for testing using default masquerades and
// certificate authorities.
func ConfigureForTest(t *testing.T) {
	ConfigureCachingForTest(t, "")
}

func ConfigureCachingForTest(t *testing.T, cacheFile string) {
	certs := trustedCACerts(t)
	p := testProviders()
	Configure(certs, p, testProviderID, cacheFile)
}

func trustedCACerts(t *testing.T) *x509.CertPool {
	certs := make([]string, 0, len(DefaultTrustedCAs))
	for _, ca := range DefaultTrustedCAs {
		certs = append(certs, ca.Cert)
	}
	pool, err := keyman.PoolContainingCerts(certs...)
	if err != nil {
		log.Errorf("Could not create pool %v", err)
		t.Fatalf("Unable to set up cert pool")
	}
	return pool
}

func testProviders() map[string]*Provider {
	return map[string]*Provider{
		testProviderID: NewProvider(testHosts, pingTestURL, testMasquerades),
	}
}
