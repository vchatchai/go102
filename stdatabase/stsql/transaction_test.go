package stsql

import "testing"

func TestTransaction(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestTransaction"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Transaction()
		})
	}
}
