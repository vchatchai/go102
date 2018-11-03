package list

import "testing"

func TestEmptyAppend(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestEmptyAppend"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EmptyAppend()
		})
	}
}
