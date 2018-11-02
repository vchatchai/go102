package stfilepath

import "testing"

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestFilter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Filter()
		})
	}
}
