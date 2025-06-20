package fronted

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	. "github.com/getlantern/waitforserver"
	tls "github.com/refraction-networking/utls"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfigUpdating(t *testing.T) {
	f := NewFronted(
		WithConfigURL("https://media.githubusercontent.com/media/getlantern/fronted/refs/heads/main/fronted.yaml.gz"),
		WithCountryCode("cn"),
	)
	time.Sleep(1 * time.Second)

	// Try to hit raw.githubusercontent.com
	rt, err := f.NewConnectedRoundTripper(context.Background(), "")
	require.NoError(t, err)
	client := &http.Client{
		Transport: rt,
	}
	resp, err := client.Get("https://raw.githubusercontent.com/getlantern/fronted/main/fronted.yaml.gz")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	// Read the full response body
	bod, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Greater(t, len(bod), 0)
	resp.Body.Close()
}

func TestYamlParsing(t *testing.T) {
	// Disable this if we're running in CI because the file is using git lfs and will just be a pointer.
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		t.Skip("Skipping test in GitHub Actions because the file is using git lfs and will be a pointer")
	}
	yamlFile, err := os.ReadFile("fronted.yaml.gz")
	require.NoError(t, err)
	pool, providers, err := processYaml(yamlFile)
	require.NoError(t, err)
	require.NotNil(t, pool)
	require.NotNil(t, providers)

	// Make sure there are some providers
	assert.Greater(t, len(providers), 0)
}

func TestDirectDomainFrontingWithoutSNIConfig(t *testing.T) {
	dir := t.TempDir()
	cacheFile := filepath.Join(dir, "cachefile.2")

	t.Log("Testing direct domain fronting without SNI config")
	doTestDomainFronting(t, cacheFile, 10)
	time.Sleep(defaultCacheSaveInterval * 2)
	// Then try again, this time reusing the existing cacheFile but a corrupted version
	n, err := corruptMasquerades(cacheFile)
	require.NoError(t, err, "Unable to corrupt masquerades")
	t.Logf("Corrupted %d masquerades", n)
	t.Log("Testing direct domain fronting without SNI config again")
	doTestDomainFronting(t, cacheFile, 10)
}

func TestDirectDomainFrontingWithSNIConfig(t *testing.T) {
	dir := t.TempDir()
	cacheFile := filepath.Join(dir, "cachefile.3")

	getURL := "https://config.example.com/global.yaml.gz"
	getHost := "config.example.com"
	getFrontedHost := "globalconfig.dsa.akamai.getiantem.org"

	hosts := map[string]string{
		getHost: getFrontedHost,
	}
	certs := trustedCACerts(t)
	p := testAkamaiProvidersWithHosts(hosts, &SNIConfig{
		UseArbitrarySNIs: true,
		ArbitrarySNIs:    []string{"mercadopago.com", "amazon.com.br", "facebook.com", "google.com", "twitter.com", "youtube.com", "instagram.com", "linkedin.com", "whatsapp.com", "netflix.com", "microsoft.com", "yahoo.com", "bing.com", "wikipedia.org", "github.com"},
	})
	defaultFrontedProviderID = "akamai"
	transport := NewFronted(WithCacheFile(cacheFile), WithCountryCode("test"))
	transport.onNewFronts(certs, p)

	client := &http.Client{
		Transport: newTransportFromDialer(transport),
	}
	require.True(t, doCheck(client, http.MethodGet, http.StatusOK, getURL))
}

func newTransportFromDialer(f Fronted) http.RoundTripper {
	rt, _ := f.NewConnectedRoundTripper(context.Background(), "")
	return rt
}

