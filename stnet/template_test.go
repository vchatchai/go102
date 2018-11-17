package stnet

import "testing"

func TestTemplate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestTemplate"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Template()
		})
	}
}
