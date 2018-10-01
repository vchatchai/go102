package stexec

import "testing"

func TestExec(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Exec"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Exec()
		})
	}
}