func doTestDomainFronting(t *testing.T, cacheFile string, expectedMasqueradesAtEnd int) int {
	getURL := "https://config.example.com/global.yaml.gz"
	getHost := "config.example.com"
	getFrontedHost := "d24ykmup0867cj.cloudfront.net"

	pingHost := "ping.example.com"
	pu, err := url.Parse(pingTestURL)
	require.NoError(t, err)
	pingFrontedHost := pu.Hostname()
	pu.Host = pingHost
	pingURL := pu.String()

	hosts := map[string]string{
		pingHost: pingFrontedHost,
		getHost:  getFrontedHost,
	}
	certs := trustedCACerts(t)
	p := testProvidersWithHosts(hosts)
	defaultFrontedProviderID = testProviderID
	transport := NewFronted(WithCacheFile(cacheFile))
	transport.onNewFronts(certs, p)

	rt := newTransportFromDialer(transport)
	client := &http.Client{
		Transport: rt,
		Timeout:   5 * time.Second,
	}
	require.True(t, doCheck(client, http.MethodPost, http.StatusAccepted, pingURL))

	defaultFrontedProviderID = testProviderID
	transport = NewFronted(WithCacheFile(cacheFile))
	transport.onNewFronts(certs, p)
	client = &http.Client{
		Transport: newTransportFromDialer(transport),
	}
	require.True(t, doCheck(client, http.MethodGet, http.StatusOK, getURL))

	d := transport.(*fronted)

	// Check the number of masquerades at the end, waiting until we get the right number
	masqueradesAtEnd := 0
	for i := 0; i < 1000; i++ {
		masqueradesAtEnd = len(d.fronts.fronts)
		if masqueradesAtEnd == expectedMasqueradesAtEnd {
			break
		}
		time.Sleep(30 * time.Millisecond)
	}
	require.GreaterOrEqual(t, masqueradesAtEnd, expectedMasqueradesAtEnd)
	return masqueradesAtEnd
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

func TestHostAliasesBasic(t *testing.T) {
	headersIn := map[string][]string{
		"X-Foo-Bar": {"Quux", "Baz"},
		"X-Bar-Foo": {"XYZ"},
		"X-Quux":    {""},
	}
	headersOut := map[string][]string{
		"X-Foo-Bar":       {"Quux", "Baz"},
		"X-Bar-Foo":       {"XYZ"},
		"X-Quux":          {""},
		"Connection":      {"close"},
		"User-Agent":      {"Go-http-client/1.1"},
		"Accept-Encoding": {"gzip"},
	}

	tests := []struct {
		url            string
		headers        map[string][]string
		expectedResult CDNResult
		expectedStatus int
	}{
		{
			"http://abc.forbidden.com/foo/bar",
			headersIn,
			CDNResult{"abc.cloudsack.biz", "/foo/bar", "", "cloudsack", headersOut},
			http.StatusAccepted,
		},
		{
			"https://abc.forbidden.com/bar?x=y&z=w",
			headersIn,
			CDNResult{"abc.cloudsack.biz", "/bar", "x=y&z=w", "cloudsack", headersOut},
			http.StatusAccepted,
		},
		{
			"http://def.forbidden.com:12345/foo",
			headersIn,
			CDNResult{"def.cloudsack.biz", "/foo", "", "cloudsack", headersOut},
			http.StatusAccepted,
		},
		{
			"https://def.forbidden.com/bar?x=y&z=w",
			headersIn,
			CDNResult{"def.cloudsack.biz", "/bar", "x=y&z=w", "cloudsack", headersOut},
			http.StatusAccepted,
		},
	}

	errtests := []struct {
		url           string
		expectedError string
	}{
		{
			"http://fff.cloudsack.biz/foo",
			`Get "http://fff.cloudsack.biz/foo": could not complete request even with retries`,
		},
		{
			"http://fff.cloudsack.biz:1234/bar?x=y&z=w",
			`Get "http://fff.cloudsack.biz:1234/bar?x=y&z=w": could not complete request even with retries`,
		},
		{
			"https://www.google.com",
			`Get "https://www.google.com": could not complete request even with retries`,
		},
	}

	cloudSack, cloudSackAddr, err := newCDN(t, "cloudsack", "cloudsack.biz")
	if !assert.NoError(t, err, "failed to start cloudsack cdn") {
		return
	}
	defer cloudSack.Close()

	masq := []*Masquerade{{Domain: "example.com", IpAddress: cloudSackAddr}}
	alias := map[string]string{
		"abc.forbidden.com": "abc.cloudsack.biz",
		"def.forbidden.com": "def.cloudsack.biz",
	}
	p := NewProvider(alias, "https://ttt.cloudsack.biz/ping", masq, nil, nil, nil, "")

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())

	defaultFrontedProviderID = "cloudsack"
	rt := NewFronted()

	rt.onNewFronts(certs, map[string]*Provider{"cloudsack": p})
	for _, test := range tests {
		client := &http.Client{Transport: newTransportFromDialer(rt)}
		req, err := http.NewRequest(http.MethodGet, test.url, nil)
		if !assert.NoError(t, err) {
			return
		}

		for k, v := range test.headers {
			req.Header[k] = v
		}
		resp, err := client.Do(req)
		if !assert.NoError(t, err, "Request %s failed", test.url) {
			continue
		}
		assert.Equal(t, test.expectedStatus, resp.StatusCode)
		if !assert.NotNil(t, resp.Body) {
			continue
		}

		var result CDNResult
		data, err := io.ReadAll(resp.Body)
		if !assert.NoError(t, err) {
			continue
		}

		err = json.Unmarshal(data, &result)
		if !assert.NoError(t, err) {
			continue
		}
		assert.Equal(t, test.expectedResult, result)
	}

	for _, test := range errtests {
		client := &http.Client{Transport: newTransportFromDialer(rt)}
		resp, err := client.Get(test.url)
		assert.EqualError(t, err, test.expectedError)
		assert.Nil(t, resp)
	}
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

	sadCloud, sadCloudAddr, err := newCDN(t, "sadcloud", "sadcloud.io")
	if !assert.NoError(t, err, "failed to start sadcloud cdn") {
		return
	}
	defer sadCloud.Close()

	cloudSack, cloudSackAddr, err := newCDN(t, "cloudsack", "cloudsack.biz")
	if !assert.NoError(t, err, "failed to start cloudsack cdn") {
		return
	}
	defer cloudSack.Close()

	masq1 := []*Masquerade{{Domain: "example.com", IpAddress: cloudSackAddr}}
	alias1 := map[string]string{
		"abc.forbidden.com": "abc.cloudsack.biz",
		"def.forbidden.com": "def.cloudsack.biz",
	}
	p1 := NewProvider(alias1, "https://ttt.cloudsack.biz/ping", masq1, nil, nil, nil, "")

	masq2 := []*Masquerade{{Domain: "example.com", IpAddress: sadCloudAddr}}
	alias2 := map[string]string{
		"abc.forbidden.com": "abc.sadcloud.io",
		"def.forbidden.com": "def.sadcloud.io",
	}
	p2 := NewProvider(alias2, "https://ttt.sadcloud.io/ping", masq2, nil, nil, nil, "")

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())
	certs.AddCert(sadCloud.Certificate())

	providers := map[string]*Provider{
		"cloudsack": p1,
		"sadcloud":  p2,
	}

	defaultFrontedProviderID = "cloudsack"
	rt := NewFronted()
	rt.onNewFronts(certs, providers)

	providerCounts := make(map[string]int)

	for i := 0; i < 10; i++ {
		for _, test := range tests {
			client := &http.Client{Transport: newTransportFromDialer(rt)}
			resp, err := client.Get(test.url)
			if !assert.NoError(t, err, "Request %s failed", test.url) {
				continue
			}
			assert.Equal(t, test.expectedStatus, resp.StatusCode)
			if !assert.NotNil(t, resp.Body) {
				continue
			}

			var result CDNResult
			data, err := io.ReadAll(resp.Body)
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

	assert.True(t, providerCounts["cloudsack"]+providerCounts["sadcloud"] > 2)
}

func TestPassthrough(t *testing.T) {
	headersIn := map[string][]string{
		"X-Foo-Bar": {"Quux", "Baz"},
		"X-Bar-Foo": {"XYZ"},
		"X-Quux":    {""},
	}
	headersOut := map[string][]string{
		"X-Foo-Bar":       {"Quux", "Baz"},
		"X-Bar-Foo":       {"XYZ"},
		"X-Quux":          {""},
		"Connection":      {"close"},
		"User-Agent":      {"Go-http-client/1.1"},
		"Accept-Encoding": {"gzip"},
	}

	tests := []struct {
		url            string
		headers        map[string][]string
		expectedResult CDNResult
		expectedStatus int
	}{
		{
			"http://fff.ok.cloudsack.biz/foo",
			headersIn,
			CDNResult{"fff.ok.cloudsack.biz", "/foo", "", "cloudsack", headersOut},
			http.StatusAccepted,
		},
		{
			"http://abc.cloudsack.biz/bar",
			headersIn,
			CDNResult{"abc.cloudsack.biz", "/bar", "", "cloudsack", headersOut},
			http.StatusAccepted,
		},
		{
			"http://XYZ.ZyZ.OK.CloudSack.BiZ/bar",
			headersIn,
			CDNResult{"xyz.zyz.ok.cloudsack.biz", "/bar", "", "cloudsack", headersOut},
			http.StatusAccepted,
		},
	}

	errtests := []struct {
		url           string
		expectedError string
	}{
		{
			"http://www.notok.cloudsack.biz",
			`Get "http://www.notok.cloudsack.biz": could not complete request even with retries`,
		},
		{
			"http://ok.cloudsack.biz",
			`Get "http://ok.cloudsack.biz": could not complete request even with retries`,
		},
		{
			"http://www.abc.cloudsack.biz",
			`Get "http://www.abc.cloudsack.biz": could not complete request even with retries`,
		},
		{
			"http://noabc.cloudsack.biz",
			`Get "http://noabc.cloudsack.biz": could not complete request even with retries`,
		},
		{
			"http://cloudsack.biz",
			`Get "http://cloudsack.biz": could not complete request even with retries`,
		},
		{
			"https://www.google.com",
			`Get "https://www.google.com": could not complete request even with retries`,
		},
	}

	cloudSack, cloudSackAddr, err := newCDN(t, "cloudsack", "cloudsack.biz")
	if !assert.NoError(t, err, "failed to start cloudsack cdn") {
		return
	}
	defer cloudSack.Close()

	masq := []*Masquerade{{Domain: "example.com", IpAddress: cloudSackAddr}}
	alias := map[string]string{}
	passthrough := []string{"*.ok.cloudsack.biz", "abc.cloudsack.biz"}
	p := NewProvider(alias, "https://ttt.cloudsack.biz/ping", masq, passthrough, nil, nil, "")

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())

	defaultFrontedProviderID = "cloudsack"
	rt := NewFronted()
	rt.onNewFronts(certs, map[string]*Provider{"cloudsack": p})

	for _, test := range tests {
		client := &http.Client{Transport: newTransportFromDialer(rt)}
		req, err := http.NewRequest(http.MethodGet, test.url, nil)
		if !assert.NoError(t, err) {
			return
		}

		for k, v := range test.headers {
			req.Header[k] = v
		}
		resp, err := client.Do(req)
		if !assert.NoError(t, err, "Request %s failed", test.url) {
			continue
		}
		assert.Equal(t, test.expectedStatus, resp.StatusCode)
		if !assert.NotNil(t, resp.Body) {
			continue
		}

		var result CDNResult
		data, err := io.ReadAll(resp.Body)
		if !assert.NoError(t, err) {
			continue
		}

		err = json.Unmarshal(data, &result)
		if !assert.NoError(t, err) {
			continue
		}
		assert.Equal(t, test.expectedResult, result)
	}

	for _, test := range errtests {
		client := &http.Client{Transport: newTransportFromDialer(rt)}
		resp, err := client.Get(test.url)
		assert.EqualError(t, err, test.expectedError)
		assert.Nil(t, resp)

	}
}

