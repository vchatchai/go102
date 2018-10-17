package stmath

import "testing"

func TestRound(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Round"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Round()
		})
	}
}
