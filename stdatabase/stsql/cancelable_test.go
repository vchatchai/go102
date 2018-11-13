package stsql

import "testing"

func TestCancelable(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCancelable"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Cancelable()
		})
	}
}
