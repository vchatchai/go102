package stio

import "testing"

func TestTempFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test TempFile"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TempFile()
		})
	}
}
