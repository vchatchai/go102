package stio

import "testing"

func TestWriteFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestWriteFile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteFile()
		})
	}
}
