package kata

import "testing"

func TestExpSum(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "Test-1", args: args{n: 1}, want: 1},
		{name: "Test-2", args: args{n: 2}, want: 2},
		{name: "Test-3", args: args{n: 3}, want: 3},
		{name: "Test-4", args: args{n: 4}, want: 5},
		{name: "Test-5", args: args{n: 5}, want: 7},
		{name: "Test-6", args: args{n: 10}, want: 42},
		{name: "Test-7", args: args{n: 50}, want: 204226},
		{name: "Test-8", args: args{n: 80}, want: 15796476},
		{name: "Test-9", args: args{n: 100}, want: 190569292},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExpSum(tt.args.n); got != tt.want {
				t.Errorf("ExpSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
