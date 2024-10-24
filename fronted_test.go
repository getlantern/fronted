package fronted

import (
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
	"sync/atomic"
	"testing"
	"time"

	. "github.com/getlantern/waitforserver"
	tls "github.com/refraction-networking/utls"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDirectDomainFronting(t *testing.T) {
	dir, err := os.MkdirTemp("", "direct_test")
	require.NoError(t, err, "Unable to create temp dir")
	defer os.RemoveAll(dir)
	cacheFile := filepath.Join(dir, "cachefile.2")
	doTestDomainFronting(t, cacheFile, 10)
	time.Sleep(defaultCacheSaveInterval * 2)
	// Then try again, this time reusing the existing cacheFile but a corrupted version
	corruptMasquerades(cacheFile)
	doTestDomainFronting(t, cacheFile, 10)
}

func TestDirectDomainFrontingWithSNIConfig(t *testing.T) {
	dir, err := os.MkdirTemp("", "direct_test")
	require.NoError(t, err, "Unable to create temp dir")
	defer os.RemoveAll(dir)
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
	Configure(certs, p, testProviderID, cacheFile)

	transport, ok := NewFronted(0)
	require.True(t, ok)
	client := &http.Client{
		Transport: transport,
	}
	require.True(t, doCheck(client, http.MethodGet, http.StatusOK, getURL))
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
	Configure(certs, p, testProviderID, cacheFile)

	transport, ok := NewFronted(30 * time.Second)
	require.True(t, ok)

	client := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
	}
	require.True(t, doCheck(client, http.MethodPost, http.StatusAccepted, pingURL))

	transport, ok = NewFronted(0)
	require.True(t, ok)
	client = &http.Client{
		Transport: transport,
	}
	require.True(t, doCheck(client, http.MethodGet, http.StatusOK, getURL))

	instance, ok := DefaultContext.instance.Get(0)
	require.True(t, ok)
	d := instance.(*fronted)

	// Check the number of masquerades at the end, waiting up to 30 seconds until we get the right number
	masqueradesAtEnd := 0
	for i := 0; i < 100; i++ {
		masqueradesAtEnd = len(d.masquerades)
		if masqueradesAtEnd == expectedMasqueradesAtEnd {
			break
		}
		time.Sleep(300 * time.Millisecond)
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

func TestLoadCandidates(t *testing.T) {
	providers := testProviders()

	expected := make(map[Masquerade]bool)
	for _, p := range providers {
		for _, m := range p.Masquerades {
			expected[*m] = true
		}
	}

	d := &fronted{
		masquerades: make(sortedMasquerades, 0, len(expected)),
	}

	d.loadCandidates(providers)

	actual := make(map[Masquerade]bool)
	count := 0
	for _, m := range d.masquerades {
		actual[Masquerade{Domain: m.getDomain(), IpAddress: m.getIpAddress()}] = true
		count++
	}

	assert.Equal(t, len(DefaultCloudfrontMasquerades), count, "Unexpected number of candidates")
	assert.Equal(t, expected, actual, "Masquerades did not load as expected")
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

	cloudSack, cloudSackAddr, err := newCDN("cloudsack", "cloudsack.biz")
	if !assert.NoError(t, err, "failed to start cloudsack cdn") {
		return
	}
	defer cloudSack.Close()

	masq := []*Masquerade{{Domain: "example.com", IpAddress: cloudSackAddr}}
	alias := map[string]string{
		"abc.forbidden.com": "abc.cloudsack.biz",
		"def.forbidden.com": "def.cloudsack.biz",
	}
	p := NewProvider(alias, "https://ttt.cloudsack.biz/ping", masq, nil, nil, nil, nil)

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())
	Configure(certs, map[string]*Provider{"cloudsack": p}, "cloudsack", "")

	rt, ok := NewFronted(10 * time.Second)
	if !assert.True(t, ok, "failed to obtain direct roundtripper") {
		return
	}
	client := &http.Client{Transport: rt}
	for _, test := range tests {
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

	masq1 := []*Masquerade{{Domain: "example.com", IpAddress: cloudSackAddr}}
	alias1 := map[string]string{
		"abc.forbidden.com": "abc.cloudsack.biz",
		"def.forbidden.com": "def.cloudsack.biz",
	}
	p1 := NewProvider(alias1, "https://ttt.cloudsack.biz/ping", masq1, nil, nil, nil, nil)

	masq2 := []*Masquerade{{Domain: "example.com", IpAddress: sadCloudAddr}}
	alias2 := map[string]string{
		"abc.forbidden.com": "abc.sadcloud.io",
		"def.forbidden.com": "def.sadcloud.io",
	}
	p2 := NewProvider(alias2, "https://ttt.sadcloud.io/ping", masq2, nil, nil, nil, nil)

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())
	certs.AddCert(sadCloud.Certificate())

	providers := map[string]*Provider{
		"cloudsack": p1,
		"sadcloud":  p2,
	}

	Configure(certs, providers, "cloudsack", "")
	rt, ok := NewFronted(10 * time.Second)
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

	cloudSack, cloudSackAddr, err := newCDN("cloudsack", "cloudsack.biz")
	if !assert.NoError(t, err, "failed to start cloudsack cdn") {
		return
	}
	defer cloudSack.Close()

	masq := []*Masquerade{{Domain: "example.com", IpAddress: cloudSackAddr}}
	alias := map[string]string{}
	passthrough := []string{"*.ok.cloudsack.biz", "abc.cloudsack.biz"}
	p := NewProvider(alias, "https://ttt.cloudsack.biz/ping", masq, nil, passthrough, nil, nil)

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())
	Configure(certs, map[string]*Provider{"cloudsack": p}, "cloudsack", "")

	rt, ok := NewFronted(10 * time.Second)
	if !assert.True(t, ok, "failed to obtain direct roundtripper") {
		return
	}
	client := &http.Client{Transport: rt}
	for _, test := range tests {
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
		resp, err := client.Get(test.url)
		assert.EqualError(t, err, test.expectedError)
		assert.Nil(t, resp)

	}
}

