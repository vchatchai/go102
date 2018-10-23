package sttime

import "testing"

func TestArithmetics(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestArithmetics"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Arithmetics()
		})
	}
}
