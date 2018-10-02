package stexec

import "testing"

func TestReadStdOut(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test ReadStdOut"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadStdOut()
		})
	}
}
