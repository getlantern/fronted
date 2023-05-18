package fronted

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/getlantern/waitforserver"
)

func TestDirectDomainFronting(t *testing.T) {
	dir, err := ioutil.TempDir("", "direct_test")
	require.NoError(t, err, "Unable to create temp dir")
	defer os.RemoveAll(dir)
	cacheFile := filepath.Join(dir, "cachefile.2")
	doTestDomainFronting(t, cacheFile, numberToVetInitially)
	time.Sleep(defaultCacheSaveInterval * 2)
	// Then try again, this time reusing the existing cacheFile but a corrupted version
	corruptMasquerades(cacheFile)
	doTestDomainFronting(t, cacheFile, numberToVetInitially)
}

func doTestDomainFronting(t *testing.T, cacheFile string, expectedMasqueradesAtEnd int) int {

	getURL := "http://config.example.com/proxies.yaml.gz"
	getHost := "config.example.com"
	getFrontedHost := "d2wi0vwulmtn99.cloudfront.net"

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

	transport, ok := NewDirect(30 * time.Second)
	require.True(t, ok)

	client := &http.Client{
		Transport: transport,
	}
	require.True(t, doCheck(client, http.MethodPost, http.StatusAccepted, pingURL))

	transport, ok = NewDirect(0)
	require.True(t, ok)
	client = &http.Client{
		Transport: transport,
	}
	require.True(t, doCheck(client, http.MethodGet, http.StatusOK, getURL))

	instance, ok := DefaultContext.instance.Get(0)
	require.True(t, ok)
	d := instance.(*direct)

	// Check the number of masquerades at the end, waiting up to 30 seconds until we get the right number
	masqueradesAtEnd := 0
	for i := 0; i < 100; i++ {
		masqueradesAtEnd = len(d.masquerades)
		if masqueradesAtEnd == expectedMasqueradesAtEnd {
			break
		}
		time.Sleep(300 * time.Millisecond)
	}
	require.Equal(t, expectedMasqueradesAtEnd, masqueradesAtEnd)
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
	p := NewProvider(alias, "https://ttt.cloudsack.biz/ping", masq, nil, nil)

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())
	Configure(certs, map[string]*Provider{"cloudsack": p}, "cloudsack", "")

	rt, ok := NewDirect(10 * time.Second)
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
	p1 := NewProvider(alias1, "https://ttt.cloudsack.biz/ping", masq1, nil, nil)

	masq2 := []*Masquerade{{Domain: "example.com", IpAddress: sadCloudAddr}}
	alias2 := map[string]string{
		"abc.forbidden.com": "abc.sadcloud.io",
		"def.forbidden.com": "def.sadcloud.io",
	}
	p2 := NewProvider(alias2, "https://ttt.sadcloud.io/ping", masq2, nil, nil)

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
	p := NewProvider(alias, "https://ttt.cloudsack.biz/ping", masq, nil, passthrough)

	certs := x509.NewCertPool()
	certs.AddCert(cloudSack.Certificate())
	Configure(certs, map[string]*Provider{"cloudsack": p}, "cloudsack", "")

	rt, ok := NewDirect(10 * time.Second)
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
		p := NewProvider(alias, "https://ttt.sadcloud.io/ping", masq, validator, nil)

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
	masqueradesExhausted := fmt.Sprintf(`Get "%v": could not dial any masquerade?`, testURL)

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
		direct, ok := NewDirect(1 * time.Second)
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
			if !assert.Nil(t, err) {
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
	data, err := ioutil.ReadFile(cacheFile)
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
	ioutil.WriteFile(cacheFile, messedUp, 0644)
}
