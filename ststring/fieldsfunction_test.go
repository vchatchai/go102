package ststring

import "testing"

func TestFieldFunction(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Field Function", args{"Mary*had,a%little_lamb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FieldFunction(tt.args.value)
		})
	}
}
