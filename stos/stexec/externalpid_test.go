package stexec

import "testing"

func TestExternalPid(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test External PID "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExternalPid()
		})
	}
}
