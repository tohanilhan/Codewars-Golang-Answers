package main

import "testing"

func TestHigh(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"man i need a taxi up to ubud", args{"man i need a taxi up to ubud"}, "taxi"},
		{"what time are we climbing up the volcano", args{"what time are we climbing up the volcano"}, "volcano"},
		{"take me to semynak", args{"take me to semynak"}, "semynak"},
		{"aa b", args{"aa b"}, "aa"},
		{"b aa", args{"b aa"}, "b"},
		{"bb d", args{"bb d"}, "bb"},
		{"d bb", args{"d bb"}, "d"},
		{"aaa b", args{"aaa b"}, "aaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := High(tt.args.s); got != tt.want {
				t.Errorf("High() = %v, want %v", got, tt.want)
			}
		})
	}
}