const (
	// set this header to an integer to force response status code
	CDNForceFail = "X-CDN-Force-Fail"
)

type CDNResult struct {
	Host, Path, Query, Provider string
	Headers                     map[string][]string
}

func newCDN(t *testing.T, providerID, domain string) (*httptest.Server, string, error) {
	var logf = t.Logf
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		// don't log request dumps in CI
		logf = func(format string, args ...any) {}
	}

	allowedSuffix := "." + domain
	srv := httptest.NewTLSServer(
		http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				logf("Failed to dump request: %s", err)
			} else {
				logf("(%s) CDN Request: %s", domain, dump)
			}

			forceFail := req.Header.Get(CDNForceFail)

			vhost := req.Host
			if vhost == domain || strings.HasSuffix(vhost, allowedSuffix) && forceFail == "" {
				logf("accepting request host=%s ff=%s", vhost, forceFail)
				body, _ := json.Marshal(&CDNResult{
					Host:     vhost,
					Path:     req.URL.Path,
					Query:    req.URL.RawQuery,
					Provider: providerID,
					Headers:  req.Header,
				})
				rw.WriteHeader(http.StatusAccepted)
				rw.Write(body)
			} else {
				logf("(%s) Rejecting request with host = %q ff=%s allowed=%s", domain, vhost, forceFail, allowedSuffix)
				errorCode := http.StatusForbidden
				if forceFail != "" {
					errorCode, err = strconv.Atoi(forceFail)
					if err != nil {
						errorCode = http.StatusInternalServerError
					}
					logf("Forcing status code to %d", errorCode)
				}
				rw.WriteHeader(errorCode)
			}
		}))
	addr := srv.Listener.Addr().String()
	logf("Waiting for origin server at %s...", addr)
	if err := WaitForServer("tcp", addr, 10*time.Second); err != nil {
		return nil, "", err
	}
	logf("Started %s CDN", domain)
	return srv, addr, nil
}

