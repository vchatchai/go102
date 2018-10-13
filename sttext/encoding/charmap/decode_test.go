package charmap

import "testing"

func TestDecode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestDecode"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Decode()
		})
	}
}
