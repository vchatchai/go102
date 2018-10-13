package stregexp

import "testing"

func TestFindSubString(t *testing.T) {
	type args struct {
		refString string
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestFindSubString", args{`[{ \"email\": \"email@example.com\" \
		"phone\": 555467890},
		{ \"email\": \"other@domain.com\" \
		"phone\": 555467890}]`}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			FindSubString(tt.args.refString)
		})
	}
}
