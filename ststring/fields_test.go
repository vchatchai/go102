package ststring

import "testing"

func TestFields(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Fields", args{`Mary_had a little_lamb`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fields(tt.args.value)
		})
	}
}
