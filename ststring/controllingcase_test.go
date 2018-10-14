package ststring

import "testing"

func TestControllingCase(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestControllingCase"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ControllingCase()
		})
	}
}
