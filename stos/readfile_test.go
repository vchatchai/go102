package stos

import "testing"

func TestReadFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestReadFile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadFile()
		})
	}
}
