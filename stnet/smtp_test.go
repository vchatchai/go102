package stnet

import "testing"

func TestSMTP(t *testing.T) {
	type args struct {
		email string
		pass  string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestSMTP", args{email: "", pass: ""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SMTP(tt.args.email, tt.args.pass)
		})
	}
}
