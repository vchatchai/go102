package stnet

import "testing"

func TestHandleGroup(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestHandleGroup"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HandleGroup()
		})
	}
}
