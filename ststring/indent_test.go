package ststring

import "testing"

func TestIndentALL(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestIndentALL"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			IndentALL()
		})
	}
}
