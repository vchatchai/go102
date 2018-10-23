package sttime

import "testing"

func TestTimeOut(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestTimeOut"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TimeOut()
		})
	}
}
