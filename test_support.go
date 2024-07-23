package fronted

import (
	"crypto/x509"
	"testing"

	"github.com/getlantern/keyman"
)

var (
	testProviderID  = "cloudfront"
	pingTestURL     = "https://d157vud77ygy87.cloudfront.net/ping"
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

func ConfigureHostAlaisesForTest(t *testing.T, hosts map[string]string) {
	certs := trustedCACerts(t)
	p := testProvidersWithHosts(hosts)
	Configure(certs, p, testProviderID, "")
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
		testProviderID: NewProvider(testHosts, pingTestURL, testMasquerades, nil, nil, nil),
	}
}

func testProvidersWithHosts(hosts map[string]string) map[string]*Provider {
	return map[string]*Provider{
		testProviderID: NewProvider(hosts, pingTestURL, testMasquerades, nil, nil, nil),
	}
}
func testAkamaiProvidersWithHosts(hosts map[string]string, sniConfig *SNIConfig) map[string]*Provider {
	return map[string]*Provider{
		testProviderID: NewProvider(hosts, pingTestURL, DefaultAkamaiMasquerades, nil, nil, sniConfig),
	}
}
