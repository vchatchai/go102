package stos

import "testing"

func TestFileSeek(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestFileSeek"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FileSeek()
		})
	}
}
