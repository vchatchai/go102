package stsync

import "testing"

func TestMap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMap"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Map()
		})
	}
}
