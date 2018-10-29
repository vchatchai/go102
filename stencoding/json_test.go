package stencoding

import (
	"fmt"
	"testing"
)

func TestJson(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestJson"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Json()
		})
	}

	fmt.Println("Done.")
}
