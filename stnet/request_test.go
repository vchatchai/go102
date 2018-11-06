package stnet

import "testing"

func TestRequest(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Request"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Request()
		})
	}
}
