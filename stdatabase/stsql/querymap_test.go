package stsql

import "testing"

func TestQuerymap(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestQuerymap"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Querymap()
		})
	}
}
