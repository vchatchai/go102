package stnet

import "testing"

func TestLookup(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestLookup"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Lookup()
		})
	}
}
