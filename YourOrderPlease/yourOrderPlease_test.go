package kata

import "testing"

func TestOrder(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "is2 Thi1s T4est 3a",
			args: args{
				sentence: "is2 Thi1s T4est 3a",
			},
			want: "Thi1s is2 3a T4est",
		},

		{
			name: "4of Fo1r pe6ople g3ood th5e the2",
			args: args{
				sentence: "4of Fo1r pe6ople g3ood th5e the2",
			},
			want: "Fo1r the2 g3ood 4of th5e pe6ople",
		},

		{
			name: "",
			args: args{
				sentence: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Order(tt.args.sentence); got != tt.want {
				t.Errorf("Order() = %v, want %v", got, tt.want)
			}
		})
	}
}
