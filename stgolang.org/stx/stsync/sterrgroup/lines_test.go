package sterrgroup

import "testing"

func TestLines(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestLines"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Lines()
		})
	}
}
