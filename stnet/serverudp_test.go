package stnet

import "testing"

func TestServerUDP(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"TestServerUDP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ServerUDP()
		})
	}
}
