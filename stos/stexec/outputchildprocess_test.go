package stexec

import "testing"

func TestOutputChildProcess(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestOutputChildProcess "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OutputChildProcess()
		})
	}
}
