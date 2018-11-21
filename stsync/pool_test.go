package stsync

import "testing"

func TestPool(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestPool"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pool()
		})
	}
}
