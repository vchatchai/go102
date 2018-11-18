package stnet

import "testing"

func TestRedirect(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"TestRedirect",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Redirect()
		})
	}
}
