package charmap

import "testing"

func TestCharset(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCharset"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Charset()
		})
	}
}
