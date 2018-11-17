package stnet

import "testing"

func TestMiddleWare(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMiddleWare"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MiddleWare()
		})
	}
}
