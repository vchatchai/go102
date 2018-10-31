package stfilepath

import "testing"

func TestListDir(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestListDir"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListDir()
		})
	}
}
