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

const (
	humansText = "Google is built by a large team of engineers, designers, researchers, robots, and others in many different sites across the globe. It is updated continuously, and built with more tools and technologies than we can shake a stick at. If you'd like to help us out, see google.com/careers.\n"
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

func TestProxying(t *testing.T) {
	dir, err := ioutil.TempDir("", "direct_test")
	if !assert.NoError(t, err, "Unable to create temp dir") {
		return
	}
	defer os.RemoveAll(dir)
	cacheFile := filepath.Join(dir, "cachefile.3")
	ConfigureCachingForTest(t, cacheFile)

	proxying, ok := NewProxyingAt("d100fjyl3713ch.cloudfront.net", 30*time.Second)
	if !assert.True(t, ok) {
		return
	}
	client := &http.Client{
		Transport: proxying,
	}
	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/humans.txt", nil)

	resp, err := client.Do(req)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		return
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, humansText, string(respBody))
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
