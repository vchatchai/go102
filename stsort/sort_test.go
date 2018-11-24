package stsort

import "testing"

func TestSort(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSort"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Sort()
		})
	}
}
