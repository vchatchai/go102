package sttime

import "testing"

func TestToday(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Time Today"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Today()
		})
	}
}
