package sttime

import (
	"testing"
)

func TestTicker(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Ticker"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ticker()
		})
	}
}
