package kata

import "testing"

func TestIsSolved(t *testing.T) {
	type args struct {
		board [3][3]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test-1", args: args{board: [3][3]int{{0, 0, 1}, {0, 1, 2}, {2, 1, 0}}}, want: -1},
		{name: "Test-2", args: args{board: [3][3]int{{1, 1, 1}, {0, 2, 2}, {0, 0, 0}}}, want: 1},
		{name: "Test-3", args: args{board: [3][3]int{{2, 1, 2}, {2, 1, 1}, {1, 1, 2}}}, want: 1},
		{name: "Test-4", args: args{board: [3][3]int{{2, 1, 2}, {2, 1, 1}, {1, 2, 1}}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSolved(tt.args.board); got != tt.want {
				t.Errorf("IsSolved() = %v, want %v", got, tt.want)
			}
		})
	}
}
