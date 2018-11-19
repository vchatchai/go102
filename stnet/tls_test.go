package stnet

import "testing"

func TestTLS(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestTLS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TLS()
		})
	}
}
