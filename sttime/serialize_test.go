package sttime

import "testing"

func TestSerialize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSerialize"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Serialize()
		})
	}
}
