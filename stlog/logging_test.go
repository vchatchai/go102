package stlog

import "testing"

func TestLogging(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestLogging"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Logging()
		})
	}
}
