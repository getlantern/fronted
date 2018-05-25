package fronted

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDirectDomainFronting(t *testing.T) {
	dir, err := ioutil.TempDir("", "direct_test")
	if !assert.NoError(t, err, "Unable to create temp dir") {
		return
	}
	defer os.RemoveAll(dir)
	cacheFile := filepath.Join(dir, "cachefile.2")
	doTestDomainFronting(t, cacheFile)
	time.Sleep(defaultCacheSaveInterval * 2)
	// Then try again, this time reusing the existing cacheFile
	doTestDomainFronting(t, cacheFile)
}

func doTestDomainFronting(t *testing.T, cacheFile string) {
	ConfigureCachingForTest(t, cacheFile)
	direct, ok := NewDirect(30 * time.Second)
	if !assert.True(t, ok) {
		return
	}
	client := &http.Client{
		Transport: direct,
	}
	assert.True(t, doCheck(client, http.MethodPost, http.StatusAccepted, testURL))

	direct, ok = NewDirect(30 * time.Second)
	if !assert.True(t, ok) {
		return
	}
	client = &http.Client{
		Transport: direct,
	}
	assert.True(t, doCheck(client, http.MethodGet, http.StatusOK, "http://d2wi0vwulmtn99.cloudfront.net/proxies.yaml.gz"))
}

func TestVet(t *testing.T) {
	pool := trustedCACerts(t)
	for _, m := range DefaultCloudfrontMasquerades {
		if Vet(m, pool) {
			return
		}
	}
	t.Fatal("None of the default masquerades vetted successfully")
}

func TestLoadCandidates(t *testing.T) {
	expected := make(map[Masquerade]bool, len(DefaultCloudfrontMasquerades))
	for _, m := range DefaultCloudfrontMasquerades {
		expected[*m] = true
	}

	d := &direct{
		candidates: make(chan masquerade, len(DefaultCloudfrontMasquerades)),
	}
	d.loadCandidates(map[string][]*Masquerade{"cloudfront": DefaultCloudfrontMasquerades})
	close(d.candidates)

	actual := make(map[Masquerade]bool)
	for m := range d.candidates {
		actual[Masquerade{m.Domain, m.IpAddress}] = true
	}

	assert.Equal(t, expected, actual, "Masquerades did not load as expected")
}
