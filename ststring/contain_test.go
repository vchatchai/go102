package ststring

import "testing"

func TestContain(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Contain", args{"Mary had a little lamb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Contain(tt.args.word)
		})
	}
}
