package ststring

import "testing"

func TestSplit(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Split", args{"Mary_had a little_lamb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Split(tt.args.value)
		})
	}
}
