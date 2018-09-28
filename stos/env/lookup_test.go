package env

import "testing"

func TestLookup(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestLookup GetValue", args{"GOPATH"}},
		{"TestLookup GetError", args{"GO"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Lookup(tt.args.key)
		})
	}
}
