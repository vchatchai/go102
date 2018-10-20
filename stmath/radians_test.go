package stmath

import "testing"

func TestRadians(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestRadians"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Radians()
		})
	}
}
