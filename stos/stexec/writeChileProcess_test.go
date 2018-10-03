package stexec

import "testing"

func TestWriteChileProcess(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestWriteChileProcess"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WriteChileProcess()
		})
	}

	t.Error("Done")
}
