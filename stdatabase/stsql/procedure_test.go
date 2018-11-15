package stsql

import "testing"

func TestProcedure(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestProcedure"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Procedure()
		})
	}
}
