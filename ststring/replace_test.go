package ststring

import "testing"

func TestRepalce(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Replace"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Repalce()
		})
	}
}
