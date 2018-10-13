package charmap

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestEncode"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Encode()
		})
	}
}
