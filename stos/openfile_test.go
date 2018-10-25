package stos

import "testing"

func TestOpenFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestOpenFile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenFile()
		})
	}
}
