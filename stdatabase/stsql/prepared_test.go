package stsql

import "testing"

func TestPrepare(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestPrepare"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Prepare()
		})
	}
}
