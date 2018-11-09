package stsql

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestConnect"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Connect()
		})
	}
}
