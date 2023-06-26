package kata

import "testing"

func TestLCS(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test 1", args{"a", "b"}, ""},
		{"Test 2", args{"abcdef", "abc"}, "abc"},
		{"Test 3", args{"132535365", "123456789"}, "12356"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCS(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("LCS() = %v, want %v", got, tt.want)
			}
		})
	}
}