func corruptMasquerades(cacheFile string) (int, error) {
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return 0, err
	}
	masquerades := make([]map[string]interface{}, 0)
	err = json.Unmarshal(data, &masquerades)
	if err != nil {
		return 0, err
	}
	for _, masquerade := range masquerades {
		ip := masquerade["IpAddress"]
		ipParts := strings.Split(ip.(string), ".")
		part0, _ := strconv.Atoi(ipParts[0])
		ipParts[0] = strconv.Itoa(part0 + 1)
		masquerade["IpAddress"] = strings.Join(ipParts, ".")
	}
	messedUp, err := json.Marshal(masquerades)
	if err != nil {
		return 0, err
	}

	return len(masquerades), os.WriteFile(cacheFile, messedUp, 0644)
}

func TestVerifyPeerCertificate(t *testing.T) {
	// raw certs generated by printing the received rawCerts from TestDirectDomainFrontingWithSNIConfig
	rawCerts := [][]byte{{48, 130, 5, 201, 48, 130, 5, 79, 160, 3, 2, 1, 2, 2, 16, 10, 104, 174, 217, 200, 97, 246, 117, 233, 137, 207, 239, 166, 251, 207, 173, 48, 10, 6, 8, 42, 134, 72, 206, 61, 4, 3, 3, 48, 86, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 48, 48, 46, 6, 3, 85, 4, 3, 19, 39, 68, 105, 103, 105, 67, 101, 114, 116, 32, 84, 76, 83, 32, 72, 121, 98, 114, 105, 100, 32, 69, 67, 67, 32, 83, 72, 65, 51, 56, 52, 32, 50, 48, 50, 48, 32, 67, 65, 49, 48, 30, 23, 13, 50, 53, 48, 51, 49, 56, 48, 48, 48, 48, 48, 48, 90, 23, 13, 50, 54, 48, 51, 49, 56, 50, 51, 53, 57, 53, 57, 90, 48, 121, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 22, 48, 20, 6, 3, 85, 4, 8, 19, 13, 77, 97, 115, 115, 97, 99, 104, 117, 115, 101, 116, 116, 115, 49, 18, 48, 16, 6, 3, 85, 4, 7, 19, 9, 67, 97, 109, 98, 114, 105, 100, 103, 101, 49, 34, 48, 32, 6, 3, 85, 4, 10, 19, 25, 65, 107, 97, 109, 97, 105, 32, 84, 101, 99, 104, 110, 111, 108, 111, 103, 105, 101, 115, 44, 32, 73, 110, 99, 46, 49, 26, 48, 24, 6, 3, 85, 4, 3, 19, 17, 97, 50, 52, 56, 46, 101, 46, 97, 107, 97, 109, 97, 105, 46, 110, 101, 116, 48, 89, 48, 19, 6, 7, 42, 134, 72, 206, 61, 2, 1, 6, 8, 42, 134, 72, 206, 61, 3, 1, 7, 3, 66, 0, 4, 137, 179, 173, 36, 7, 148, 115, 87, 89, 242, 68, 18, 67, 219, 1, 116, 170, 189, 59, 31, 134, 167, 129, 151, 124, 33, 237, 233, 140, 3, 235, 210, 253, 232, 46, 212, 95, 18, 33, 192, 156, 90, 87, 31, 233, 17, 0, 194, 110, 216, 164, 93, 216, 34, 45, 203, 136, 211, 69, 207, 17, 17, 92, 225, 163, 130, 3, 218, 48, 130, 3, 214, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 10, 188, 8, 41, 23, 140, 165, 57, 109, 122, 14, 206, 51, 199, 46, 179, 237, 251, 195, 122, 48, 29, 6, 3, 85, 29, 14, 4, 22, 4, 20, 84, 136, 107, 213, 162, 51, 87, 199, 43, 162, 19, 8, 242, 17, 240, 170, 80, 244, 78, 184, 48, 110, 6, 3, 85, 29, 17, 4, 103, 48, 101, 130, 17, 97, 50, 52, 56, 46, 101, 46, 97, 107, 97, 109, 97, 105, 46, 110, 101, 116, 130, 15, 42, 46, 97, 107, 97, 109, 97, 105, 122, 101, 100, 46, 110, 101, 116, 130, 23, 42, 46, 97, 107, 97, 109, 97, 105, 122, 101, 100, 45, 115, 116, 97, 103, 105, 110, 103, 46, 110, 101, 116, 130, 14, 42, 46, 97, 107, 97, 109, 97, 105, 104, 100, 46, 110, 101, 116, 130, 22, 42, 46, 97, 107, 97, 109, 97, 105, 104, 100, 45, 115, 116, 97, 103, 105, 110, 103, 46, 110, 101, 116, 48, 62, 6, 3, 85, 29, 32, 4, 55, 48, 53, 48, 51, 6, 6, 103, 129, 12, 1, 2, 2, 48, 41, 48, 39, 6, 8, 43, 6, 1, 5, 5, 7, 2, 1, 22, 27, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 67, 80, 83, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 3, 136, 48, 29, 6, 3, 85, 29, 37, 4, 22, 48, 20, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 6, 8, 43, 6, 1, 5, 5, 7, 3, 2, 48, 129, 155, 6, 3, 85, 29, 31, 4, 129, 147, 48, 129, 144, 48, 70, 160, 68, 160, 66, 134, 64, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 51, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 84, 76, 83, 72, 121, 98, 114, 105, 100, 69, 67, 67, 83, 72, 65, 51, 56, 52, 50, 48, 50, 48, 67, 65, 49, 45, 49, 46, 99, 114, 108, 48, 70, 160, 68, 160, 66, 134, 64, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 52, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 84, 76, 83, 72, 121, 98, 114, 105, 100, 69, 67, 67, 83, 72, 65, 51, 56, 52, 50, 48, 50, 48, 67, 65, 49, 45, 49, 46, 99, 114, 108, 48, 129, 133, 6, 8, 43, 6, 1, 5, 5, 7, 1, 1, 4, 121, 48, 119, 48, 36, 6, 8, 43, 6, 1, 5, 5, 7, 48, 1, 134, 24, 104, 116, 116, 112, 58, 47, 47, 111, 99, 115, 112, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 48, 79, 6, 8, 43, 6, 1, 5, 5, 7, 48, 2, 134, 67, 104, 116, 116, 112, 58, 47, 47, 99, 97, 99, 101, 114, 116, 115, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 84, 76, 83, 72, 121, 98, 114, 105, 100, 69, 67, 67, 83, 72, 65, 51, 56, 52, 50, 48, 50, 48, 67, 65, 49, 45, 49, 46, 99, 114, 116, 48, 12, 6, 3, 85, 29, 19, 1, 1, 255, 4, 2, 48, 0, 48, 130, 1, 127, 6, 10, 43, 6, 1, 4, 1, 214, 121, 2, 4, 2, 4, 130, 1, 111, 4, 130, 1, 107, 1, 105, 0, 119, 0, 14, 87, 148, 188, 243, 174, 169, 62, 51, 27, 44, 153, 7, 179, 247, 144, 223, 155, 194, 61, 113, 50, 37, 221, 33, 169, 37, 172, 97, 197, 78, 33, 0, 0, 1, 149, 170, 99, 183, 135, 0, 0, 4, 3, 0, 72, 48, 70, 2, 33, 0, 235, 0, 242, 233, 98, 214, 73, 85, 137, 193, 228, 5, 25, 30, 42, 94, 189, 169, 209, 120, 117, 2, 217, 215, 176, 205, 75, 242, 208, 168, 238, 17, 2, 33, 0, 198, 118, 215, 89, 151, 11, 1, 238, 193, 90, 1, 193, 70, 74, 22, 58, 217, 225, 104, 133, 172, 94, 20, 99, 124, 128, 254, 138, 79, 55, 54, 123, 0, 118, 0, 73, 156, 155, 105, 222, 29, 124, 236, 252, 54, 222, 205, 135, 100, 166, 184, 91, 175, 10, 135, 128, 25, 209, 85, 82, 251, 233, 235, 41, 221, 248, 195, 0, 0, 1, 149, 170, 99, 183, 201, 0, 0, 4, 3, 0, 71, 48, 69, 2, 33, 0, 214, 242, 234, 42, 116, 165, 172, 238, 126, 119, 200, 152, 85, 19, 124, 2, 180, 24, 8, 241, 20, 135, 91, 42, 14, 66, 55, 214, 136, 129, 23, 27, 2, 32, 113, 102, 37, 203, 235, 86, 220, 129, 181, 162, 156, 200, 129, 80, 125, 97, 82, 224, 31, 181, 77, 27, 136, 123, 70, 2, 130, 127, 3, 92, 16, 2, 0, 118, 0, 203, 56, 247, 21, 137, 124, 132, 161, 68, 95, 91, 193, 221, 251, 201, 110, 242, 154, 89, 205, 71, 10, 105, 5, 133, 176, 203, 20, 195, 20, 88, 231, 0, 0, 1, 149, 170, 99, 183, 157, 0, 0, 4, 3, 0, 71, 48, 69, 2, 32, 74, 67, 53, 194, 171, 185, 139, 172, 120, 13, 130, 224, 25, 132, 23, 1, 186, 169, 182, 193, 127, 84, 186, 57, 200, 221, 58, 255, 253, 149, 89, 203, 2, 33, 0, 222, 152, 192, 126, 210, 194, 215, 227, 190, 66, 199, 219, 75, 231, 144, 230, 249, 77, 28, 106, 124, 143, 55, 68, 212, 178, 188, 36, 28, 126, 135, 152, 48, 10, 6, 8, 42, 134, 72, 206, 61, 4, 3, 3, 3, 104, 0, 48, 101, 2, 48, 74, 189, 221, 93, 21, 162, 48, 200, 90, 164, 63, 233, 219, 209, 90, 6, 177, 129, 252, 19, 92, 4, 173, 221, 13, 8, 160, 9, 229, 206, 17, 220, 154, 107, 76, 136, 54, 79, 185, 131, 246, 235, 144, 87, 215, 248, 58, 246, 2, 49, 0, 179, 157, 243, 87, 121, 153, 72, 22, 90, 230, 197, 138, 127, 234, 20, 45, 23, 37, 48, 186, 234, 163, 23, 206, 103, 92, 5, 159, 100, 167, 161, 141, 193, 223, 215, 23, 62, 4, 91, 91, 103, 13, 39, 75, 50, 14, 60, 43}, {48, 130, 4, 23, 48, 130, 2, 255, 160, 3, 2, 1, 2, 2, 16, 7, 242, 243, 92, 135, 168, 119, 175, 122, 239, 233, 71, 153, 53, 37, 189, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 12, 5, 0, 48, 97, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 25, 48, 23, 6, 3, 85, 4, 11, 19, 16, 119, 119, 119, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 49, 32, 48, 30, 6, 3, 85, 4, 3, 19, 23, 68, 105, 103, 105, 67, 101, 114, 116, 32, 71, 108, 111, 98, 97, 108, 32, 82, 111, 111, 116, 32, 67, 65, 48, 30, 23, 13, 50, 49, 48, 52, 49, 52, 48, 48, 48, 48, 48, 48, 90, 23, 13, 51, 49, 48, 52, 49, 51, 50, 51, 53, 57, 53, 57, 90, 48, 86, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 48, 48, 46, 6, 3, 85, 4, 3, 19, 39, 68, 105, 103, 105, 67, 101, 114, 116, 32, 84, 76, 83, 32, 72, 121, 98, 114, 105, 100, 32, 69, 67, 67, 32, 83, 72, 65, 51, 56, 52, 32, 50, 48, 50, 48, 32, 67, 65, 49, 48, 118, 48, 16, 6, 7, 42, 134, 72, 206, 61, 2, 1, 6, 5, 43, 129, 4, 0, 34, 3, 98, 0, 4, 193, 27, 198, 154, 91, 152, 217, 164, 41, 160, 233, 212, 4, 181, 219, 235, 166, 178, 108, 85, 192, 255, 237, 152, 198, 73, 47, 6, 39, 81, 203, 191, 112, 193, 5, 122, 195, 177, 157, 135, 137, 186, 173, 180, 19, 23, 201, 168, 180, 131, 200, 184, 144, 209, 204, 116, 53, 54, 60, 131, 114, 176, 181, 208, 247, 34, 105, 200, 241, 128, 196, 123, 64, 143, 207, 104, 135, 38, 92, 57, 137, 241, 77, 145, 77, 218, 137, 139, 228, 3, 195, 67, 229, 191, 47, 115, 163, 130, 1, 130, 48, 130, 1, 126, 48, 18, 6, 3, 85, 29, 19, 1, 1, 255, 4, 8, 48, 6, 1, 1, 255, 2, 1, 0, 48, 29, 6, 3, 85, 29, 14, 4, 22, 4, 20, 10, 188, 8, 41, 23, 140, 165, 57, 109, 122, 14, 206, 51, 199, 46, 179, 237, 251, 195, 122, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 3, 222, 80, 53, 86, 209, 76, 187, 102, 240, 163, 226, 27, 27, 195, 151, 178, 61, 209, 85, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 1, 134, 48, 29, 6, 3, 85, 29, 37, 4, 22, 48, 20, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 6, 8, 43, 6, 1, 5, 5, 7, 3, 2, 48, 118, 6, 8, 43, 6, 1, 5, 5, 7, 1, 1, 4, 106, 48, 104, 48, 36, 6, 8, 43, 6, 1, 5, 5, 7, 48, 1, 134, 24, 104, 116, 116, 112, 58, 47, 47, 111, 99, 115, 112, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 48, 64, 6, 8, 43, 6, 1, 5, 5, 7, 48, 2, 134, 52, 104, 116, 116, 112, 58, 47, 47, 99, 97, 99, 101, 114, 116, 115, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 82, 111, 111, 116, 67, 65, 46, 99, 114, 116, 48, 66, 6, 3, 85, 29, 31, 4, 59, 48, 57, 48, 55, 160, 53, 160, 51, 134, 49, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 51, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 82, 111, 111, 116, 67, 65, 46, 99, 114, 108, 48, 61, 6, 3, 85, 29, 32, 4, 54, 48, 52, 48, 11, 6, 9, 96, 134, 72, 1, 134, 253, 108, 2, 1, 48, 7, 6, 5, 103, 129, 12, 1, 1, 48, 8, 6, 6, 103, 129, 12, 1, 2, 1, 48, 8, 6, 6, 103, 129, 12, 1, 2, 2, 48, 8, 6, 6, 103, 129, 12, 1, 2, 3, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 12, 5, 0, 3, 130, 1, 1, 0, 71, 89, 129, 127, 212, 27, 31, 176, 113, 246, 152, 93, 24, 186, 152, 71, 152, 176, 126, 118, 43, 234, 255, 26, 139, 172, 38, 179, 66, 141, 49, 230, 74, 232, 25, 208, 239, 218, 20, 231, 215, 20, 146, 161, 146, 242, 167, 46, 45, 175, 251, 29, 246, 251, 83, 176, 138, 63, 252, 216, 22, 10, 233, 176, 46, 182, 165, 11, 24, 144, 53, 38, 162, 218, 246, 168, 183, 50, 252, 149, 35, 75, 198, 69, 185, 196, 207, 228, 124, 238, 230, 201, 248, 144, 189, 114, 227, 153, 195, 29, 11, 5, 124, 106, 151, 109, 178, 171, 2, 54, 216, 194, 188, 44, 1, 146, 63, 4, 163, 139, 117, 17, 199, 185, 41, 188, 17, 208, 134, 186, 146, 188, 38, 249, 101, 200, 55, 205, 38, 246, 134, 19, 12, 4, 170, 137, 229, 120, 177, 193, 78, 121, 188, 118, 163, 11, 81, 228, 197, 208, 158, 106, 254, 26, 44, 86, 174, 6, 54, 39, 163, 115, 28, 8, 125, 147, 50, 208, 194, 68, 25, 218, 141, 244, 14, 123, 29, 40, 3, 43, 9, 138, 118, 202, 119, 220, 135, 122, 172, 123, 82, 38, 85, 167, 114, 15, 157, 210, 136, 79, 254, 177, 33, 197, 26, 161, 170, 57, 245, 86, 219, 194, 132, 196, 53, 31, 112, 218, 187, 70, 240, 134, 191, 100, 0, 196, 62, 247, 159, 70, 27, 157, 35, 5, 185, 125, 179, 79, 15, 169, 69, 58, 227, 116, 48, 152}}

	var tests = []struct {
		name            string
		givenRawCerts   [][]byte
		givenRoots      *x509.CertPool
		givenVerifyHost string
		assert          func(t *testing.T, err error)
	}{
		{
			name:          "should return no certificates present when not providing rawCerts",
			givenRawCerts: make([][]byte, 0),
			assert: func(t *testing.T, err error) {
				if assert.Error(t, err) {
					assert.ErrorContains(t, err, "no certificates presented")
				}
			},
		},
		{
			name:          "should return an error when providing invalid first rawCert",
			givenRawCerts: [][]byte{{}},
			assert: func(t *testing.T, err error) {
				if assert.Error(t, err) {
					assert.ErrorContains(t, err, "unable to parse certificate")
				}
			},
		},
		{
			name:          "should return an error when unable to load intermediate certificates",
			givenRawCerts: [][]byte{rawCerts[0], {}},
			assert: func(t *testing.T, err error) {
				if assert.Error(t, err) {
					assert.ErrorContains(t, err, "unable to parse intermediate certificate")
				}
			},
		},
		{
			name:            "should return an error when failing to verify the certificate for the given verifyHost",
			givenRawCerts:   rawCerts,
			givenRoots:      trustedCACerts(t),
			givenVerifyHost: "cloudfront.net",
			assert: func(t *testing.T, err error) {
				if assert.Error(t, err) {
					assert.ErrorContains(t, err, "certificate verification failed")
				}
			},
		},
		{
			name:            "should succeed when providing valid rawCerts, roots, domain and sni",
			givenRawCerts:   rawCerts,
			givenRoots:      trustedCACerts(t),
			givenVerifyHost: "potato.akamaihd.net",
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name:            "should succeed when providing valid rawCerts, roots even without verifying the host",
			givenRawCerts:   rawCerts,
			givenRoots:      trustedCACerts(t),
			givenVerifyHost: "",
			assert: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := verifyPeerCertificate(tt.givenRawCerts, tt.givenRoots, tt.givenVerifyHost)
			tt.assert(t, err)
		})
	}
}

