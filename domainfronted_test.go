package domainfronted

import (
	"strconv"
	"strings"
	"testing"

	"github.com/getlantern/proxytest"
)

func TestRoundTrip(t *testing.T) {
	server := &Server{
		Addr: "localhost:0",
		AllowNonGlobalDestinations: true,
		CertContext: &CertContext{
			PKFile:         "testpk.pem",
			ServerCertFile: "testcert.pem",
		},
	}
	l, err := server.Listen()
	if err != nil {
		t.Fatalf("Unable to listen: %s", err)
	}
	go func() {
		err = server.Serve(l)
		if err != nil {
			t.Fatalf("Unable to serve: %s", err)
		}
	}()

	addrParts := strings.Split(l.Addr().String(), ":")
	host := addrParts[0]
	port, err := strconv.Atoi(addrParts[1])
	if err != nil {
		t.Fatalf("Unable to parse port: %s", err)
	}
	client := NewClient(&ClientConfig{
		Host:               host,
		Port:               port,
		InsecureSkipVerify: true,
	})

	proxytest.Go(t, client.Dial)
}
