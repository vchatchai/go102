package stos

import "testing"

func TestFileMode(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestFileMode"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FileMode()
		})
	}
}
