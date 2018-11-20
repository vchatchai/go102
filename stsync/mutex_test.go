package stsync

import "testing"

func TestMutex(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMutex"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Mutex()
		})
	}
}
