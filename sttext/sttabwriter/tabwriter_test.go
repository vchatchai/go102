package sttabwriter

import "testing"

func TestTabWriter(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test TabWriter"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TabWriter()
		})
	}
}
