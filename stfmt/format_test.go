package stfmt

import "testing"

func TestFormat(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Format"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Format()
		})
	}
}