func TestFindWorkingMasquerades(t *testing.T) {
	tests := []struct {
		name                string
		masquerades         []Front
		expectedSuccessful  int
		expectedMasquerades int
	}{
		{
			name: "All successful",
			masquerades: []Front{
				newMockFront("domain1.com", "1.1.1.1", 0, true),
				newMockFront("domain2.com", "2.2.2.2", 0, true),
				newMockFront("domain3.com", "3.3.3.3", 0, true),
				newMockFront("domain4.com", "4.4.4.4", 0, true),
				newMockFront("domain1.com", "1.1.1.1", 0, true),
				newMockFront("domain1.com", "1.1.1.1", 0, true),
			},
			expectedSuccessful: 4,
		},
		{
			name: "Some successful",
			masquerades: []Front{
				newMockFront("domain1.com", "1.1.1.1", 0, true),
				newMockFront("domain2.com", "2.2.2.2", 0, false),
				newMockFront("domain3.com", "3.3.3.3", 0, true),
				newMockFront("domain4.com", "4.4.4.4", 0, false),
				newMockFront("domain1.com", "1.1.1.1", 0, true),
			},
			expectedSuccessful: 2,
		},
		{
			name: "None successful",
			masquerades: []Front{
				newMockFront("domain1.com", "1.1.1.1", 0, false),
				newMockFront("domain2.com", "2.2.2.2", 0, false),
				newMockFront("domain3.com", "3.3.3.3", 0, false),
				newMockFront("domain4.com", "4.4.4.4", 0, false),
			},
			expectedSuccessful: 0,
		},
		{
			name: "Batch processing",
			masquerades: func() []Front {
				var masquerades []Front
				for i := 0; i < 50; i++ {
					masquerades = append(masquerades, newMockFront(fmt.Sprintf("domain%d.com", i), fmt.Sprintf("1.1.1.%d", i), 0, i%2 == 0))
				}
				return masquerades
			}(),
			expectedSuccessful: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fronted{
				stopCh:   make(chan interface{}, 10),
				frontsCh: make(chan Front, 10),
			}
			f.providers = make(map[string]*Provider)
			f.providers["testProviderId"] = NewProvider(nil, "", nil, nil, nil, nil, "")
			f.fronts = newThreadSafeFronts(len(tt.masquerades))
			f.fronts.addFronts(tt.masquerades...)

			f.tryAllFronts()

			tries := 0
			for len(f.frontsCh) < tt.expectedSuccessful && tries < 100 {
				time.Sleep(30 * time.Millisecond)
				tries++
			}

			assert.GreaterOrEqual(t, len(f.frontsCh), tt.expectedSuccessful)
		})
	}
}

