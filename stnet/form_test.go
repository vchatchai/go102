package stnet

import "testing"

func TestParseForm(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestParseForm"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseForm()
		})
	}
}
