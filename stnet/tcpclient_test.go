package stnet

import "testing"

func TestTcpClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestTcpClient"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TcpClient()
		})
	}
}
