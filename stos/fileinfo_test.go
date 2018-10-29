package stos

import "testing"

func TestFileInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestFileInfo"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FileInfo()
		})
	}
}
