package stjson

import "testing"

func TestFileConfigulation(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestFileConfigulation"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FileConfigulation()
		})
	}
	t.Error("done.")
}
