package domainfronted

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/getlantern/connpool"
	"github.com/getlantern/enproxy"
	"github.com/getlantern/golog"

	"gopkg.in/getlantern/tlsdialer.v2"
)

const (
	CONNECT = "CONNECT" // HTTP CONNECT method
)

var (
	log = golog.LoggerFor("domainfronted")

	// Cutoff for logging warnings about a dial having taken a long time.
	longDialLimit = 10 * time.Second

	// idleTimeout needs to be small enough that we stop using connections
	// before the upstream server/CDN closes them itself.
	// TODO: make this configurable.
	idleTimeout = 10 * time.Second
)

// ClientConfig captures the configuration of a domain-fronted client.
type ClientConfig struct {
	// Host: the host (e.g. getiantem.org)
	Host string

	// Port: the port (e.g. 443)
	Port int

	// Masquerades: the Masquerades to use when domain-fronting. These will be
	// verified when the client starts.
	Masquerades []*Masquerade

	// InsecureSkipVerify: if true, server's certificate is not verified.
	InsecureSkipVerify bool

	// RootCAs: optional CertPool specifying the root CAs to use for verifying
	// servers
	RootCAs *x509.CertPool

	// BufferRequests: if true, requests to the proxy will be buffered and sent
	// with identity encoding.  If false, they'll be streamed with chunked
	// encoding.
	BufferRequests bool

	// DialTimeoutMillis: how long to wait on dialing server before timing out
	// (defaults to 5 seconds)
	DialTimeoutMillis int

	// RedialAttempts: number of times to try redialing. The total number of
	// dial attempts will be 1 + RedialAttempts.
	RedialAttempts int

	// Weight: relative weight versus other servers (for round-robin)
	Weight int

	// QOS: relative quality of service offered. Should be >= 0, with higher
	// values indicating higher QOS.
	QOS int

	// OnDialStats is an optional callback that will get called on every dial to
	// the server to report stats on what was dialed and how long each step
	// took.
	OnDialStats func(domain, addr string, resolutionTime, connectTime, handshakeTime time.Duration)
}

// Client provides a mechanism for dialing domain-fronted servers.
type Client struct {
	cfg             *ClientConfig
	masquerades     *verifiedMasqueradeSet
	connPool        *connpool.Pool
	enproxyConfig   *enproxy.Config
	tlsConfigs      map[string]*tls.Config
	tlsConfigsMutex sync.Mutex
}

// NewClient creates a new client for the given ClientConfig.
func NewClient(cfg *ClientConfig) *Client {
	client := &Client{
		cfg:        cfg,
		tlsConfigs: make(map[string]*tls.Config),
	}
	if client.cfg.Masquerades != nil {
		client.masquerades = client.verifiedMasquerades()
	}
	client.connPool = &connpool.Pool{
		MinSize:      30,
		ClaimTimeout: idleTimeout,
		Dial:         client.dialServer,
	}
	client.enproxyConfig = client.enproxyConfigWith(func(addr string) (net.Conn, error) {
		return client.connPool.Get()
	})
	return client
}

// Dial dials upstream using domain-fronting.
func (client *Client) Dial(network, addr string) (net.Conn, error) {
	if !strings.Contains(network, "tcp") {
		return nil, fmt.Errorf("Protocol %s is not supported, only tcp is supported", network)
	}

	conn := &enproxy.Conn{
		Addr:   addr,
		Config: client.enproxyConfig,
	}
	err := conn.Connect()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// Close closes the cilent, in particular closing the underlying connection
// pool.
func (client *Client) Close() {
	if client.connPool != nil {
		// We stop the connPool on a goroutine so as not to wait for Stop to finish
		go client.connPool.Stop()
	}
}

// HttpClientUsing creates a simple domain-fronted HTTP client using the
// specified Masquerade.
func (client *Client) HttpClientUsing(masquerade *Masquerade) *http.Client {
	enproxyConfig := client.enproxyConfigWith(func(addr string) (net.Conn, error) {
		return client.dialServerWith(masquerade)
	})

	return &http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				conn := &enproxy.Conn{
					Addr:   addr,
					Config: enproxyConfig,
				}
				err := conn.Connect()
				if err != nil {
					return nil, err
				}
				return conn, nil
			},
		},
	}
}

