package kata

import "testing"

func TestFindMissingNumber(t *testing.T) {
	type args struct {
		seq []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test-1", args: args{seq: []int{3, 9, 1, 11, 13, 5}}, want: 7},
		{name: "Test-2", args: args{seq: []int{5, -1, 0, 3, 4, -3, 2, -2}}, want: 1},
		{name: "Test-3", args: args{seq: []int{2, -2, 8, -8, 4, -4, 6, -6}}, want: 0},
		{name: "Test-4", args: args{seq: []int{1, 1, 1, 1, 1, 1, 1, 1}}, want: 1}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMissingNumber(tt.args.seq); got != tt.want {
				t.Errorf("FindMissingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