func TestLoadFronts(t *testing.T) {
	providers := map[string]*Provider{
		"provider1": {
			Masquerades: []*Masquerade{
				{Domain: "domain1.com", IpAddress: "1.1.1.1"},
				{Domain: "domain2.com", IpAddress: "2.2.2.2"},
			},
		},
		"provider2": {
			Masquerades: []*Masquerade{
				{Domain: "domain3.com", IpAddress: "3.3.3.3"},
				{Domain: "domain4.com", IpAddress: "4.4.4.4"},
			},
		},
	}

	expected := map[string]bool{
		"domain1.com": true,
		"domain2.com": true,
		"domain3.com": true,
		"domain4.com": true,
	}

	// Create the cache dirty channel
	cacheDirty := make(chan interface{}, 10)
	masquerades := loadFronts(providers, cacheDirty)

	assert.Equal(t, 4, len(masquerades), "Unexpected number of masquerades loaded")

	for _, m := range masquerades {
		assert.True(t, expected[m.getDomain()], "Unexpected masquerade domain: %s", m.getDomain())
	}
}

func TestRandRange(t *testing.T) {
	tests := []struct {
		min, max int
	}{
		{1, 10},
		{5, 15},
		{0, 100},
		{-10, 10},
		{50, 60},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("min=%d,max=%d", tt.min, tt.max), func(t *testing.T) {
			for i := 0; i < 100; i++ {
				result := randRange(tt.min, tt.max)
				assert.GreaterOrEqual(t, result, tt.min)
				assert.Less(t, result, tt.max)
			}
		})
	}
}

