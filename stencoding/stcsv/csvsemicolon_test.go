package stcsv

import "testing"

func TestCSVSemiColon(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCSVSemiColon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CSVSemiColon()
		})
	}
}
