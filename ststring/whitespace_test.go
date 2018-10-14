package ststring

import "testing"

func TestWhiteSpace(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestWhiteSpace"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WhiteSpace()
		})
	}
}
