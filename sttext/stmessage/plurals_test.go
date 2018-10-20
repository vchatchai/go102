package stmessage

import "testing"

func TestPlurals(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestPlurals"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Plurals()
		})
	}
}
