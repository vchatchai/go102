package sttime

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Parse"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Parse()
		})
	}
}
