package stnet

import "testing"

func TestPush(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestPush"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Push()
		})
	}
}
