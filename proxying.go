package fronted

import (
	"net"
	"net/http"
	"time"

	"github.com/getlantern/uuid"
)

// DialTimeout dials out using masquerading
func DialTimeout(timeout time.Duration) (net.Conn, error) {
	_d, ok := _instance.Get(timeout)
	if !ok {
		return nil, dialError("Timeout waiting for masquerade")
	}
	d := _d.(*direct)
	conn, masqueradeGood, err := d.dial()
	return &proxyingConn{Conn: conn, masqueradeGood: masqueradeGood}, err
}

// PrepareForProxyingVia prepares a request for domain-fronted proxying via the
// specified proxyHost.
func PrepareForProxyingVia(proxyHost string, req *http.Request) {
	// Store original URL for domain-fronting
	req.Header.Set("X-Ddf-Url", req.URL.String())
	// Set a unique request-id just to make sure we bust the cache
	req.Header.Set("X-Ddf-Request-Id", uuid.NewRandom().String())
	req.URL.Scheme = "http"
	req.URL.Host = proxyHost
	req.URL.Path = ""
	req.URL.RawPath = ""
	req.URL.RawQuery = ""
	req.Host = proxyHost
}

// AfterProxying records the result of proxying
func AfterProxying(conn net.Conn, req *http.Request, resp *http.Response, err error) {
	conn.(*proxyingConn).masqueradeGood(err == nil && resp != nil && resp.Header.Get("X-Cache") == "Error from cloudfront")
}

type dialError string

func (e dialError) Error() string {
	return string(e)
}

func (e dialError) Timeout() bool {
	return true
}

func (e dialError) Temporary() bool {
	return true
}

type proxyingConn struct {
	net.Conn
	masqueradeGood func(bool) bool
}
