package kata

import "testing"

func TestZeros(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Zeros(0)", args{0}, 0},
		{"Zeros(6)", args{6}, 1},
		{"Zeros(12)", args{12}, 2},
		{"Zeros(30)", args{30}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zeros(tt.args.n); got != tt.want {
				t.Errorf("Zeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
