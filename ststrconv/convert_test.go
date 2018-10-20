package strconv

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestConvert"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Convert()
		})
	}
}
