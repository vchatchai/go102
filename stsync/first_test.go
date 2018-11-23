package stsync

import "testing"

func TestFirst(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestFirst"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			First()
		})
	}
}
