package stsql

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestVerify(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestVerify"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Verify()
		})
	}
}
