package stio

import "testing"

func TestPipe(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Pipe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pipe()
		})
	}
}
