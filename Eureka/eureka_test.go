package kata

import (
	"reflect"
	"testing"
)

func TestSumDigPow(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{"1, 10", args{1, 10}, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"1, 100", args{1, 100}, []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9, 89}},
		{"10, 89", args{10, 89}, []uint64{89}},
		{"10, 100", args{10, 100}, []uint64{89}},
		{"90, 100", args{90, 100}, nil},
		{"89, 135", args{89, 135}, []uint64{89, 135}},
		{"12157692622039623503, 12157692622039625605", args{12157692622039623503, 12157692622039625605}, []uint64{12157692622039623539}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumDigPow(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumDigPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
