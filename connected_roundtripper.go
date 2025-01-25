package fronted

import (
	"fmt"
	"net"
	"net/http"
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
	ctx, span := tracer.Start(req.Context(), "frontRoundTrip")
	defer span.End()
	req = req.WithContext(ctx)
	originHost := req.URL.Hostname()
	frontedHost := crt.provider.Lookup(originHost)
	if frontedHost == "" {
		// this error is not the masquerade's fault in particular
		// so it is returned as good.
		crt.Conn.Close()
		crt.front.markWithResult(true)
		err := fmt.Errorf("no domain fronting mapping for '%s'. Please add it to provider_map.yaml or equivalent", originHost)
		op.FailIf(err)
		return nil, err
	}
	log.Debugf("Translated origin %s -> %s.", originHost, frontedHost)

	reqi := withDomainFront(req, frontedHost)
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
	log.Debug("Request completed successfully")
	return resp, nil
}

// connectedConnHTTPTransport uses a preconnected connection to the CDN to make HTTP requests.
// This uses the pre-established connection to the CDN on the fronting domain.
func connectedConnHTTPTransport(conn net.Conn, disableKeepAlives bool) *http.Transport {
	return &http.Transport{
		Dial: func(network, addr string) (net.Conn, error) {
			return conn, nil
		},
		TLSHandshakeTimeout: 20 * time.Second,
		DisableKeepAlives:   disableKeepAlives,
		IdleConnTimeout:     70 * time.Second,
	}
}

func withDomainFront(req *http.Request, frontedHost string) *http.Request {
	// The connection is already encrypted by domain fronting.  We need to rewrite URLs starting
	// with "https://" to "http://", lest we get an error for doubling up on TLS.

	// The RoundTrip interface requires that we not modify the memory in the request, so we just
	// create a copy.
	newReq := req.Clone(req.Context())
	newReq.URL.Scheme = "http"
	newReq.URL.Host = frontedHost

	return newReq
}
