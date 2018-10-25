package stos

import "testing"

func TestStdouterr(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestStdouterr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Stdouterr()
		})
	}
}
