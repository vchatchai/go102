package sttime

import "testing"

func TestDelay(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestDelay"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Delay()
		})
	}
}
