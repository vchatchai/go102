package stnet

import "testing"

func TestServerTCP(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestServerTCP"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ServerTCP()
		})
	}
}
