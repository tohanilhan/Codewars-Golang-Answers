package kata

import "testing"

func TestSumOfIntervals(t *testing.T) {
	type args struct {
		intervals [][2]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1", args{[][2]int{{1, 5}}}, 4},
		{"Test 2", args{[][2]int{{1, 5}, {6, 10}}}, 8},
		{"Test 3", args{[][2]int{{1, 5}, {1, 5}}}, 4},
		{"Test 4", args{[][2]int{{1, 4}, {7, 10}, {3, 5}}}, 7},
		{"Test 5", args{[][2]int{{-1000000000, 1000000000}}}, 2000000000},
		{"Test 6", args{[][2]int{{0, 20}, {-100000000, 10}, {30, 40}}}, 100000030},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("SumOfIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}
