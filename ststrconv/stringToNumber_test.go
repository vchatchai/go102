package strconv

import "testing"

func TestStringToNumber(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestStringToNumber"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StringToNumber()
		})
	}
}
