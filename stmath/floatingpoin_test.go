package stmath

import "testing"

func TestFloatingPoint(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestFloatingPoint"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FloatingPoint()
		})
	}
}
