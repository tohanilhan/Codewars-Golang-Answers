package kata

import (
	"reflect"
	"testing"
)

func TestRemovNb(t *testing.T) {
	type args struct {
		m uint64
	}
	tests := []struct {
		name string
		args args
		want [][2]uint64
	}{
		{"26", args{26}, [][2]uint64{{15, 21}, {21, 15}}},
		{"100", args{100}, nil},
		{"101", args{101}, [][2]uint64{{55, 91}, {91, 55}}},
		{"102", args{102}, [][2]uint64{{70, 73}, {73, 70}}},
		{"110", args{110}, [][2]uint64{{70, 85}, {85, 70}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemovNb(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemovNb() = %v, want %v", got, tt.want)
			}
		})
	}
}
