package stnet

import "testing"

func TestRest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TEST Rest Service"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rest()
		})
	}
}
