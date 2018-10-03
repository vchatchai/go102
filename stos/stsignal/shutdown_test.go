package stsignal

import "testing"

func TestShutdown(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Shutdown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Shutdown()
		})
	}
}
