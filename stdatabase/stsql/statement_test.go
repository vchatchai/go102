package stsql

import "testing"

func TestStatment(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestStatment"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Statment()
		})
	}
}
