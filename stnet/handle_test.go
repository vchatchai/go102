package stnet

import "testing"

func TestHandle(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHandle"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Handle()
		})
	}
}
