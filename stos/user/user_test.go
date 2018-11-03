package user

import "testing"

func TestUser(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestUser"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			User()
		})
	}
}