func TestCustomValidators(t *testing.T) {

	sadCloud, sadCloudAddr, err := newCDN("sadcloud", "sadcloud.io")
	if !assert.NoError(t, err, "failed to start sadcloud cdn") {
		return
	}
	defer sadCloud.Close()

	sadCloudCodes := []int{http.StatusPaymentRequired, http.StatusTeapot, http.StatusBadGateway}
	sadCloudValidator := NewStatusCodeValidator(sadCloudCodes)
	testURL := "https://abc.forbidden.com/quux"

	setup := func(validator ResponseValidator) {
		masq := []*Masquerade{{Domain: "example.com", IpAddress: sadCloudAddr}}
		alias := map[string]string{
			"abc.forbidden.com": "abc.sadcloud.io",
		}
		p := NewProvider(alias, "https://ttt.sadcloud.io/ping", masq, validator, nil, nil, nil)

		certs := x509.NewCertPool()
		certs.AddCert(sadCloud.Certificate())

		providers := map[string]*Provider{
			"sadcloud": p,
		}

		Configure(certs, providers, "sadcloud", "")
	}

	// This error indicates that the validator has discarded all masquerades.
	// Each test starts with one masquerade, which is vetted during the
	// call to NewDirect.
	masqueradesExhausted := fmt.Sprintf(`Get "%v": could not complete request even with retries`, testURL)

	tests := []struct {
		responseCode  int
		validator     ResponseValidator
		expectedError string
	}{
		// with the default validator, only 403s are rejected
		{
			responseCode:  http.StatusForbidden,
			validator:     nil,
			expectedError: masqueradesExhausted,
		},
		{
			responseCode:  http.StatusAccepted,
			validator:     nil,
			expectedError: "",
		},
		{
			responseCode:  http.StatusPaymentRequired,
			validator:     nil,
			expectedError: "",
		},
		{
			responseCode:  http.StatusTeapot,
			validator:     nil,
			expectedError: "",
		},
		{
			responseCode:  http.StatusBadGateway,
			validator:     nil,
			expectedError: "",
		},

		// with the custom validator, 403 is allowed, listed codes are rejected
		{
			responseCode:  http.StatusForbidden,
			validator:     sadCloudValidator,
			expectedError: "",
		},
		{
			responseCode:  http.StatusAccepted,
			validator:     sadCloudValidator,
			expectedError: "",
		},
		{
			responseCode:  http.StatusPaymentRequired,
			validator:     sadCloudValidator,
			expectedError: masqueradesExhausted,
		},
		{
			responseCode:  http.StatusTeapot,
			validator:     sadCloudValidator,
			expectedError: masqueradesExhausted,
		},
		{
			responseCode:  http.StatusBadGateway,
			validator:     sadCloudValidator,
			expectedError: masqueradesExhausted,
		},
	}

	for _, test := range tests {
		setup(test.validator)
		direct, ok := NewFronted(1 * time.Second)
		if !assert.True(t, ok) {
			return
		}
		client := &http.Client{
			Transport: direct,
		}

		req, err := http.NewRequest(http.MethodGet, testURL, nil)
		if !assert.NoError(t, err) {
			return
		}
		if test.responseCode != http.StatusAccepted {
			val := strconv.Itoa(test.responseCode)
			log.Debugf("requesting forced response code %s", val)
			req.Header.Set(CDNForceFail, val)
		}

		res, err := client.Do(req)
		if test.expectedError == "" {
			if !assert.NoError(t, err) {
				continue
			}
			assert.Equal(t, test.responseCode, res.StatusCode, "Failed to force response status code")
		} else {
			assert.EqualError(t, err, test.expectedError)
		}
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

			forceFail := req.Header.Get(CDNForceFail)

			vhost := req.Host
			if vhost == domain || strings.HasSuffix(vhost, allowedSuffix) && forceFail == "" {
				log.Debugf("accepting request host=%s ff=%s", vhost, forceFail)
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
				log.Debugf("(%s) Rejecting request with host = %q ff=%s allowed=%s", domain, vhost, forceFail, allowedSuffix)
				errorCode := http.StatusForbidden
				if forceFail != "" {
					errorCode, err = strconv.Atoi(forceFail)
					if err != nil {
						errorCode = http.StatusInternalServerError
					}
					log.Debugf("Forcing status code to %d", errorCode)
				}
				rw.WriteHeader(errorCode)
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

func corruptMasquerades(cacheFile string) {
	log.Debug("Corrupting masquerades")
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		log.Error(err)
		return
	}
	masquerades := make([]map[string]interface{}, 0)
	err = json.Unmarshal(data, &masquerades)
	if err != nil {
		log.Error(err)
		return
	}
	log.Debugf("Number of masquerades to corrupt: %d", len(masquerades))
	for _, masquerade := range masquerades {
		domain := masquerade["Domain"]
		ip := masquerade["IpAddress"]
		ipParts := strings.Split(ip.(string), ".")
		part0, _ := strconv.Atoi(ipParts[0])
		ipParts[0] = strconv.Itoa(part0 + 1)
		masquerade["IpAddress"] = strings.Join(ipParts, ".")
		log.Debugf("Corrupted masquerade %v", domain)
	}
	messedUp, err := json.Marshal(masquerades)
	if err != nil {
		return
	}
	os.WriteFile(cacheFile, messedUp, 0644)
}

func TestVerifyPeerCertificate(t *testing.T) {
	// raw certs generated by printing the received rawCerts from TestDirectDomainFrontingWithSNIConfig
	rawCerts := [][]byte{{48, 130, 6, 78, 48, 130, 5, 54, 160, 3, 2, 1, 2, 2, 16, 11, 14, 250, 105, 152, 72, 112, 146, 165, 214, 78, 192, 231, 165, 110, 242, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 48, 79, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 41, 48, 39, 6, 3, 85, 4, 3, 19, 32, 68, 105, 103, 105, 67, 101, 114, 116, 32, 84, 76, 83, 32, 82, 83, 65, 32, 83, 72, 65, 50, 53, 54, 32, 50, 48, 50, 48, 32, 67, 65, 49, 48, 30, 23, 13, 50, 52, 48, 52, 49, 56, 48, 48, 48, 48, 48, 48, 90, 23, 13, 50, 53, 48, 52, 49, 57, 50, 51, 53, 57, 53, 57, 90, 48, 121, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 22, 48, 20, 6, 3, 85, 4, 8, 19, 13, 77, 97, 115, 115, 97, 99, 104, 117, 115, 101, 116, 116, 115, 49, 18, 48, 16, 6, 3, 85, 4, 7, 19, 9, 67, 97, 109, 98, 114, 105, 100, 103, 101, 49, 34, 48, 32, 6, 3, 85, 4, 10, 19, 25, 65, 107, 97, 109, 97, 105, 32, 84, 101, 99, 104, 110, 111, 108, 111, 103, 105, 101, 115, 44, 32, 73, 110, 99, 46, 49, 26, 48, 24, 6, 3, 85, 4, 3, 19, 17, 97, 50, 52, 56, 46, 101, 46, 97, 107, 97, 109, 97, 105, 46, 110, 101, 116, 48, 89, 48, 19, 6, 7, 42, 134, 72, 206, 61, 2, 1, 6, 8, 42, 134, 72, 206, 61, 3, 1, 7, 3, 66, 0, 4, 5, 224, 177, 69, 53, 250, 80, 142, 150, 138, 229, 168, 82, 249, 163, 196, 35, 150, 140, 182, 86, 208, 48, 132, 211, 49, 12, 169, 58, 148, 19, 105, 223, 193, 88, 236, 160, 208, 199, 150, 32, 252, 119, 75, 85, 5, 247, 130, 138, 242, 186, 184, 107, 67, 177, 230, 40, 36, 104, 131, 178, 228, 231, 148, 163, 130, 3, 197, 48, 130, 3, 193, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 183, 107, 162, 234, 168, 170, 132, 140, 121, 234, 180, 218, 15, 152, 178, 197, 149, 118, 185, 244, 48, 29, 6, 3, 85, 29, 14, 4, 22, 4, 20, 115, 183, 92, 115, 61, 0, 51, 82, 107, 67, 69, 86, 236, 116, 51, 65, 161, 9, 34, 162, 48, 110, 6, 3, 85, 29, 17, 4, 103, 48, 101, 130, 17, 97, 50, 52, 56, 46, 101, 46, 97, 107, 97, 109, 97, 105, 46, 110, 101, 116, 130, 15, 42, 46, 97, 107, 97, 109, 97, 105, 122, 101, 100, 46, 110, 101, 116, 130, 23, 42, 46, 97, 107, 97, 109, 97, 105, 122, 101, 100, 45, 115, 116, 97, 103, 105, 110, 103, 46, 110, 101, 116, 130, 14, 42, 46, 97, 107, 97, 109, 97, 105, 104, 100, 46, 110, 101, 116, 130, 22, 42, 46, 97, 107, 97, 109, 97, 105, 104, 100, 45, 115, 116, 97, 103, 105, 110, 103, 46, 110, 101, 116, 48, 62, 6, 3, 85, 29, 32, 4, 55, 48, 53, 48, 51, 6, 6, 103, 129, 12, 1, 2, 2, 48, 41, 48, 39, 6, 8, 43, 6, 1, 5, 5, 7, 2, 1, 22, 27, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 67, 80, 83, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 3, 136, 48, 29, 6, 3, 85, 29, 37, 4, 22, 48, 20, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 6, 8, 43, 6, 1, 5, 5, 7, 3, 2, 48, 129, 143, 6, 3, 85, 29, 31, 4, 129, 135, 48, 129, 132, 48, 64, 160, 62, 160, 60, 134, 58, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 51, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 84, 76, 83, 82, 83, 65, 83, 72, 65, 50, 53, 54, 50, 48, 50, 48, 67, 65, 49, 45, 52, 46, 99, 114, 108, 48, 64, 160, 62, 160, 60, 134, 58, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 52, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 84, 76, 83, 82, 83, 65, 83, 72, 65, 50, 53, 54, 50, 48, 50, 48, 67, 65, 49, 45, 52, 46, 99, 114, 108, 48, 127, 6, 8, 43, 6, 1, 5, 5, 7, 1, 1, 4, 115, 48, 113, 48, 36, 6, 8, 43, 6, 1, 5, 5, 7, 48, 1, 134, 24, 104, 116, 116, 112, 58, 47, 47, 111, 99, 115, 112, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 48, 73, 6, 8, 43, 6, 1, 5, 5, 7, 48, 2, 134, 61, 104, 116, 116, 112, 58, 47, 47, 99, 97, 99, 101, 114, 116, 115, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 84, 76, 83, 82, 83, 65, 83, 72, 65, 50, 53, 54, 50, 48, 50, 48, 67, 65, 49, 45, 49, 46, 99, 114, 116, 48, 12, 6, 3, 85, 29, 19, 1, 1, 255, 4, 2, 48, 0, 48, 130, 1, 125, 6, 10, 43, 6, 1, 4, 1, 214, 121, 2, 4, 2, 4, 130, 1, 109, 4, 130, 1, 105, 1, 103, 0, 118, 0, 78, 117, 163, 39, 92, 154, 16, 195, 56, 91, 108, 212, 223, 63, 82, 235, 29, 240, 224, 142, 27, 141, 105, 192, 177, 250, 100, 177, 98, 154, 57, 223, 0, 0, 1, 142, 241, 217, 223, 134, 0, 0, 4, 3, 0, 71, 48, 69, 2, 33, 0, 182, 60, 198, 96, 136, 128, 205, 139, 42, 82, 117, 248, 90, 158, 186, 210, 179, 163, 225, 68, 48, 33, 54, 42, 66, 129, 205, 220, 227, 47, 241, 24, 2, 32, 47, 50, 19, 81, 103, 101, 88, 38, 67, 79, 20, 225, 232, 59, 123, 77, 100, 243, 60, 99, 22, 213, 169, 109, 122, 35, 153, 88, 59, 40, 193, 180, 0, 118, 0, 125, 89, 30, 18, 225, 120, 42, 123, 28, 97, 103, 124, 94, 253, 248, 208, 135, 92, 20, 160, 78, 149, 158, 185, 3, 47, 217, 14, 140, 46, 121, 184, 0, 0, 1, 142, 241, 217, 223, 135, 0, 0, 4, 3, 0, 71, 48, 69, 2, 33, 0, 236, 206, 233, 76, 152, 193, 240, 13, 15, 141, 73, 58, 88, 53, 123, 217, 228, 185, 26, 35, 9, 53, 191, 231, 1, 223, 99, 28, 200, 188, 2, 47, 2, 32, 39, 67, 173, 42, 123, 38, 247, 178, 220, 3, 89, 37, 218, 105, 45, 249, 17, 111, 222, 84, 173, 197, 17, 26, 177, 217, 193, 163, 221, 229, 129, 134, 0, 117, 0, 230, 210, 49, 99, 64, 119, 140, 193, 16, 65, 6, 215, 113, 185, 206, 193, 210, 64, 246, 150, 132, 134, 251, 186, 135, 50, 29, 253, 30, 55, 142, 80, 0, 0, 1, 142, 241, 217, 223, 156, 0, 0, 4, 3, 0, 70, 48, 68, 2, 32, 63, 238, 16, 71, 200, 160, 240, 218, 87, 96, 100, 137, 184, 151, 189, 202, 191, 140, 193, 138, 110, 83, 166, 225, 152, 192, 33, 228, 72, 60, 146, 9, 2, 32, 20, 216, 203, 133, 251, 181, 154, 237, 126, 11, 120, 77, 219, 28, 73, 93, 254, 23, 141, 52, 195, 145, 216, 145, 16, 187, 133, 16, 140, 184, 135, 183, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 3, 130, 1, 1, 0, 98, 147, 27, 116, 164, 135, 78, 19, 1, 11, 53, 227, 221, 49, 154, 147, 19, 174, 118, 228, 188, 90, 81, 60, 70, 72, 54, 95, 222, 204, 55, 191, 171, 254, 126, 228, 34, 208, 165, 74, 135, 252, 133, 131, 205, 71, 216, 124, 81, 208, 146, 28, 219, 168, 108, 81, 76, 30, 114, 121, 71, 134, 116, 156, 58, 85, 38, 176, 202, 33, 124, 189, 155, 252, 217, 111, 116, 7, 83, 186, 149, 7, 7, 127, 39, 167, 50, 69, 97, 162, 65, 90, 234, 59, 114, 92, 19, 87, 118, 143, 216, 97, 192, 226, 95, 230, 244, 208, 237, 199, 7, 3, 99, 108, 69, 214, 95, 36, 69, 116, 75, 195, 254, 18, 207, 11, 34, 253, 237, 248, 127, 152, 29, 58, 131, 49, 178, 141, 72, 111, 11, 151, 30, 3, 56, 6, 6, 156, 45, 103, 3, 25, 210, 95, 235, 109, 29, 45, 59, 21, 36, 81, 146, 160, 165, 185, 201, 100, 150, 126, 160, 230, 126, 128, 222, 243, 49, 119, 188, 163, 162, 98, 153, 174, 185, 234, 44, 226, 102, 184, 207, 2, 193, 66, 77, 199, 39, 219, 64, 44, 145, 6, 207, 52, 237, 50, 200, 55, 253, 21, 208, 124, 150, 3, 136, 196, 70, 121, 86, 75, 41, 76, 71, 193, 94, 73, 151, 255, 164, 127, 129, 242, 35, 125, 80, 24, 21, 121, 184, 18, 224, 212, 70, 58, 206, 122, 34, 250, 119, 203, 84, 55, 11, 9, 221, 103}, {48, 130, 4, 190, 48, 130, 3, 166, 160, 3, 2, 1, 2, 2, 16, 6, 216, 217, 4, 213, 88, 67, 70, 246, 138, 47, 167, 84, 34, 126, 196, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 48, 97, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 25, 48, 23, 6, 3, 85, 4, 11, 19, 16, 119, 119, 119, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 49, 32, 48, 30, 6, 3, 85, 4, 3, 19, 23, 68, 105, 103, 105, 67, 101, 114, 116, 32, 71, 108, 111, 98, 97, 108, 32, 82, 111, 111, 116, 32, 67, 65, 48, 30, 23, 13, 50, 49, 48, 52, 49, 52, 48, 48, 48, 48, 48, 48, 90, 23, 13, 51, 49, 48, 52, 49, 51, 50, 51, 53, 57, 53, 57, 90, 48, 79, 49, 11, 48, 9, 6, 3, 85, 4, 6, 19, 2, 85, 83, 49, 21, 48, 19, 6, 3, 85, 4, 10, 19, 12, 68, 105, 103, 105, 67, 101, 114, 116, 32, 73, 110, 99, 49, 41, 48, 39, 6, 3, 85, 4, 3, 19, 32, 68, 105, 103, 105, 67, 101, 114, 116, 32, 84, 76, 83, 32, 82, 83, 65, 32, 83, 72, 65, 50, 53, 54, 32, 50, 48, 50, 48, 32, 67, 65, 49, 48, 130, 1, 34, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 1, 5, 0, 3, 130, 1, 15, 0, 48, 130, 1, 10, 2, 130, 1, 1, 0, 193, 75, 179, 101, 71, 112, 188, 221, 79, 88, 219, 236, 156, 237, 195, 102, 229, 31, 49, 19, 84, 173, 74, 102, 70, 31, 44, 10, 236, 100, 7, 229, 46, 220, 220, 185, 10, 32, 237, 223, 227, 196, 208, 158, 154, 169, 122, 29, 130, 136, 229, 17, 86, 219, 30, 159, 88, 194, 81, 231, 44, 52, 13, 46, 210, 146, 225, 86, 203, 241, 121, 95, 179, 187, 135, 202, 37, 3, 123, 154, 82, 65, 102, 16, 96, 79, 87, 19, 73, 240, 232, 55, 103, 131, 223, 231, 211, 75, 103, 76, 34, 81, 166, 223, 14, 153, 16, 237, 87, 81, 116, 38, 226, 125, 199, 202, 98, 46, 19, 27, 127, 35, 136, 37, 83, 111, 193, 52, 88, 0, 139, 132, 255, 248, 190, 167, 88, 73, 34, 123, 150, 173, 162, 136, 155, 21, 188, 160, 124, 223, 233, 81, 168, 213, 176, 237, 55, 226, 54, 180, 130, 75, 98, 181, 73, 154, 236, 199, 103, 214, 227, 62, 245, 227, 214, 18, 94, 68, 241, 191, 113, 66, 125, 88, 132, 3, 128, 177, 129, 1, 250, 249, 202, 50, 187, 180, 142, 39, 135, 39, 197, 43, 116, 212, 168, 214, 151, 222, 195, 100, 249, 202, 206, 83, 162, 86, 188, 120, 23, 142, 73, 3, 41, 174, 251, 73, 79, 164, 21, 185, 206, 242, 92, 25, 87, 109, 107, 121, 167, 43, 162, 39, 32, 19, 181, 208, 61, 64, 211, 33, 48, 7, 147, 234, 153, 245, 2, 3, 1, 0, 1, 163, 130, 1, 130, 48, 130, 1, 126, 48, 18, 6, 3, 85, 29, 19, 1, 1, 255, 4, 8, 48, 6, 1, 1, 255, 2, 1, 0, 48, 29, 6, 3, 85, 29, 14, 4, 22, 4, 20, 183, 107, 162, 234, 168, 170, 132, 140, 121, 234, 180, 218, 15, 152, 178, 197, 149, 118, 185, 244, 48, 31, 6, 3, 85, 29, 35, 4, 24, 48, 22, 128, 20, 3, 222, 80, 53, 86, 209, 76, 187, 102, 240, 163, 226, 27, 27, 195, 151, 178, 61, 209, 85, 48, 14, 6, 3, 85, 29, 15, 1, 1, 255, 4, 4, 3, 2, 1, 134, 48, 29, 6, 3, 85, 29, 37, 4, 22, 48, 20, 6, 8, 43, 6, 1, 5, 5, 7, 3, 1, 6, 8, 43, 6, 1, 5, 5, 7, 3, 2, 48, 118, 6, 8, 43, 6, 1, 5, 5, 7, 1, 1, 4, 106, 48, 104, 48, 36, 6, 8, 43, 6, 1, 5, 5, 7, 48, 1, 134, 24, 104, 116, 116, 112, 58, 47, 47, 111, 99, 115, 112, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 48, 64, 6, 8, 43, 6, 1, 5, 5, 7, 48, 2, 134, 52, 104, 116, 116, 112, 58, 47, 47, 99, 97, 99, 101, 114, 116, 115, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 82, 111, 111, 116, 67, 65, 46, 99, 114, 116, 48, 66, 6, 3, 85, 29, 31, 4, 59, 48, 57, 48, 55, 160, 53, 160, 51, 134, 49, 104, 116, 116, 112, 58, 47, 47, 99, 114, 108, 51, 46, 100, 105, 103, 105, 99, 101, 114, 116, 46, 99, 111, 109, 47, 68, 105, 103, 105, 67, 101, 114, 116, 71, 108, 111, 98, 97, 108, 82, 111, 111, 116, 67, 65, 46, 99, 114, 108, 48, 61, 6, 3, 85, 29, 32, 4, 54, 48, 52, 48, 11, 6, 9, 96, 134, 72, 1, 134, 253, 108, 2, 1, 48, 7, 6, 5, 103, 129, 12, 1, 1, 48, 8, 6, 6, 103, 129, 12, 1, 2, 1, 48, 8, 6, 6, 103, 129, 12, 1, 2, 2, 48, 8, 6, 6, 103, 129, 12, 1, 2, 3, 48, 13, 6, 9, 42, 134, 72, 134, 247, 13, 1, 1, 11, 5, 0, 3, 130, 1, 1, 0, 128, 50, 206, 94, 11, 221, 110, 90, 13, 10, 175, 225, 214, 132, 203, 192, 142, 250, 133, 112, 237, 218, 93, 179, 12, 247, 43, 117, 64, 254, 133, 10, 250, 243, 49, 120, 183, 112, 75, 26, 137, 88, 186, 128, 189, 243, 107, 29, 233, 126, 207, 11, 186, 88, 156, 89, 212, 144, 211, 253, 108, 253, 208, 152, 109, 183, 113, 130, 91, 207, 109, 11, 90, 9, 208, 123, 222, 196, 67, 216, 42, 164, 222, 158, 65, 38, 95, 187, 143, 153, 203, 221, 174, 225, 168, 111, 159, 135, 254, 116, 183, 31, 27, 32, 171, 177, 79, 198, 245, 103, 93, 93, 155, 60, 233, 255, 105, 247, 97, 108, 214, 217, 243, 253, 54, 198, 171, 3, 136, 118, 210, 75, 46, 117, 134, 227, 252, 216, 85, 125, 38, 194, 17, 119, 223, 62, 2, 182, 124, 243, 171, 123, 122, 134, 54, 111, 184, 247, 216, 147, 113, 207, 134, 223, 115, 48, 250, 123, 171, 237, 42, 89, 200, 66, 132, 59, 17, 23, 26, 82, 243, 201, 14, 20, 125, 162, 91, 114, 103, 186, 113, 237, 87, 71, 102, 197, 184, 2, 74, 101, 52, 94, 139, 208, 42, 60, 32, 156, 81, 153, 76, 231, 82, 158, 247, 107, 17, 43, 13, 146, 126, 29, 232, 138, 235, 54, 22, 67, 135, 234, 42, 99, 191, 117, 63, 235, 222, 196, 3, 187, 10, 60, 247, 48, 239, 235, 175, 76, 252, 139, 54, 16, 115, 62, 243, 164}}

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
		masquerades         []*mockMasquerade
		expectedSuccessful  int
		expectedMasquerades int
	}{
		{
			name: "All successful",
			masquerades: []*mockMasquerade{
				newMockMasquerade("domain1.com", "1.1.1.1", 0, true),
				newMockMasquerade("domain2.com", "2.2.2.2", 0, true),
				newMockMasquerade("domain3.com", "3.3.3.3", 0, true),
				newMockMasquerade("domain4.com", "4.4.4.4", 0, true),
				newMockMasquerade("domain1.com", "1.1.1.1", 0, true),
				newMockMasquerade("domain1.com", "1.1.1.1", 0, true),
			},
			expectedSuccessful: 4,
		},
		{
			name: "Some successful",
			masquerades: []*mockMasquerade{
				newMockMasquerade("domain1.com", "1.1.1.1", 0, true),
				newMockMasquerade("domain2.com", "2.2.2.2", 0, false),
				newMockMasquerade("domain3.com", "3.3.3.3", 0, true),
				newMockMasquerade("domain4.com", "4.4.4.4", 0, false),
				newMockMasquerade("domain1.com", "1.1.1.1", 0, true),
			},
			expectedSuccessful: 2,
		},
		{
			name: "None successful",
			masquerades: []*mockMasquerade{
				newMockMasquerade("domain1.com", "1.1.1.1", 0, false),
				newMockMasquerade("domain2.com", "2.2.2.2", 0, false),
				newMockMasquerade("domain3.com", "3.3.3.3", 0, false),
				newMockMasquerade("domain4.com", "4.4.4.4", 0, false),
			},
			expectedSuccessful: 0,
		},
		{
			name: "Batch processing",
			masquerades: func() []*mockMasquerade {
				var masquerades []*mockMasquerade
				for i := 0; i < 50; i++ {
					masquerades = append(masquerades, newMockMasquerade(fmt.Sprintf("domain%d.com", i), fmt.Sprintf("1.1.1.%d", i), 0, i%2 == 0))
				}
				return masquerades
			}(),
			expectedSuccessful: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &fronted{}
			d.providers = make(map[string]*Provider)
			d.providers["testProviderId"] = NewProvider(nil, "", nil, nil, nil, nil, nil)
			d.masquerades = make(sortedMasquerades, len(tt.masquerades))
			for i, m := range tt.masquerades {
				d.masquerades[i] = m
			}

			var successful atomic.Uint32
			d.vetBatch(0, 10, &successful)

			tries := 0
			for successful.Load() < uint32(tt.expectedSuccessful) && tries < 100 {
				time.Sleep(30 * time.Millisecond)
				tries++
			}

			assert.GreaterOrEqual(t, int(successful.Load()), tt.expectedSuccessful)
		})
	}
}

