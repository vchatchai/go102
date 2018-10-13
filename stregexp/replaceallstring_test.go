package stregexp

import "testing"

func TestReplaceAllString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestReplaceAllString", args{"Mary had a little lamb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReplaceAllString(tt.args.value)
		})
	}
}
