package stnet

import "testing"

func TestRedirects(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Redirect"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Redirects()
		})
	}
}
