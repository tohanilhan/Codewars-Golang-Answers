package kata

import "testing"

func TestDuplicateEncode(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"din", args{"din"}, "((("},
		{"recede", args{"recede"}, "()()()"},
		{"Success", args{"Success"}, ")())())"},
		{"(( @", args{"(( @"}, "))(("},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DuplicateEncode(tt.args.word); got != tt.want {
				t.Errorf("DuplicateEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}
