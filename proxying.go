package fronted

import (
	"net"
	"net/http"
	"time"

	"github.com/getlantern/uuid"
)

// DialTimeout dials out using masquerading, proxying via the given proxyHost
// and optionally notifying onRequest for each proxied request.
func DialTimeout(proxyHost string, timeout time.Duration, onRequest func(*http.Request)) (net.Conn, error) {
	_d, ok := _instance.Get(timeout)
	if !ok {
		return nil, dialError("Timeout waiting for masquerade")
	}
	d := _d.(*direct)
	// TODO: apply timeout to dial here too
	conn, masqueradeGood, err := d.dial()
	if onRequest == nil {
		onRequest = func(*http.Request) {}
	}
	return &proxyingConn{Conn: conn, proxyHost: proxyHost, onRequest: onRequest, masqueradeGood: masqueradeGood}, err
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
	proxyHost      string
	onRequest      func(*http.Request)
	masqueradeGood func(bool) bool
}

// OnRequest implements the proxy.RequestAware interface to prepare the domain-
// fronted request.
func (conn *proxyingConn) OnRequest(req *http.Request) {
	conn.onRequest(req)
	if req.URL.Scheme == "" {
		// HTTPS requests had their scheme stripped, add it back
		req.URL.Scheme = "https"
	}
	// Store original URL for domain-fronting
	req.Header.Set("X-Ddf-Url", req.URL.String())
	// Set a unique request-id just to make sure we bust the cache
	req.Header.Set("X-Ddf-Request-Id", uuid.NewRandom().String())
	req.URL.Scheme = "http"
	req.URL.Host = conn.proxyHost
	req.URL.Path = ""
	req.URL.RawPath = ""
	req.URL.RawQuery = ""
	req.Host = conn.proxyHost
}

// OnResponse implements the proxy.ResponseAware interface
func (conn *proxyingConn) OnResponse(req *http.Request, resp *http.Response, err error) {
	if conn.masqueradeGood == nil {
		return
	}
	conn.masqueradeGood(err == nil && resp != nil && resp.Header.Get("X-Cache") == "Error from cloudfront")
	conn.masqueradeGood = nil // only call masqueradeGood for first response
}

// MITMSkipEncryption is a marker to tell the mitm library not to bother
// encrypting this conn.
func (conn *proxyingConn) MITMSkipEncryption() {}
