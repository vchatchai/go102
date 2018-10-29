package stencoding

import "testing"

func TestXML(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestXML"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			XML()
		})
	}
}