// Generate a mock of a MasqueradeInterface with a Dial method that can optionally
// return an error after a specified number of milliseconds.
func newMockFront(domain string, ipAddress string, timeout time.Duration, passesCheck bool) *mockFront {
	return newMockFrontWithLastSuccess(domain, ipAddress, timeout, passesCheck, time.Time{})
}

// Generate a mock of a MasqueradeInterface with a Dial method that can optionally
// return an error after a specified number of milliseconds.
func newMockFrontWithLastSuccess(domain string, ipAddress string, timeout time.Duration, passesCheck bool, lastSucceededTime time.Time) *mockFront {
	return &mockFront{
		Domain:            domain,
		IpAddress:         ipAddress,
		timeout:           timeout,
		passesCheck:       passesCheck,
		lastSucceededTime: lastSucceededTime,
	}
}

type mockFront struct {
	Domain            string
	IpAddress         string
	timeout           time.Duration
	passesCheck       bool
	lastSucceededTime time.Time
}

// setLastSucceeded implements MasqueradeInterface.
func (m *mockFront) setLastSucceeded(succeededTime time.Time) {
	m.lastSucceededTime = succeededTime
}

// lastSucceeded implements MasqueradeInterface.
func (m *mockFront) lastSucceeded() time.Time {
	return m.lastSucceededTime
}

