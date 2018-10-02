package stexec

import "testing"

func TestReadBufferChildProcess(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestReadBufferChildProcess"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadBufferChildProcess()
		})
	}
}
