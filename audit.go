package fronted

import (
	"net"
	"sync"
)

var (
	dialAuditor   DialAuditor = noopDialAuditor
	dialAuditorMx sync.RWMutex
)

// DialAuditor allows auditing of all dialed connections and optionally replacing them
type DialAuditor func(net.Conn, error) (net.Conn, error)

// SetDialAuditor sets the DialAuditor to use when dialing connections
func SetDialAuditor(auditor DialAuditor) {
	if auditor == nil {
		auditor = noopDialAuditor
	}
	dialAuditorMx.Lock()
	dialAuditor = auditor
	dialAuditorMx.Unlock()
}

func getDialAuditor() DialAuditor {
	dialAuditorMx.RLock()
	auditor := dialAuditor
	dialAuditorMx.RUnlock()
	return auditor
}

func noopDialAuditor(conn net.Conn, err error) (net.Conn, error) {
	return conn, err
}
