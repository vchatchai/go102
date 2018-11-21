package stsync

import "testing"

func TestOnce(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestOnce"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Once()
		})
	}
}
