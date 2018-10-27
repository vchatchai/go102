package stencoding

import "testing"

func TestGob(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestGob"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gob()
		})
	}
}
