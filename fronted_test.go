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
	yamlFile, err := os.ReadFile("fronted.yaml.gz")
	require.NoError(t, err)
	pool, providers, err := processYaml(yamlFile)
	require.NoError(t, err)
	require.NotNil(t, pool)
	require.NotNil(t, providers)

	// Make sure there are some providers
	assert.Greater(t, len(providers), 0)
}

func TestDomainFrontingWithoutSNIConfig(t *testing.T) {
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

func TestDomainFrontingWithSNIConfig(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		t.Skip("Skipping Akamai integration test in CI: real Akamai endpoints are unreliable from CI runners")
	}
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
	transport := NewFronted(WithCacheFile(cacheFile), WithCountryCode("test"), WithEmbeddedConfigName("noconfig.yaml"), WithDefaultProviderID("akamai"))
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
	transport := NewFronted(WithCacheFile(cacheFile), WithDefaultProviderID(testProviderID))
	transport.onNewFronts(certs, p)

	rt := newTransportFromDialer(transport)
	client := &http.Client{
		Transport: rt,
		Timeout:   5 * time.Second,
	}
	require.True(t, doCheck(client, http.MethodPost, http.StatusAccepted, pingURL))

	transport = NewFronted(WithCacheFile(cacheFile), WithDefaultProviderID(testProviderID))
	transport.onNewFronts(certs, p)
	client = &http.Client{
		Transport: newTransportFromDialer(transport),
	}
	require.True(t, doCheck(client, http.MethodGet, http.StatusOK, getURL))

	d := transport.(*fronted)

	// Check the number of masquerades at the end, waiting until we get the right number
	masqueradesAtEnd := 0
	for range 1000 {
		masqueradesAtEnd = len(d.fronts.fronts)
		if masqueradesAtEnd >= expectedMasqueradesAtEnd {
			break
		}
		time.Sleep(30 * time.Millisecond)
	}
	require.GreaterOrEqual(t, masqueradesAtEnd, expectedMasqueradesAtEnd)
	return masqueradesAtEnd
}

