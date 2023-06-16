package kata

import "testing"

func TestDigitalRoot(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"16", args{16}, 7},
		{"195", args{195}, 6},
		{"992", args{992}, 2},
		{"167346", args{167346}, 9},
		{"0", args{0}, 0},
		{"942", args{942}, 6},
		{"132189", args{132189}, 6},
		{"493193", args{493193}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DigitalRoot(tt.args.n); got != tt.want {
				t.Errorf("DigitalRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
