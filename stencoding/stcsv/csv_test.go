package stcsv

import "testing"

func TestCSV(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestCSV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CSV()
		})
	}
}
