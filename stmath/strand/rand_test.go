package strand

import "testing"

func TestRand(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestRand"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Rand()
		})
	}
}