func TestVet(t *testing.T) {
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		t.Skip("Skipping integration test in CI: vets masquerades sequentially against real CDN endpoints")
	}
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
			`Get "http://fff.cloudsack.biz/foo": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for fff.cloudsack.biz`,
		},
		{
			"http://fff.cloudsack.biz:1234/bar?x=y&z=w",
			`Get "http://fff.cloudsack.biz:1234/bar?x=y&z=w": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for fff.cloudsack.biz`,
		},
		{
			"https://www.google.com",
			`Get "https://www.google.com": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for www.google.com`,
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

	rt := NewFronted(WithDefaultProviderID("cloudsack"))

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

	rt := NewFronted(WithDefaultProviderID("cloudsack"))
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
			`Get "http://www.notok.cloudsack.biz": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for www.notok.cloudsack.biz`,
		},
		{
			"http://ok.cloudsack.biz",
			`Get "http://ok.cloudsack.biz": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for ok.cloudsack.biz`,
		},
		{
			"http://www.abc.cloudsack.biz",
			`Get "http://www.abc.cloudsack.biz": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for www.abc.cloudsack.biz`,
		},
		{
			"http://noabc.cloudsack.biz",
			`Get "http://noabc.cloudsack.biz": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for noabc.cloudsack.biz`,
		},
		{
			"http://cloudsack.biz",
			`Get "http://cloudsack.biz": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for cloudsack.biz`,
		},
		{
			"https://www.google.com",
			`Get "https://www.google.com": no domain fronting mapping for 'cloudsack'. Please add it to provider_map.yaml or equivalent for www.google.com`,
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

	rt := NewFronted(WithDefaultProviderID("cloudsack"))
	rt.onNewFronts(certs, map[string]*Provider{"cloudsack": p})

	for _, test := range tests {
		log.Debug("Testing passthrough", "url", test.url, "headers", test.headers)
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
	// raw certs fetched from a248.e.akamai.net on 2026-03-19 (leaf valid 2025-12-22 to 2026-12-22)
	// To refresh: connect to a248.e.akamai.net:443 and extract PeerCertificates as raw bytes.
	rawCerts := [][]byte{{48, 130, 5, 201, 48, 130, 5, 79, 160, 3, 2, 1, 2, 2, 16, 3, 115, 171, 66, 15, 84, 148, 27, 85, 87, 66, 217, 172, 137, 6, 38, 48, 10, 6, 8, 42, 134, 72, 206, 61, 4, 3, 3, 48, 89, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 51, 48, 49, 6, 3, 85, 4, 3, 19, 42, 68, 105, 103, 105, 67, 101, 114, 116, 32, 71, 108, 111, 98, 97, 108, 32, 71, 51, 32, 84, 76, 83, 32, 69, 67, 67, 32, 83, 72, 65, 51, 56, 52, 32, 50, 48, 50, 48, 32, 67, 65, 49, 48, 30, 23, 13, 50, 53, 49, 50, 50, 50, 48, 48, 48, 48, 48, 48, 90, 23, 13, 50, 54, 49, 50, 50, 50, 50, 51, 53, 57, 53, 57, 90, 48, 121, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 22, 48, 20, 6, 3, 85, 4, 8, 19, 13, 77, 97, 115, 115, 97, 99, 104, 117, 115, 101, 116, 116, 115, 49, 18, 48, 16, 6, 3, 85, 4, 7, 19, 9, 67, 97, 109, 98, 114, 105, 100, 103, 101, 49, 34, 48, 32, 6, 3, 85, 4, 10, 19, 25, 65, 107, 97, 109, 97, 105, 32, 84, 101, 99, 104, 110, 111, 108, 111, 103, 105, 101, 115, 44, 32, 73, 110, 99, 46, 49, 26, 48, 24, 6, 3, 85, 4, 3, 19, 17, 97, 50, 52, 56, 46, 101, 46, 97, 107, 97, 109, 97, 105, 46, 110, 101, 116, 48, 89, 48, 19, 6, 7, 42, 134, 72, 206, 61, 2, 1, 6, 8, 42, 134, 72, 206, 61, 3, 1, 7, 3, 66, 0, 4, 111, 186, 198, 28, 140, 210, 95, 69, 166, 32, 10, 8, 148, 120, 2, 169, 163, 29, 116, 53, 247, 176, 207, 132, 247, 126, 133, 217, 90, 254, 197, 204, 161, 221, 162, 45, 40, 93, 124, 215, 173, 109, 242, 231, 189, 68, 138, 78, 158, 124, 200, 219, 211, 208, 130, 202, 71, 245, 147, 123, 110, 135, 176, 174, 163, 130, 3, 215, 48, 130, 3, 211, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 138, 35, 235, 158, 107, 215, 249, 55, 93, 249, 109, 33, 57, 118, 154, 161, 103, 222, 16, 168, 48, 29, 6, 3, 85, 29, 14, 4, 22, 4, 20, 175, 50, 71, 249, 169, 75, 98, 88, 227, 19, 56, 139, 138, 197, 234, 121, 107, 93, 97, 180, 48, 110, 6, 3, 85, 29, 17, 4, 103, 48, 101, 130, 17, 97, 50, 52, 56, 46, 101, 46, 97, 107, 97, 109, 97, 105, 46, 110, 101, 116, 130, 15, 42, 46, 97, 107, 97, 109, 97, 105, 122, 101, 100, 46, 110, 101, 116, 130, 23, 42, 46, 97, 107, 97, 109, 97, 105, 122, 101, 100, 45, 115, 116, 97, 103, 105, 110, 103, 46, 110, 101, 116, 130, 14, 42, 46, 97, 107, 97, 109, 97, 105, 104, 100, 46, 110, 101, 116, 130, 22, 42, 46, 97, 107, 97, 109, 97, 105, 104, 100, 45, 115, 116, 97, 103, 105, 110, 103, 46, 110, 101, 116, 48, 62, 6, 3, 85, 29, 32, 4, 55, 48, 53, 48, 51, 6, 6, 103, 129, 12, 1, 2, 2, 48, 41, 48, 39, 6, 8, 43, 6, 1, 5, 5, 7, 2, 1, 22, 27, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 67, 80, 83, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 3, 136, 48, 19, 6, 3, 85, 29, 37, 4, 12, 48, 10, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 48, 129, 159, 6, 3, 85, 29, 31, 4, 129, 151, 48, 129, 148, 48, 72, 160, 70, 160, 68, 134, 66, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 51, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 71, 51, 84, 76, 83, 69, 67, 67, 83, 72, 65, 51, 56, 52, 50, 48, 50, 48, 67, 65, 49, 45, 50, 46, 99, 114, 108, 48, 72, 160, 70, 160, 68, 134, 66, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 52, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 71, 51, 84, 76, 83, 69, 67, 67, 83, 72, 65, 51, 56, 52, 50, 48, 50, 48, 67, 65, 49, 45, 50, 46, 99, 114, 108, 48, 129, 135, 6, 8, 43, 6, 1, 5, 5, 7, 1, 1, 4, 123, 48, 121, 48, 36, 6, 8, 43, 6, 1, 5, 5, 7, 48, 1, 134, 24, 104, 116, 116, 112, 58, 47, 47, 111, 99, 115, 112, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 48, 81, 6, 8, 43, 6, 1, 5, 5, 7, 48, 2, 134, 69, 104, 116, 116, 112, 58, 47, 47, 99, 97, 99, 101, 114, 116, 115, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 71, 51, 84, 76, 83, 69, 67, 67, 83, 72, 65, 51, 56, 52, 50, 48, 50, 48, 67, 65, 49, 45, 50, 46, 99, 114, 116, 48, 12, 6, 3, 85, 29, 19, 1, 1, 255, 4, 2, 48, 0, 48, 130, 1, 128, 6, 10, 43, 6, 1, 4, 1, 214, 121, 2, 4, 2, 4, 130, 1, 112, 4, 130, 1, 108, 1, 106, 0, 119, 0, 216, 9, 85, 59, 148, 79, 122, 255, 200, 22, 25, 111, 148, 79, 133, 171, 176, 248, 252, 94, 135, 85, 38, 15, 21, 209, 46, 114, 187, 69, 75, 20, 0, 0, 1, 155, 71, 44, 132, 37, 0, 0, 4, 3, 0, 72, 48, 70, 2, 33, 0, 158, 239, 46, 20, 200, 22, 143, 122, 32, 169, 168, 36, 34, 171, 17, 160, 242, 99, 250, 113, 52, 152, 147, 215, 65, 170, 144, 115, 138, 2, 209, 143, 2, 33, 0, 241, 103, 177, 230, 12, 45, 78, 165, 173, 93, 247, 231, 142, 233, 1, 84, 78, 188, 128, 104, 48, 175, 194, 219, 199, 202, 188, 108, 185, 246, 140, 168, 0, 118, 0, 200, 163, 196, 127, 199, 179, 173, 185, 53, 107, 1, 63, 106, 122, 18, 109, 227, 58, 78, 67, 165, 198, 70, 249, 151, 173, 57, 117, 153, 29, 207, 154, 0, 0, 1, 155, 71, 44, 132, 94, 0, 0, 4, 3, 0, 71, 48, 69, 2, 32, 30, 43, 243, 178, 133, 25, 152, 197, 50, 214, 210, 84, 47, 137, 244, 144, 108, 208, 20, 112, 231, 241, 170, 95, 140, 160, 173, 229, 147, 32, 32, 130, 2, 33, 0, 244, 24, 46, 179, 140, 173, 98, 78, 147, 190, 97, 173, 133, 157, 31, 31, 40, 155, 114, 124, 249, 209, 107, 71, 77, 183, 52, 179, 54, 127, 243, 103, 0, 119, 0, 194, 49, 126, 87, 69, 25, 163, 69, 238, 127, 56, 222, 178, 144, 65, 235, 199, 194, 33, 90, 34, 191, 127, 213, 181, 173, 118, 154, 217, 14, 82, 205, 0, 0, 1, 155, 71, 44, 132, 51, 0, 0, 4, 3, 0, 72, 48, 70, 2, 33, 0, 136, 72, 13, 74, 213, 175, 80, 26, 100, 90, 38, 231, 231, 13, 200, 54, 16, 62, 160, 225, 74, 11, 232, 106, 210, 170, 57, 127, 106, 163, 3, 128, 2, 33, 0, 129, 9, 223, 236, 95, 171, 3, 167, 165, 117, 237, 65, 73, 5, 166, 58, 150, 57, 136, 11, 171, 61, 39, 189, 153, 90, 222, 175, 199, 66, 159, 172, 48, 10, 6, 8, 42, 134, 72, 206, 61, 4, 3, 3, 3, 104, 0, 48, 101, 2, 48, 107, 182, 71, 108, 4, 218, 17, 79, 182, 69, 42, 22, 248, 54, 241, 143, 118, 155, 201, 39, 83, 15, 165, 234, 140, 53, 63, 223, 164, 29, 44, 76, 81, 64, 204, 38, 27, 143, 88, 24, 224, 126, 22, 106, 173, 134, 123, 182, 2, 49, 0, 238, 81, 173, 172, 28, 31, 243, 138, 237, 192, 179, 6, 131, 198, 133, 126, 181, 63, 143, 84, 161, 243, 146, 74, 168, 108, 249, 164, 34, 232, 22, 87, 70, 121, 197, 36, 208, 94, 88, 253, 223, 101, 108, 73, 217, 244, 239, 225}, {48, 130, 3, 121, 48, 130, 2, 255, 160, 3, 2, 1, 2, 2, 16, 11, 0, 233, 45, 77, 109, 115, 31, 202, 48, 89, 199, 203, 30, 24, 134, 48, 10, 6, 8, 42, 134, 72, 206, 61, 4, 3, 3, 48, 97, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 25, 48, 23, 6, 3, 85, 4, 11, 19, 16, 119, 119, 119, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 49, 32, 48, 30, 6, 3, 85, 4, 3, 19, 23, 68, 105, 103, 105, 67, 101, 114, 116, 32, 71, 108, 111, 98, 97, 108, 32, 82, 111, 111, 116, 32, 71, 51, 48, 30, 23, 13, 50, 49, 48, 52, 49, 52, 48, 48, 48, 48, 48, 48, 90, 23, 13, 51, 49, 48, 52, 49, 51, 50, 51, 53, 57, 53, 57, 90, 48, 89, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 51, 48, 49, 6, 3, 85, 4, 3, 19, 42, 68, 105, 103, 105, 67, 101, 114, 116, 32, 71, 108, 111, 98, 97, 108, 32, 71, 51, 32, 84, 76, 83, 32, 69, 67, 67, 32, 83, 72, 65, 51, 56, 52, 32, 50, 48, 50, 48, 32, 67, 65, 49, 48, 118, 48, 16, 6, 7, 42, 134, 72, 206, 61, 2, 1, 6, 5, 43, 129, 4, 0, 34, 3, 98, 0, 4, 120, 169, 156, 117, 174, 136, 93, 99, 164, 173, 93, 134, 216, 16, 73, 214, 175, 146, 89, 99, 67, 35, 133, 244, 72, 101, 48, 205, 74, 52, 149, 166, 14, 62, 217, 124, 8, 215, 87, 5, 40, 72, 158, 11, 171, 235, 194, 211, 150, 158, 237, 69, 210, 139, 138, 206, 1, 75, 23, 67, 225, 115, 207, 109, 115, 72, 52, 220, 0, 70, 9, 181, 86, 84, 201, 95, 122, 199, 19, 7, 208, 108, 24, 23, 108, 202, 219, 199, 11, 38, 86, 46, 141, 7, 245, 103, 163, 130, 1, 130, 48, 130, 1, 126, 48, 18, 6, 3, 85, 29, 19, 1, 1, 255, 4, 8, 48, 6, 1, 1, 255, 2, 1, 0, 48, 29, 6, 3, 85, 29, 14, 4, 22, 4, 20, 138, 35, 235, 158, 107, 215, 249, 55, 93, 249, 109, 33, 57, 118, 154, 161, 103, 222, 16, 168, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 179, 219, 72, 164, 249, 161, 197, 216, 174, 54, 65, 204, 17, 99, 105, 98, 41, 188, 75, 198, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 1, 134, 48, 29, 6, 3, 85, 29, 37, 4, 22, 48, 20, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 6, 8, 43, 6, 1, 5, 5, 7, 3, 2, 48, 118, 6, 8, 43, 6, 1, 5, 5, 7, 1, 1, 4, 106, 48, 104, 48, 36, 6, 8, 43, 6, 1, 5, 5, 7, 48, 1, 134, 24, 104, 116, 116, 112, 58, 47, 47, 111, 99, 115, 112, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 48, 64, 6, 8, 43, 6, 1, 5, 5, 7, 48, 2, 134, 52, 104, 116, 116, 112, 58, 47, 47, 99, 97, 99, 101, 114, 116, 115, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 82, 111, 111, 116, 71, 51, 46, 99, 114, 116, 48, 66, 6, 3, 85, 29, 31, 4, 59, 48, 57, 48, 55, 160, 53, 160, 51, 134, 49, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 51, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 82, 111, 111, 116, 71, 51, 46, 99, 114, 108, 48, 61, 6, 3, 85, 29, 32, 4, 54, 48, 52, 48, 11, 6, 9, 96, 134, 72, 1, 134, 253, 108, 2, 1, 48, 7, 6, 5, 103, 129, 12, 1, 1, 48, 8, 6, 6, 103, 129, 12, 1, 2, 1, 48, 8, 6, 6, 103, 129, 12, 1, 2, 2, 48, 8, 6, 6, 103, 129, 12, 1, 2, 3, 48, 10, 6, 8, 42, 134, 72, 206, 61, 4, 3, 3, 3, 104, 0, 48, 101, 2, 48, 126, 38, 88, 110, 238, 136, 236, 12, 221, 21, 65, 238, 122, 184, 153, 153, 112, 209, 98, 101, 79, 160, 32, 158, 71, 177, 91, 193, 178, 103, 49, 29, 204, 114, 122, 175, 34, 114, 64, 66, 110, 101, 132, 254, 135, 75, 15, 25, 2, 49, 0, 230, 191, 214, 174, 52, 135, 91, 63, 103, 199, 29, 168, 111, 213, 18, 120, 181, 230, 135, 49, 68, 169, 93, 198, 184, 120, 204, 207, 239, 212, 50, 88, 17, 255, 58, 133, 6, 60, 29, 132, 111, 211, 245, 249, 218, 51, 28, 164}}

	// Check if the leaf cert has expired; skip with a clear message if so.
	if leafCert, err := x509.ParseCertificate(rawCerts[0]); err == nil {
		if time.Now().After(leafCert.NotAfter) {
			t.Skipf("Test leaf certificate expired on %s — fetch fresh certs from a248.e.akamai.net:443", leafCert.NotAfter.Format("2006-01-02"))
		}
	}

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
			expectedSuccessful: 2,
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
			expectedSuccessful: 2,
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
	masquerades := loadFronts(providers, cacheDirty, nil)

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

func TestWithDialer(t *testing.T) {
	called := false
	customDialer := func(ctx context.Context, network, addr string) (net.Conn, error) {
		called = true
		return (&net.Dialer{}).DialContext(ctx, network, addr)
	}

	f := NewFronted(
		WithDialer(customDialer),
		WithEmbeddedConfigName("noconfig.yaml"),
	)
	defer f.Close()

	d := f.(*fronted)
	assert.NotNil(t, d.dialFunc, "dialFunc should be set")

	// Verify the custom dialer is stored (we can't compare funcs directly, but we can
	// verify it's not the default by calling it and checking our flag).
	_, _ = d.dialFunc(context.Background(), "tcp", "localhost:0")
	assert.True(t, called, "custom dialer should have been called")
}

func TestWithDialerDefault(t *testing.T) {
	f := NewFronted(WithEmbeddedConfigName("noconfig.yaml"))
	defer f.Close()

	d := f.(*fronted)
	assert.NotNil(t, d.dialFunc, "default dialFunc should be set when WithDialer is not used")
}

func TestWithDialerFlowsToFronts(t *testing.T) {
	called := false
	customDialer := func(ctx context.Context, network, addr string) (net.Conn, error) {
		called = true
		return nil, errors.New("custom dialer called")
	}

	f := NewFronted(
		WithDialer(customDialer),
		WithEmbeddedConfigName("noconfig.yaml"),
	)
	defer f.Close()

	d := f.(*fronted)

	// Create a provider and fronts using the fronted's dialer
	masquerades := []*Masquerade{{Domain: "example.com", IpAddress: "127.0.0.1"}}
	providers := map[string]*Provider{
		"test": NewProvider(nil, "", masquerades, nil, nil, nil, ""),
	}
	fronts := loadFronts(providers, d.cacheDirty, d.dialFunc)
	assert.Equal(t, 1, len(fronts))

	// Dialing through the front should use the custom dialer
	_, err := fronts[0].dial(nil, tls.HelloChrome_131)
	assert.Error(t, err)
	assert.True(t, called, "custom dialer should flow through to fronts")
}
