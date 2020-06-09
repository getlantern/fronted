package fronted

import (
	"crypto/x509"
	"testing"

	"github.com/getlantern/keyman"
)

var (
	testProviderID  = "cloudfront"
	pingTestURL     = "http://d157vud77ygy87.cloudfront.net/ping"
	testHosts       = map[string]string(nil)
	testMasquerades = DefaultCloudfrontMasquerades
)

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
		testProviderID: NewProvider(testHosts, pingTestURL, testMasquerades, nil, nil),
	}
}

func testProvidersWithHosts(hosts map[string]string) map[string]*Provider {
	return map[string]*Provider{
		testProviderID: NewProvider(hosts, pingTestURL, testMasquerades, nil, nil),
	}
}
