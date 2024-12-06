package fronted

import (
	"context"
)

type connectingFronts interface {
	onConnected(m Front)
	connectingFront(context.Context) (Front, error)
	size() int
}

type connecting struct {
	// Create a channel of fronts that are connecting.
	frontsCh chan Front
}

// Make sure that connectingFronts is a connectListener
var _ connectingFronts = &connecting{}

// newConnectingFronts creates a new ConnectingFronts struct with a channel of fronts that have
// successfully connected.
func newConnectingFronts(size int) *connecting {
	return &connecting{
		// We overallocate the channel to avoid blocking.
		frontsCh: make(chan Front, size),
	}
}

// onConnected adds a working front to the channel of working fronts.
func (cf *connecting) onConnected(m Front) {
	cf.frontsCh <- m
}

func (cf *connecting) connectingFront(ctx context.Context) (Front, error) {
	for {
		select {
		// This is typically the context of the HTTP request. If the context is done, return an error.
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

func (cf *connecting) size() int {
	return len(cf.frontsCh)
}
