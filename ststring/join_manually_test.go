package ststring

import "testing"

func TestJoinManually(t *testing.T) {
	type args struct {
		value []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test Join", args{[]string{
			"FIRST_NAME = 'Jack' ",
			"INSURANCE_NO = 333444555 ",
			"EFFECTIVE_FROM = SYSDATE "}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JoinManually(tt.args.value)
		})
	}
}
