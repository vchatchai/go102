package streflect

import "testing"

func TestStructTag(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestStructTag"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StructTag()
		})
	}
}
