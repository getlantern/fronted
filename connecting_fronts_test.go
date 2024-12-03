package fronted

import (
	"testing"
)

func TestConnectingFrontsSize(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *connectingFronts
		expected int
	}{
		{
			name: "empty channel",
			setup: func() *connectingFronts {
				return newConnectingFronts(10)
			},
			expected: 0,
		},
		{
			name: "non-empty channel",
			setup: func() *connectingFronts {
				cf := newConnectingFronts(10)
				cf.onConnected(&mockFront{})
				return cf
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cf := tt.setup()
			if got := cf.size(); got != tt.expected {
				t.Errorf("size() = %d, want %d", got, tt.expected)
			}
		})
	}
}
