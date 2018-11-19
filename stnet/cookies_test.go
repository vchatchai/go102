package stnet

import "testing"

func TestCookies(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCookies"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Cookies()
		})
	}
}
