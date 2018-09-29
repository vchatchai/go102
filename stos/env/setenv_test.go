package env

import "testing"

func TestSetENV(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TEST ENV", args{"test", "testvalue"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetENV(tt.args.key, tt.args.value)
		})
	}
	t.Error()
}