func (client *Client) enproxyConfigWith(dialProxy func(addr string) (net.Conn, error)) *enproxy.Config {
	return &enproxy.Config{
		DialProxy: dialProxy,
		NewRequest: func(upstreamHost string, method string, body io.Reader) (req *http.Request, err error) {
			if upstreamHost == "" {
				// No specific host requested, use configured one
				upstreamHost = client.cfg.Host
			}
			return http.NewRequest(method, "http://"+upstreamHost+"/", body)
		},
		BufferRequests: client.cfg.BufferRequests,
		IdleTimeout:    idleTimeout, // TODO: make this configurable
	}
}

func (client *Client) dialServer() (net.Conn, error) {
	var masquerade *Masquerade
	if client.masquerades != nil {
		masquerade = client.masquerades.nextVerified()
	}
	return client.dialServerWith(masquerade)
}

func (client *Client) dialServerWith(masquerade *Masquerade) (net.Conn, error) {
	dialTimeout := time.Duration(client.cfg.DialTimeoutMillis) * time.Millisecond
	if dialTimeout == 0 {
		dialTimeout = 20 * time.Second
	}

	// Note - we need to suppress the sending of the ServerName in the client
	// handshake to make host-spoofing work with Fastly.  If the client Hello
	// includes a server name, Fastly checks to make sure that this matches the
	// Host header in the HTTP request and if they don't match, it returns
	// a 400 Bad Request error.
	sendServerNameExtension := false

	cwt, err := tlsdialer.DialForTimings(
		&net.Dialer{
			Timeout: dialTimeout,
		},
		"tcp",
		client.addressForServer(masquerade),
		sendServerNameExtension,
		client.tlsConfig(masquerade))

	if client.cfg.OnDialStats != nil {
		domain := ""
		if masquerade != nil {
			domain = masquerade.Domain
		}

		resultAddr := ""
		if err == nil {
			resultAddr = cwt.Conn.RemoteAddr().String()
		}

		client.cfg.OnDialStats(domain, resultAddr, cwt.ResolutionTime, cwt.ConnectTime, cwt.HandshakeTime)
	}

	if err != nil && masquerade != nil {
		err = fmt.Errorf("Unable to dial masquerade %s: %s", masquerade.Domain, err)
	}
	return cwt.Conn, err
}

// Get the address to dial for reaching the server
func (client *Client) addressForServer(masquerade *Masquerade) string {
	return fmt.Sprintf("%s:%d", client.serverHost(masquerade), client.cfg.Port)
}

func (client *Client) serverHost(masquerade *Masquerade) string {
	serverHost := client.cfg.Host
	if masquerade != nil {
		if masquerade.IpAddress != "" {
			serverHost = masquerade.IpAddress
		} else if masquerade.Domain != "" {
			serverHost = masquerade.Domain
		}
	}
	return serverHost
}

// tlsConfig builds a tls.Config for dialing the upstream host. Constructed
// tls.Configs are cached on a per-masquerade basis to enable client session
// caching and reduce the amount of PEM certificate parsing.
func (client *Client) tlsConfig(masquerade *Masquerade) *tls.Config {
	client.tlsConfigsMutex.Lock()
	defer client.tlsConfigsMutex.Unlock()

	if client.tlsConfigs == nil {
		client.tlsConfigs = make(map[string]*tls.Config)
	}

	serverName := client.cfg.Host
	if masquerade != nil {
		serverName = masquerade.Domain
	}
	tlsConfig := client.tlsConfigs[serverName]
	if tlsConfig == nil {
		tlsConfig = &tls.Config{
			ClientSessionCache: tls.NewLRUClientSessionCache(1000),
			InsecureSkipVerify: client.cfg.InsecureSkipVerify,
			ServerName:         serverName,
			RootCAs:            client.cfg.RootCAs,
		}
		client.tlsConfigs[serverName] = tlsConfig
	}

	return tlsConfig
}
