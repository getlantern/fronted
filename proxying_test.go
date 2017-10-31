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

func TestProxying(t *testing.T) {
	dir, err := ioutil.TempDir("", "direct_test")
	if !assert.NoError(t, err, "Unable to create temp dir") {
		return
	}
	defer os.RemoveAll(dir)
	cacheFile := filepath.Join(dir, "cachefile.3")
	ConfigureCachingForTest(t, cacheFile)

	conn, err := DialTimeout(30 * time.Second)
	if !assert.NoError(t, err) {
		return
	}
	defer conn.Close()

	req, _ := http.NewRequest(http.MethodGet, "https://www.google.com/humans.txt", nil)
	req.Header.Set("X-Lantern-Auth-Token", "pj6mWPafKzP26KZvUf7FIs24eB2ubjUKFvXktodqgUzZULhGeRUT0mwhyHb9jY2b")
	PrepareForProxyingVia("d100fjyl3713ch.cloudfront.net", req)
	resp, err := httpTransport(conn, clientSessionCache).RoundTrip(req)
	if !assert.NoError(t, err) {
		return
	}
	AfterProxying(conn, req, resp, err)
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
