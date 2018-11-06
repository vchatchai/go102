package stnet

import "testing"

func TestHeader(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Header"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Header()
		})
	}
}
