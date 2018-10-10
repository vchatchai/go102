package buildin

import "testing"

func TestConcatCopy(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Concat Copy", args{[]string{"This ", "is ", "even ",
			"more ", "performant "}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConcatCopy(tt.args.values)
		})
	}
}
