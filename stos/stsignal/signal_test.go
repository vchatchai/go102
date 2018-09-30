package stsignal

import "testing"

func TestSignal(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Signal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Signal()
		})
	}
	t.Error("Done.")
}
