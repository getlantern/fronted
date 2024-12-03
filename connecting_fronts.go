package fronted

import (
	"context"
)

type workingFronts interface {
	onConnected(m Front)
	connectingFront(context.Context) (Front, error)
	size() int
}

type connectingFronts struct {
	// Create a channel of fronts that are connecting.
	frontsCh chan Front
}

// Make sure that connectingFronts is a connectListener
var _ workingFronts = &connectingFronts{}

// newConnectingFronts creates a new ConnectingFronts struct with an empty slice of Masquerade IPs and domains.
func newConnectingFronts(size int) *connectingFronts {
	return &connectingFronts{
		// We overallocate the channel to avoid blocking.
		frontsCh: make(chan Front, size),
	}
}

// AddFront adds a new front to the list of fronts.
func (cf *connectingFronts) onConnected(m Front) {
	cf.frontsCh <- m
}

func (cf *connectingFronts) connectingFront(ctx context.Context) (Front, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case m := <-cf.frontsCh:
			// The front may have stopped succeeding since we last checked,
			// so only return it if it's still succeeding.
			if m.isSucceeding() {
				// Add the front back to the channel.
				cf.frontsCh <- m
				return m, nil
			}
		}
	}
}

func (cf *connectingFronts) size() int {
	return len(cf.frontsCh)
}
