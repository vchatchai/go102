package stnet

import "testing"

func TestJsonrpc(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestJsonrpc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Jsonrpc()
		})
	}
}
