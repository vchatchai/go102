package stnet

import "testing"

func TestURL(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test URL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			URL()
		})
	}
}
