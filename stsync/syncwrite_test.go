package stsync

import "testing"

func TestSyncwrite(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestSyncwrite"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Syncwrite()
		})
	}
}
