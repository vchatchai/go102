package sttime

import "testing"

func TestEpoch(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestEpoch"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Epoch()
		})
	}
}