// Generate a mock of a MasqueradeInterface with a Dial method that can optionally
// return an error after a specified number of milliseconds.
func newMockMasquerade(domain string, ipAddress string, timeout time.Duration, passesCheck bool) *mockMasquerade {
	return &mockMasquerade{
		Domain:      domain,
		IpAddress:   ipAddress,
		timeout:     timeout,
		passesCheck: passesCheck,
	}
}

type mockMasquerade struct {
	Domain            string
	IpAddress         string
	timeout           time.Duration
	passesCheck       bool
	lastSucceededTime time.Time
}

// setLastSucceeded implements MasqueradeInterface.
func (m *mockMasquerade) setLastSucceeded(succeededTime time.Time) {
	m.lastSucceededTime = succeededTime
}

// lastSucceeded implements MasqueradeInterface.
func (m *mockMasquerade) lastSucceeded() time.Time {
	return m.lastSucceededTime
}

// postCheck implements MasqueradeInterface.
func (m *mockMasquerade) postCheck(net.Conn, string) bool {
	return m.passesCheck
}

// dial implements MasqueradeInterface.
func (m *mockMasquerade) dial(rootCAs *x509.CertPool, clientHelloID tls.ClientHelloID) (net.Conn, error) {
	if m.timeout > 0 {
		time.Sleep(m.timeout)
		return nil, errors.New("mock dial error")
	}
	m.lastSucceededTime = time.Now()
	return &net.TCPConn{}, nil
}

// getDomain implements MasqueradeInterface.
func (m *mockMasquerade) getDomain() string {
	return m.Domain
}

// getIpAddress implements MasqueradeInterface.
func (m *mockMasquerade) getIpAddress() string {
	return m.IpAddress
}

// getProviderID implements MasqueradeInterface.
func (m *mockMasquerade) getProviderID() string {
	return "testProviderId"
}

// markFailed implements MasqueradeInterface.
func (m *mockMasquerade) markFailed() {

}

// markSucceeded implements MasqueradeInterface.
func (m *mockMasquerade) markSucceeded() {
}

// Make sure that the mockMasquerade implements the MasqueradeInterface
var _ MasqueradeInterface = (*mockMasquerade)(nil)
