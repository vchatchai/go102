package datastructure

import "testing"

func TestLists(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test List"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Lists()
		})
	}
}
