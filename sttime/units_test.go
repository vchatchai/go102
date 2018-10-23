package sttime

import "testing"

func TestUnit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Unit"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Unit()
		})
	}
}
