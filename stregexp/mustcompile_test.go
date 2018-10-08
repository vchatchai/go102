package stregexp

import "testing"

func TestMustCompile(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Must Compile", args{"Mary*had,a%little_lamb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MustCompile(tt.args.value)
		})
	}
}
