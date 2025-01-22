package fronted

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/getlantern/ops"
)

type connectedRoundTripper struct {
	front Front
	net.Conn
	provider *Provider
}

func newConnectedRoundTripper(fr Front, conn net.Conn, provider *Provider) connectedRoundTripper {
	return connectedRoundTripper{
		front:    fr,
		Conn:     conn,
		provider: provider,
	}
}

// Also implements http.RoundTripper
func (crt connectedRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	op := ops.Begin("fronted_request")
	defer op.End()
	originHost := req.URL.Hostname()
	frontedHost := crt.provider.Lookup(originHost)
	if frontedHost == "" {
		// this error is not the masquerade's fault in particular
		// so it is returned as good.
		crt.Conn.Close()
		crt.front.markWithResult(true)
		err := fmt.Errorf("no domain fronting mapping for '%s'. Please add it to provider_map.yaml or equivalent for %s",
			crt.front.getProviderID(), originHost)
		op.FailIf(err)
		return nil, err
	}
	log.Debugf("Translated origin %s -> %s for provider %s...", originHost, frontedHost, crt.front.getProviderID())

	reqi, err := withDomainFront(req, frontedHost, req.Body)
	if err != nil {
		return nil, op.FailIf(log.Errorf("Failed to copy http request with origin translated to %v?: %v", frontedHost, err))
	}
	disableKeepAlives := true
	if strings.EqualFold(reqi.Header.Get("Connection"), "upgrade") {
		disableKeepAlives = false
	}

	tr := connectedConnHTTPTransport(crt.Conn, disableKeepAlives)
	resp, err := tr.RoundTrip(reqi)
	if err != nil {
		log.Debugf("Could not complete request: %v", err)
		crt.front.markWithResult(false)
		return nil, err
	}

	err = crt.provider.ValidateResponse(resp)
	if err != nil {
		log.Debugf("Could not complete request: %v", err)
		resp.Body.Close()
		crt.front.markWithResult(false)
		return nil, err
	}

	crt.front.markWithResult(true)
	return resp, nil
}

// connectedConnHTTPTransport uses a preconnected connection to the CDN to make HTTP requests.
// This uses the pre-established connection to the CDN on the fronting domain.
func connectedConnHTTPTransport(conn net.Conn, disableKeepAlives bool) http.RoundTripper {
	return &connectedTransport{
		Transport: http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				return conn, nil
			},
			TLSHandshakeTimeout: 40 * time.Second,
			DisableKeepAlives:   disableKeepAlives,
			IdleConnTimeout:     70 * time.Second,
		},
	}
}

// connectedTransport is a wrapper struct enabling us to modify the protocol of outgoing
// requests to make them all HTTP instead of potentially HTTPS, which breaks our particular
// implemenation of direct domain fronting.
type connectedTransport struct {
	http.Transport
}

func (ct *connectedTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	defer func(op ops.Op) { op.End() }(ops.Begin("direct_transport_roundtrip"))

	// The connection is already encrypted by domain fronting.  We need to rewrite URLs starting
	// with "https://" to "http://", lest we get an error for doubling up on TLS.

	// The RoundTrip interface requires that we not modify the memory in the request, so we just
	// create a copy.
	norm := new(http.Request)
	*norm = *req // includes shallow copies of maps, but okay
	norm.URL = new(url.URL)
	*norm.URL = *req.URL
	norm.URL.Scheme = "http"
	return ct.Transport.RoundTrip(norm)
}

func withDomainFront(req *http.Request, frontedHost string, body io.ReadCloser) (*http.Request, error) {
	urlCopy := *req.URL
	urlCopy.Host = frontedHost
	r, err := http.NewRequestWithContext(req.Context(), req.Method, urlCopy.String(), body)
	if err != nil {
		return nil, err
	}

	for k, vs := range req.Header {
		if !strings.EqualFold(k, "Host") {
			v := make([]string, len(vs))
			copy(v, vs)
			r.Header[k] = v
		}
	}
	return r, nil
}
