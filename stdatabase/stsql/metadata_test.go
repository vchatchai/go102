package stsql

import "testing"

func TestMetadata(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMetadata"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Metadata()
		})
	}
}
