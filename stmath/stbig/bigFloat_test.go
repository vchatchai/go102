package stbig

import "testing"

func TestBigFloat(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestBigFloat"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BigFloat()
		})
	}
}
