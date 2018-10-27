package stio

import "testing"

func TestMultiWR(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"MultiWR"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MultiWR()
		})
	}
}
