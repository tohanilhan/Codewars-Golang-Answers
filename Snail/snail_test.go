package kata

import (
	"reflect"
	"testing"
)

func TestSnail(t *testing.T) {
	type args struct {
		snaipMap [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				snaipMap: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "Test 2",
			args: args{
				snaipMap: [][]int{
					{1, 2, 3, 1},
					{4, 5, 6, 4},
					{7, 8, 9, 7},
					{7, 8, 9, 7},
				},
			},
			want: []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8},
		},
		{
			name: "Test 3",
			args: args{
				snaipMap: [][]int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Snail(tt.args.snaipMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Snail() = %v, want %v", got, tt.want)
			}
		})
	}
}
