package starchive

import "testing"

func TestZip(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test ZIp"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Zip()
		})
	}
}
