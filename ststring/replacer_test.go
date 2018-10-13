package ststring

import "testing"

func TestReplacer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestReplacer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Replacer()
		})
	}
}
