package fronted

import (
	"errors"
	"sort"
	"sync"
	"time"
)

type connectTimeFront struct {
	MasqueradeInterface
	connectTime time.Duration
}

type connectingFronts struct {
	fronts []connectTimeFront
	//frontsChan chan MasqueradeInterface
	sync.RWMutex
}

// Make sure that connectingFronts is a connectListener
var _ workingFronts = &connectingFronts{}

// newConnectingFronts creates a new ConnectingFronts struct with an empty slice of Masquerade IPs and domains.
func newConnectingFronts() *connectingFronts {
	return &connectingFronts{
		fronts: make([]connectTimeFront, 0),
		//frontsChan: make(chan MasqueradeInterface),
	}
}

// AddFront adds a new front to the list of fronts.
func (cf *connectingFronts) onConnected(m MasqueradeInterface, connectTime time.Duration) {
	cf.Lock()
	defer cf.Unlock()

	cf.fronts = append(cf.fronts, connectTimeFront{
		MasqueradeInterface: m,
		connectTime:         connectTime,
	})
	// Sort fronts by connect time.
	sort.Slice(cf.fronts, func(i, j int) bool {
		return cf.fronts[i].connectTime < cf.fronts[j].connectTime
	})
	//cf.frontsChan <- m
}

func (cf *connectingFronts) onError(m MasqueradeInterface) {
	cf.Lock()
	defer cf.Unlock()

	// Remove the front from connecting fronts.
	for i, front := range cf.fronts {
		if front.MasqueradeInterface == m {
			cf.fronts = append(cf.fronts[:i], cf.fronts[i+1:]...)
			return
		}
	}
}

func (cf *connectingFronts) workingFront() (MasqueradeInterface, error) {
	cf.RLock()
	defer cf.RUnlock()
	if len(cf.fronts) == 0 {
		return nil, errors.New("no fronts available")
	}
	return cf.fronts[0].MasqueradeInterface, nil
}
