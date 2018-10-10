package stbyte

import "testing"

func TestConcatBuffer(t *testing.T) {
	type args struct {
		strings []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Concat Buffer", args{[]string{"This ", "is ", "even ",
			"more ", "performant "}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConcatBuffer(tt.args.strings)
		})
	}
}
