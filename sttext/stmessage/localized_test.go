package stmessage

import "testing"

func TestLocalized(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Localized"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Localized()
		})
	}
}
