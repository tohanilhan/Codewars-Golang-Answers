package kata

import "testing"

func TestPartitions(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1}, 1},
		{"3", args{3}, 3},
		{"4", args{4}, 5},
		{"5", args{5}, 7},
		{"6", args{6}, 11},
		{"7", args{7}, 15},
		{"8", args{8}, 22},
		{"9", args{9}, 30},
		{"10", args{10}, 42},
		{"11", args{11}, 56},
		{"12", args{12}, 77},
		{"13", args{25}, 1958},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Partitions(tt.args.n); got != tt.want {
				t.Errorf("Partitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
