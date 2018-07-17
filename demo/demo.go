package main

import (
	"bufio"
	"bytes"
	"crypto/x509"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/getlantern/fronted"
	"github.com/getlantern/go-update/check"
	"github.com/getlantern/keyman"
	"github.com/getlantern/withtimeout"
	"github.com/getlantern/yaml"
)

const (
	vetTimeout = 2 * time.Minute
)

func terminate(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
	fmt.Println("\n\n******************** HIT RETURN TO EXIT **************************")
	bufio.NewReader(os.Stdin).ReadLine()
	os.Exit(0)
}

func main() {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		terminate("Unable to set up tmpDir: %v", err)
	}

	// Read embedded configuration
	cfg := &Config{}
	err = yaml.Unmarshal([]byte(configYaml), cfg)
	if err != nil {
		terminate("Unable to read config: %v", err)
	}

	// Initialize certificate pool with all trusted CAs
	trustedCAs := x509.NewCertPool()
	for _, pem := range cfg.TrustedCAs {
		cert, err := keyman.LoadCertificateFromPEMBytes([]byte(pem))
		if err != nil {
			terminate("Error parsing certificate: %v", err)
		}
		trustedCAs.AddCert(cert.X509())
	}

	var results sortedRequestResults

	test := func(name string, rt http.RoundTripper) {
		for _, req := range testRequests() {
			_resp, _, err := withtimeout.Do(30*time.Second, func() (interface{}, error) {
				return rt.RoundTrip(req)
			})
			var statusCode int
			var bodySize int64
			if err == nil {
				resp := _resp.(*http.Response)
				statusCode = resp.StatusCode
				if resp.Body != nil {
					bodySize, _ = io.Copy(ioutil.Discard, resp.Body)
					resp.Body.Close()
				}
			}
			results = append(results, &requestResult{name, req, statusCode, err, uint64(bodySize)})
		}
	}

	// Test requests with standard http first
	test("_baseline", &http.Transport{})

	// Test all requests with each provider
	for name, provider := range cfg.Providers {
		provider.Validator = fronted.NewStatusCodeValidator(provider.RejectStatus)
		fronted.Configure(trustedCAs, map[string]*fronted.Provider{name: &provider.Provider}, "cloudfront", filepath.Join(tmpDir, "cache"))
		direct, ok := fronted.NewDirect(vetTimeout)
		if ok {
			test(name, direct)
		}
	}

	sort.Sort(results)
	fmt.Println("--------- RESULTS ---------")
	for _, rr := range results {
		fmt.Println(rr)
	}

	csvFile, err := os.Create(filepath.Join(tmpDir, "results.csv"))
	if err != nil {
		terminate("Unable to create CSV file: %v", err)
	}
	defer csvFile.Close()

	w := csv.NewWriter(csvFile)
	w.Write([]string{"url", "provider", "status code", "error", "body size"})
	for _, rr := range results {
		w.Write(rr.strings())
	}
	w.Flush()

	terminate("\n\n\nCSV file with results at %v\n\n", csvFile.Name())
}

func testRequests() []*http.Request {
	var requests []*http.Request
	addRequest := func(method string, url string, body io.Reader) *http.Request {
		req, _ := http.NewRequest(method, url, body)
		requests = append(requests, req)
		return req
	}

	addRequest(http.MethodGet, "https://api.getiantem.org/plans", nil)
	addRequest(http.MethodGet, "https://api-staging.getiantem.org/plans", nil)
	addRequest(http.MethodPost, "https://borda.lantern.io/measurements", nil)
	addRequest(http.MethodGet, "http://config.getiantem.org/proxies.yaml.gz", nil)
	addRequest(http.MethodGet, "http://geo.getiantem.org/lookup", nil)
	addRequest(http.MethodPost, "https://update.getlantern.org/update", makeAutoUpdateRequestBody()).Header.Set("Content-Type", "application/json")
	return requests
}

func makeRequest(method string, url string, body io.ReadCloser) *http.Request {
	req, _ := http.NewRequest(method, url, body)
	return req
}

