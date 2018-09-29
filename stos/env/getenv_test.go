package env

import "testing"

func TestGetENV(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test GETENV GOPATH", args{"GOPATH"}},
		{"Test GETENV test", args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetENV(tt.args.key)
		})
	}
	t.Error()
}
