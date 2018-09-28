package env

import "testing"

func TestGetENV(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"GET ENV"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetENV()
		})
	}

	t.Error("DONE.")
}