func makeAutoUpdateRequestBody() io.Reader {
	p := &check.Params{
		Channel: "stable",
		OS:      "windows",
		Arch:    "amd64",
		Version: 1,
		Tags: map[string]string{
			"channel": "stable",
			"os":      "windows",
			"arch":    "amd64",
		},
	}

	body, _ := json.Marshal(p)
	return bytes.NewReader(body)
}

type requestResult struct {
	provider   string
	req        *http.Request
	statusCode int
	err        error
	bodySize   uint64
}

func (rr *requestResult) String() string {
	return fmt.Sprintf("RESULT %-40v   %-10v   %v   %v   %v", rr.req.URL.String(), rr.provider, rr.statusCode, errOrEmpty(rr.err), humanize.Bytes(rr.bodySize))
}

func (rr *requestResult) strings() []string {
	return []string{rr.req.URL.String(), rr.provider, fmt.Sprint(rr.statusCode), errOrEmpty(rr.err), humanize.Bytes(rr.bodySize)}
}

func errOrEmpty(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

type sortedRequestResults []*requestResult

func (rr sortedRequestResults) Len() int { return len(rr) }

func (rr sortedRequestResults) Swap(i, j int) {
	rr[i], rr[j] = rr[j], rr[i]
}

func (rr sortedRequestResults) Less(i, j int) bool {
	a, b := rr[i], rr[j]
	urlComparison := strings.Compare(a.req.URL.String(), b.req.URL.String())
	if urlComparison == 0 {
		return strings.Compare(a.provider, b.provider) < 0
	}
	return urlComparison < 0
}

type Config struct {
	Providers  map[string]*ProviderConfig
	TrustedCAs []string
}

type ProviderConfig struct {
	fronted.Provider
	RejectStatus []int
}

const configYaml = `
providers:
  cloudfront:
    rejectstatus:                                            [403]
    provider:
      hostaliases:
        api.getiantem.org:                                         d2n32kma9hyo9f.cloudfront.net
        api-staging.getiantem.org:                                 d16igwq64x5e11.cloudfront.net
        borda.lantern.io:                                          d157vud77ygy87.cloudfront.net
        config.getiantem.org:                                      d2wi0vwulmtn99.cloudfront.net
        config-staging.getiantem.org:                              d33pfmbpauhmvd.cloudfront.net
        geo.getiantem.org:                                         d3u5fqukq7qrhd.cloudfront.net
        globalconfig.flashlightproxy.com:                          d24ykmup0867cj.cloudfront.net
        update.getlantern.org:                                     d2yl1zps97e5mx.cloudfront.net
        github.com:                                                d2yl1zps97e5mx.cloudfront.net
        github-production-release-asset-2e65be.s3.amazonaws.com:   d37kom4pw4aa7b.cloudfront.net
      testurl:
        http://d157vud77ygy87.cloudfront.net/ping
      masquerades:
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.197
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.184
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.191
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.193
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.199
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.190
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.204
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.194
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.195
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.150
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.146
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.170
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.154
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.206
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.207
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.168
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.175
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.171
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.196
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.176
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.163
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.153
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.203
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.145
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.179
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.177
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.158
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.192
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.165
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.172
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.147
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.200
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.152
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.149
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.202
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.164
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.162
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.182
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.180
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.205
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.173
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.186
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.185
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.188
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.167
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.151
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.166
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.183
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.181
      - domain:                                                  cloudfront.net
        ipaddress:                                               54.230.11.159
  akamai:
    rejectstatus:                                            [403]
    provider:
      hostaliases:
        api-staging.getiantem.org:                               api-staging.akamai.getiantem.org
        api.getiantem.org:                                       api.akamai.getiantem.org
        borda.lantern.io:                                        borda.akamai.getiantem.org
        config-staging.getiantem.org:                            config-staging.akamai.getiantem.org
        config.getiantem.org:                                    config.akamai.getiantem.org
        geo.getiantem.org:                                       geo.akamai.getiantem.org
        github-production-release-asset-2e65be.s3.amazonaws.com: github-release-asset.akamai.getiantem.org
        github.com:                                              github.akamai.getiantem.org
        globalconfig.flashlightproxy.com:                        globalconfig.akamai.getiantem.org
        update.getlantern.org:                                   update.akamai.getiantem.org
      testurl:                                                   https://borda.akamai.getiantem.org/ping
      masquerades:
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               96.17.109.101
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               96.17.109.85
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               96.17.109.83
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.37
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.102
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               96.17.109.72
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.112
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.156
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.4
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.127
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.18
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.115
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.140
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.144
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.41
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.43
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.138
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.111
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.157
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.168
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.121
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               96.17.109.49
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.122
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.25
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               63.239.233.93
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.113
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.35
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.11
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.105
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               63.239.233.76
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.107
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.98
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.48
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               63.239.233.90
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.36
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.53
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               96.17.109.11
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               63.239.233.78
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               96.17.109.84
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.57
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.34
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.92
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.52
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               209.116.151.125
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               2.16.4.66
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               2.16.4.115
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               2.16.4.26
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               2.16.4.144
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               2.16.4.57
      - domain:                                                  a248.e.akamai.net
        ipaddress:                                               184.50.27.102
  edgecast:
    rejectstatus:                                            [403]
    provider:
      hostaliases:
        api-staging.getiantem.org:                               api-staging.edgecast.getiantem.org
        api.getiantem.org:                                       api.edgecast.getiantem.org
        borda.lantern.io:                                        borda.edgecast.getiantem.org
        config-staging.getiantem.org:                            config-staging.edgecast.getiantem.org
        config.getiantem.org:                                    config.edgecast.getiantem.org
        geo.getiantem.org:                                       geo.edgecast.getiantem.org
        github-production-release-asset-2e65be.s3.amazonaws.com: github-release-asset.edgecast.getiantem.org
        github.com:                                              github.edgecast.getiantem.org
        globalconfig.flashlightproxy.com:                        globalconfig.edgecast.getiantem.org
        update.getlantern.org:                                   update.edgecast.getiantem.org
      testurl:                                                   https://borda.edgecast.getiantem.org/ping
      masquerades:
      - domain:                                                  ams.allbet999.net
        ipaddress:                                               72.21.92.61
      - domain:                                                  app.visanow.com
        ipaddress:                                               72.21.92.217
      - domain:                                                  bb17.edgecastcdn.net
        ipaddress:                                               72.21.92.244
      - domain:                                                  bighistoryproject.com
        ipaddress:                                               72.21.92.204
      - domain:                                                  cargurus.com
        ipaddress:                                               72.21.92.79
      - domain:                                                  crucial.com
        ipaddress:                                               72.21.92.124
      - domain:                                                  dailymotion.com
        ipaddress:                                               72.21.92.102
      - domain:                                                  ewtn.com
        ipaddress:                                               72.21.92.228
      - domain:                                                  ezalber.com
        ipaddress:                                               72.21.92.104
      - domain:                                                  geccdn.net
        ipaddress:                                               72.21.92.9
      - domain:                                                  gh.wego.com
        ipaddress:                                               72.21.92.77
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.120
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.164
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.117
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.183
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.121
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.67
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.84
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.219
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.24
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.122
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.93
      - domain:                                                  gp1.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.125
      - domain:                                                  ifastps.com.cn
        ipaddress:                                               72.21.92.206
      - domain:                                                  mx-abgame88.garcade.net
        ipaddress:                                               72.21.92.209
      - domain:                                                  mylife.com
        ipaddress:                                               72.21.92.110
      - domain:                                                  mythingsmedia.net
        ipaddress:                                               72.21.92.224
      - domain:                                                  nossl.edgecastcdn.net
        ipaddress:                                               72.21.92.138
      - domain:                                                  nossl.edgecastcdn.net
        ipaddress:                                               72.21.92.22
      - domain:                                                  nossl.edgecastcdn.net
        ipaddress:                                               72.21.92.81
      - domain:                                                  nossl.edgecastcdn.net
        ipaddress:                                               72.21.92.143
      - domain:                                                  nossl.edgecastcdn.net
        ipaddress:                                               72.21.92.47
      - domain:                                                  nossl.edgecastcdn.net
        ipaddress:                                               72.21.92.141
      - domain:                                                  s5.adn.edgecastcdn.net
        ipaddress:                                               72.21.92.251
      - domain:                                                  smartertravel.com
        ipaddress:                                               72.21.92.230
      - domain:                                                  speedtest.net
        ipaddress:                                               72.21.92.82
      - domain:                                                  stage.web3.justwink.com
        ipaddress:                                               72.21.92.159
      - domain:                                                  starlimscloud.com
        ipaddress:                                               72.21.92.171
      - domain:                                                  surveys.com
        ipaddress:                                               72.21.92.133
      - domain:                                                  target.com
        ipaddress:                                               72.21.92.128
      - domain:                                                  trivago.com
        ipaddress:                                               72.21.92.145
      - domain:                                                  understood.org
        ipaddress:                                               72.21.92.157
      - domain:                                                  uprinting.com
        ipaddress:                                               72.21.92.95
      - domain:                                                  www.dog.com
        ipaddress:                                               72.21.92.112
      - domain:                                                  www.oneservice.sg
        ipaddress:                                               72.21.92.78
      - domain:                                                  www.plus500.com
        ipaddress:                                               72.21.92.83
      - domain:                                                  www.regus.fr
        ipaddress:                                               72.21.92.215
      - domain:                                                  www.ura.gov.sg
        ipaddress:                                               72.21.92.193
      - domain:                                                  www.ziprealty.com
        ipaddress:                                               72.21.92.26
      - domain:                                                  xda-developers.com
        ipaddress:                                               72.21.92.105
trustedcas:
  - "-----BEGIN CERTIFICATE-----\nMIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD\nQTAeFw0wNjExMTAwMDAwMDBaFw0zMTExMTAwMDAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IENBMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4jvhEXLeqKTTo1eqUKKPC3eQyaKl7hLOllsB\nCSDMAZOnTjC3U/dDxGkAV53ijSLdhwZAAIEJzs4bg7/fzTtxRuLWZscFs3YnFo97\nnh6Vfe63SKMI2tavegw5BmV/Sl0fvBf4q77uKNd0f3p4mVmFaG5cIzJLv07A6Fpt\n43C/dxC//AH2hdmoRBBYMql1GNXRor5H4idq9Joz+EkIYIvUX7Q6hL+hqkpMfT7P\nT19sdl6gSzeRntwi5m3OFBqOasv+zbMUZBfHWymeMr/y7vrTC0LUq7dBMtoM1O/4\ngdW7jVg/tRvoSSiicNoxBN33shbyTApOB6jtSj1etX+jkMOvJwIDAQABo2MwYTAO\nBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUA95QNVbR\nTLtm8KPiGxvDl7I90VUwHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUw\nDQYJKoZIhvcNAQEFBQADggEBAMucN6pIExIK+t1EnE9SsPTfrgT1eXkIoyQY/Esr\nhMAtudXH/vTBH1jLuG2cenTnmCmrEbXjcKChzUyImZOMkXDiqw8cvpOp/2PV5Adg\n06O/nVsJ8dWO41P0jmP6P6fbtGbfYmbW0W5BjfIttep3Sp+dWOIrWcBAI+0tKIJF\nPnlUkiaY4IBIqDfv8NZ5YBberOgOzW6sRBc4L0na4UU+Krk2U886UAb3LujEV0ls\nYSEY1QSteDwsOoBrp+uvFRTp2InBuThs4pFsiv9kuXclVzDAGySj4dzp30d8tbQk\nCAUw7C29C79Fv1C5qfPrmAESrciIxpg0X40KPMbp1ZWVbd4=\n-----END CERTIFICATE-----\n"
  - "-----BEGIN CERTIFICATE-----\nMIIDjjCCAnagAwIBAgIQAzrx5qcRqaC7KGSxHQn65TANBgkqhkiG9w0BAQsFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBH\nMjAeFw0xMzA4MDExMjAwMDBaFw0zODAxMTUxMjAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IEcyMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuzfNNNx7a8myaJCtSnX/RrohCgiN9RlUyfuI\n2/Ou8jqJkTx65qsGGmvPrC3oXgkkRLpimn7Wo6h+4FR1IAWsULecYxpsMNzaHxmx\n1x7e/dfgy5SDN67sH0NO3Xss0r0upS/kqbitOtSZpLYl6ZtrAGCSYP9PIUkY92eQ\nq2EGnI/yuum06ZIya7XzV+hdG82MHauVBJVJ8zUtluNJbd134/tJS7SsVQepj5Wz\ntCO7TG1F8PapspUwtP1MVYwnSlcUfIKdzXOS0xZKBgyMUNGPHgm+F6HmIcr9g+UQ\nvIOlCsRnKPZzFBQ9RnbDhxSJITRNrw9FDKZJobq7nMWxM4MphQIDAQABo0IwQDAP\nBgNVHRMBAf8EBTADAQH/MA4GA1UdDwEB/wQEAwIBhjAdBgNVHQ4EFgQUTiJUIBiV\n5uNu5g/6+rkS7QYXjzkwDQYJKoZIhvcNAQELBQADggEBAGBnKJRvDkhj6zHd6mcY\n1Yl9PMWLSn/pvtsrF9+wX3N3KjITOYFnQoQj8kVnNeyIv/iPsGEMNKSuIEyExtv4\nNeF22d+mQrvHRAiGfzZ0JFrabA0UWTW98kndth/Jsw1HKj2ZL7tcu7XUIOGZX1NG\nFdtom/DzMNU+MeKNhJ7jitralj41E6Vf8PlwUHBHQRFXGU7Aj64GxJUTFy8bJZ91\n8rGOmaFvE7FBcf6IKshPECBV1/MUReXgRPTqh5Uykw7+U0b6LJ3/iyK5S9kJRaTe\npLiaWN0bfVKfjllDiIGknibVb63dDcY3fe0Dkhvld1927jyNxF1WW6LZZm6zNTfl\nMrY=\n-----END CERTIFICATE-----\n"
  - "-----BEGIN CERTIFICATE-----\nMIIDrzCCApegAwIBAgIQCDvgVpBCRrGhdWrJWZHHSjANBgkqhkiG9w0BAQUFADBh\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSAwHgYDVQQDExdEaWdpQ2VydCBHbG9iYWwgUm9vdCBD\nQTAeFw0wNjExMTAwMDAwMDBaFw0zMTExMTAwMDAwMDBaMGExCzAJBgNVBAYTAlVT\nMRUwEwYDVQQKEwxEaWdpQ2VydCBJbmMxGTAXBgNVBAsTEHd3dy5kaWdpY2VydC5j\nb20xIDAeBgNVBAMTF0RpZ2lDZXJ0IEdsb2JhbCBSb290IENBMIIBIjANBgkqhkiG\n9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4jvhEXLeqKTTo1eqUKKPC3eQyaKl7hLOllsB\nCSDMAZOnTjC3U/dDxGkAV53ijSLdhwZAAIEJzs4bg7/fzTtxRuLWZscFs3YnFo97\nnh6Vfe63SKMI2tavegw5BmV/Sl0fvBf4q77uKNd0f3p4mVmFaG5cIzJLv07A6Fpt\n43C/dxC//AH2hdmoRBBYMql1GNXRor5H4idq9Joz+EkIYIvUX7Q6hL+hqkpMfT7P\nT19sdl6gSzeRntwi5m3OFBqOasv+zbMUZBfHWymeMr/y7vrTC0LUq7dBMtoM1O/4\ngdW7jVg/tRvoSSiicNoxBN33shbyTApOB6jtSj1etX+jkMOvJwIDAQABo2MwYTAO\nBgNVHQ8BAf8EBAMCAYYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUA95QNVbR\nTLtm8KPiGxvDl7I90VUwHwYDVR0jBBgwFoAUA95QNVbRTLtm8KPiGxvDl7I90VUw\nDQYJKoZIhvcNAQEFBQADggEBAMucN6pIExIK+t1EnE9SsPTfrgT1eXkIoyQY/Esr\nhMAtudXH/vTBH1jLuG2cenTnmCmrEbXjcKChzUyImZOMkXDiqw8cvpOp/2PV5Adg\n06O/nVsJ8dWO41P0jmP6P6fbtGbfYmbW0W5BjfIttep3Sp+dWOIrWcBAI+0tKIJF\nPnlUkiaY4IBIqDfv8NZ5YBberOgOzW6sRBc4L0na4UU+Krk2U886UAb3LujEV0ls\nYSEY1QSteDwsOoBrp+uvFRTp2InBuThs4pFsiv9kuXclVzDAGySj4dzp30d8tbQk\nCAUw7C29C79Fv1C5qfPrmAESrciIxpg0X40KPMbp1ZWVbd4=\n-----END CERTIFICATE-----\n"
  - "-----BEGIN CERTIFICATE-----\nMIIDxTCCAq2gAwIBAgIQAqxcJmoLQJuPC3nyrkYldzANBgkqhkiG9w0BAQUFADBs\nMQswCQYDVQQGEwJVUzEVMBMGA1UEChMMRGlnaUNlcnQgSW5jMRkwFwYDVQQLExB3\nd3cuZGlnaWNlcnQuY29tMSswKQYDVQQDEyJEaWdpQ2VydCBIaWdoIEFzc3VyYW5j\nZSBFViBSb290IENBMB4XDTA2MTExMDAwMDAwMFoXDTMxMTExMDAwMDAwMFowbDEL\nMAkGA1UEBhMCVVMxFTATBgNVBAoTDERpZ2lDZXJ0IEluYzEZMBcGA1UECxMQd3d3\nLmRpZ2ljZXJ0LmNvbTErMCkGA1UEAxMiRGlnaUNlcnQgSGlnaCBBc3N1cmFuY2Ug\nRVYgUm9vdCBDQTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMbM5XPm\n+9S75S0tMqbf5YE/yc0lSbZxKsPVlDRnogocsF9ppkCxxLeyj9CYpKlBWTrT3JTW\nPNt0OKRKzE0lgvdKpVMSOO7zSW1xkX5jtqumX8OkhPhPYlG++MXs2ziS4wblCJEM\nxChBVfvLWokVfnHoNb9Ncgk9vjo4UFt3MRuNs8ckRZqnrG0AFFoEt7oT61EKmEFB\nIk5lYYeBQVCmeVyJ3hlKV9Uu5l0cUyx+mM0aBhakaHPQNAQTXKFx01p8VdteZOE3\nhzBWBOURtCmAEvF5OYiiAhF8J2a3iLd48soKqDirCmTCv2ZdlYTBoSUeh10aUAsg\nEsxBu24LUTi4S8sCAwEAAaNjMGEwDgYDVR0PAQH/BAQDAgGGMA8GA1UdEwEB/wQF\nMAMBAf8wHQYDVR0OBBYEFLE+w2kD+L9HAdSYJhoIAu9jZCvDMB8GA1UdIwQYMBaA\nFLE+w2kD+L9HAdSYJhoIAu9jZCvDMA0GCSqGSIb3DQEBBQUAA4IBAQAcGgaX3Nec\nnzyIZgYIVyHbIUf4KmeqvxgydkAQV8GK83rZEWWONfqe/EW1ntlMMUu4kehDLI6z\neM7b41N5cdblIZQB2lWHmiRk9opmzN6cN82oNLFpmyPInngiK3BD41VHMWEZ71jF\nhS9OMPagMRYjyOfiZRYzy78aG6A9+MpeizGLYAiJLQwGXFK3xPkKmNEVX58Svnw2\nYzi9RKR/5CYrCsSXaQ3pjOLAEFe4yHYSkVXySGnYvCoCWw9E1CAx2/S6cCZdkGCe\nvEsXCS+0yx5DaMkHJ8HSXPfqIbloEpw8nL+e/IBcm2PN7EeqJSdnoDfzAIJ9VNep\n+OkuE6N36B9K\n-----END CERTIFICATE-----\n"
`
