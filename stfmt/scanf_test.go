package stfmt

import "testing"

func TestScanf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestScanf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Scanf()
		})
	}
}
