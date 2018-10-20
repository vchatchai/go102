package stmath

import "testing"

func TestComplex(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestComplex"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Complex()
		})
	}
}
