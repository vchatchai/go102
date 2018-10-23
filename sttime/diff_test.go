package sttime

import "testing"

func TestDiff(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Diff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Diff()
		})
	}
}
