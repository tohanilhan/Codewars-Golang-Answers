package kata

import "testing"

func TestIsValidWalk(t *testing.T) {
	type args struct {
		walk []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{"10 min walk", args{[]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}}, true},
		{"Not a 10 min walk", args{[]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}}, false},
		{"Not a 10 min walk", args{[]rune{'w'}}, false},
		{"Not a 10 min walk", args{[]rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}}, false},
		{"Not a 10 min walk", args{[]rune{'e', 'e', 'e', 'e', 'w', 'w', 's', 's', 's', 's'}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidWalk(tt.args.walk); got != tt.want {
				t.Errorf("IsValidWalk() = %v, want %v", got, tt.want)
			}
		})
	}
}
