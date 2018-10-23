package sttime

import "testing"

func TestTimeZones(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test TimeZone"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TimeZones()
		})
	}
}
