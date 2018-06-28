package fronted

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	. "github.com/getlantern/waitforserver"
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
	assert.True(t, doCheck(client, http.MethodPost, http.StatusAccepted, pingTestURL))

	direct, ok = NewDirect(30 * time.Second)
	if !assert.True(t, ok) {
		return
	}
	client = &http.Client{
		Transport: direct,
	}
	assert.True(t, doCheck(client, http.MethodGet, http.StatusOK, getTestURL))
}

func TestVet(t *testing.T) {
	pool := trustedCACerts(t)
	for _, m := range testMasquerades {
		if Vet(m, pool, pingTestURL) {
			return
		}
	}
	t.Fatal("None of the default masquerades vetted successfully")
}

func TestLoadCandidates(t *testing.T) {
	providers := testProviders()

	expected := make(map[Masquerade]bool)
	for _, p := range providers {
		for _, m := range p.Masquerades {
			expected[*m] = true
		}
	}

	d := &direct{
		candidates: make(chan masquerade, len(expected)),
	}

	d.loadCandidates(providers)
	close(d.candidates)

	actual := make(map[Masquerade]bool)
	count := 0
	for m := range d.candidates {
		actual[Masquerade{m.Domain, m.IpAddress}] = true
		count++
	}

	assert.Equal(t, len(DefaultCloudfrontMasquerades), count, "Unexpected number of candidates")
	assert.Equal(t, expected, actual, "Masquerades did not load as expected")
}

func TestHostAliasesBasic(t *testing.T) {

	tests := []struct {
		url            string
		expectedResult CDNResult
		expectedStatus int
	}{
		{
			"http://abc.forbidden.com/foo/bar",
			CDNResult{"abc.cloudsack.biz", "/foo/bar", "", "cloudsack"},
			http.StatusAccepted,
		},
		{
			"https://abc.forbidden.com/bar?x=y&z=w",
			CDNResult{"abc.cloudsack.biz", "/bar", "x=y&z=w", "cloudsack"},
			http.StatusAccepted,
		},
		{
			"http://def.forbidden.com/foo",
			CDNResult{"def.cloudsack.biz", "/foo", "", "cloudsack"},
			http.StatusAccepted,
		},
		{
			"https://def.forbidden.com/bar?x=y&z=w",
			CDNResult{"def.cloudsack.biz", "/bar", "x=y&z=w", "cloudsack"},
			http.StatusAccepted,
		},
		// not translated, but permitted
		{
			"http://fff.cloudsack.biz/foo",
			CDNResult{"fff.cloudsack.biz", "/foo", "", "cloudsack"},
			http.StatusAccepted,
		},
		{
			"http://fff.cloudsack.biz/bar?x=y&z=w",
			CDNResult{"fff.cloudsack.biz", "/bar", "x=y&z=w", "cloudsack"},
			http.StatusAccepted,
		},
	}

	cloudSack, cloudSackAddr, err := newCDN("cloudsack", "cloudsack.biz")
	if !assert.NoError(t, err, "failed to start cloudsack cdn") {
		return
	}
	defer cloudSack.Close()

	masq := []*Masquerade{&Masquerade{Domain: "example.com", IpAddress: cloudSackAddr}}
	alias := map[string]string{
		"abc.forbidden.com": "abc.cloudsack.biz",
		"def.forbidden.com": "def.cloudsack.biz",
	}
	p := NewProvider(alias, "https://ttt.cloudsack.biz/ping", masq)

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())
	Configure(certs, map[string]*Provider{"cloudsack": p}, "cloudsack", "")

	rt, ok := NewDirect(10 * time.Second)
	if !assert.True(t, ok, "failed to obtain direct roundtripper") {
		return
	}
	client := &http.Client{Transport: rt}
	for _, test := range tests {
		resp, err := client.Get(test.url)
		if !assert.NoError(t, err, "Request %s failed", test.url) {
			continue
		}
		assert.Equal(t, test.expectedStatus, resp.StatusCode)
		if !assert.NotNil(t, resp.Body) {
			continue
		}

		var result CDNResult
		data, err := ioutil.ReadAll(resp.Body)
		if !assert.NoError(t, err) {
			continue
		}

		err = json.Unmarshal(data, &result)
		if !assert.NoError(t, err) {
			continue
		}
		assert.Equal(t, test.expectedResult, result)
	}

	// this is not allowed, so masqurades are discarded and
	// an error results...
	_, err = client.Get("https://example.biz/baz")
	assert.NotNil(t, err)
}

