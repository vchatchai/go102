package stnet

import "testing"

func TestInterfaces(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test interfaces"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Interfaces()
		})
	}
}