// isSucceeding implements MasqueradeInterface.
func (m *mockFront) isSucceeding() bool {
	return m.lastSucceededTime.After(time.Time{})
}

// verifyWithPost implements MasqueradeInterface.
func (m *mockFront) verifyWithPost(net.Conn, string) bool {
	return m.passesCheck
}

// dial implements MasqueradeInterface.
func (m *mockFront) dial(rootCAs *x509.CertPool, clientHelloID tls.ClientHelloID) (net.Conn, error) {
	if m.timeout > 0 {
		time.Sleep(m.timeout)
		return nil, errors.New("mock dial error")
	}
	m.lastSucceededTime = time.Now()
	return &net.TCPConn{}, nil
}

// getDomain implements MasqueradeInterface.
func (m *mockFront) getDomain() string {
	return m.Domain
}

// getIpAddress implements MasqueradeInterface.
func (m *mockFront) getIpAddress() string {
	return m.IpAddress
}

// getProviderID implements MasqueradeInterface.
func (m *mockFront) getProviderID() string {
	return "testProviderId"
}

// markFailed implements MasqueradeInterface.
func (m *mockFront) markFailed() {

}

// markSucceeded implements MasqueradeInterface.
func (m *mockFront) markSucceeded() {
}

// markCacheDirty implements MasqueradeInterface.
func (m *mockFront) markCacheDirty() {
}

func (m *mockFront) markWithResult(good bool) bool {
	if good {
		m.markSucceeded()
	} else {
		m.markFailed()
	}
	m.markCacheDirty()
	return good
}

// Make sure that the mockMasquerade implements the MasqueradeInterface
var _ Front = (*mockFront)(nil)
