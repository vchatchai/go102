package struntime

import "testing"

func TestVersion(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Print Version"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Version()
		})
	}
}
