package stencoding

import "testing"

func TestRWbinary(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TESTRWBianry"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RWbinary()
		})
	}
}
