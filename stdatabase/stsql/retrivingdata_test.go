package stsql

import "testing"

func TestRetrivingData(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestRetrivingData"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RetrivingData()
		})
	}
}
