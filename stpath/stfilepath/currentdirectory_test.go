package stfilepath

import "testing"

func TestCurrentDirectory(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Current Directory"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CurrentDirectory()
		})
	}
}