func TestHostAliasesMulti(t *testing.T) {

	tests := []struct {
		url            string
		expectedStatus int
		expectedPath   string
		expectedQuery  string
		expectedHosts  []string
	}{
		{
			"http://abc.forbidden.com/foo/bar",
			http.StatusAccepted,
			"/foo/bar",
			"",
			[]string{
				"abc.cloudsack.biz",
				"abc.sadcloud.io",
			},
		},
		{
			"http://def.forbidden.com/bar?x=y&z=w",
			http.StatusAccepted,
			"/bar",
			"x=y&z=w",
			[]string{
				"def.cloudsack.biz",
				"def.sadcloud.io",
			},
		},
	}

	sadCloud, sadCloudAddr, err := newCDN("sadcloud", "sadcloud.io")
	if !assert.NoError(t, err, "failed to start sadcloud cdn") {
		return
	}
	defer sadCloud.Close()

	cloudSack, cloudSackAddr, err := newCDN("cloudsack", "cloudsack.biz")
	if !assert.NoError(t, err, "failed to start cloudsack cdn") {
		return
	}
	defer cloudSack.Close()

	masq1 := []*Masquerade{&Masquerade{Domain: "example.com", IpAddress: cloudSackAddr}}
	alias1 := map[string]string{
		"abc.forbidden.com": "abc.cloudsack.biz",
		"def.forbidden.com": "def.cloudsack.biz",
	}
	p1 := NewProvider(alias1, "https://ttt.cloudsack.biz/ping", masq1)

	masq2 := []*Masquerade{&Masquerade{Domain: "example.com", IpAddress: sadCloudAddr}}
	alias2 := map[string]string{
		"abc.forbidden.com": "abc.sadcloud.io",
		"def.forbidden.com": "def.sadcloud.io",
	}
	p2 := NewProvider(alias2, "https://ttt.sadcloud.io/ping", masq2)

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())
	certs.AddCert(sadCloud.Certificate())

	providers := map[string]*Provider{
		"cloudsack": p1,
		"sadcloud":  p2,
	}

	Configure(certs, providers, "cloudsack", "")
	rt, ok := NewDirect(10 * time.Second)
	if !assert.True(t, ok, "failed to obtain direct roundtripper") {
		return
	}
	client := &http.Client{Transport: rt}

	providerCounts := make(map[string]int)

	for i := 0; i < 10; i++ {
		for _, test := range tests {
			resp, err := client.Get(test.url)
			if !assert.NoError(t, err, "Request %s failed", test.url) {
				continue
			}
			assert.Equal(t, test.expectedStatus, resp.StatusCode)
			if !assert.NotNil(t, resp.Body) {
				continue
			}

			var result CDNResult
			data, err := ioutil.ReadAll(resp.Body)
			if !assert.NoError(t, err) {
				continue
			}

			err = json.Unmarshal(data, &result)
			if !assert.NoError(t, err) {
				continue
			}
			assert.Contains(t, test.expectedHosts, result.Host)
			assert.Equal(t, test.expectedQuery, result.Query)
			assert.Equal(t, test.expectedPath, result.Path)

			providerCounts[result.Provider] += 1
		}
	}

	assert.True(t, providerCounts["cloudsack"] > 1)
	assert.True(t, providerCounts["sadcloud"] > 1)
}

type CDNResult struct {
	Host, Path, Query, Provider string
}

func newCDN(providerID, domain string) (*httptest.Server, string, error) {
	allowedSuffix := fmt.Sprintf(".%s", domain)
	srv := httptest.NewTLSServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				log.Errorf("Failed to dump request: %s", err)
			} else {
				log.Debugf("(%s) CDN Request: %s", domain, dump)
			}

			vhost := req.Host
			if strings.HasSuffix(vhost, allowedSuffix) {
				body, _ := json.Marshal(&CDNResult{
					Host:     vhost,
					Path:     req.URL.Path,
					Query:    req.URL.RawQuery,
					Provider: providerID,
				})
				rw.WriteHeader(http.StatusAccepted)
				rw.Write(body)
			} else {
				log.Debugf("(%s) Rejecting request with host = %q", domain, vhost)
				rw.WriteHeader(http.StatusForbidden)
			}
		}))
	addr := srv.Listener.Addr().String()
	log.Debugf("Waiting for origin server at %s...", addr)
	if err := WaitForServer("tcp", addr, 10*time.Second); err != nil {
		return nil, "", err
	}
	log.Debugf("Started %s CDN", domain)
	return srv, addr, nil
}
