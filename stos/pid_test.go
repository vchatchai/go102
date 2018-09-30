package stos

import "testing"

func TestPID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test PID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PID()
		})
	}

	t.Error("Done.")
}
