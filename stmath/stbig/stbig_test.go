package stbig

import "testing"

func TestBig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestBig"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Big()
		})
	}
}
