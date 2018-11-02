package stcrypto

import "testing"

func TestComparison(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestComparison"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Comparison()
		})
	}
}
