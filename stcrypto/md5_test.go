package stcrypto

import "testing"

func TestMD5Test(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test MD5"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MD5Test()
		})
	}
}
