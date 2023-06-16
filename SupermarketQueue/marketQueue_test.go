package kata

import "testing"

func TestQueueTime(t *testing.T) {
	type args struct {
		customers []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Empty queue", args{[]int{}, 1}, 0},
		{"Single till", args{[]int{1, 2, 3, 4}, 1}, 10},
		{"Multiple tills", args{[]int{2, 2, 3, 3, 4, 4}, 2}, 9},
		{"Multiple tills", args{[]int{1, 2, 3, 4, 5}, 100}, 5},
		{"Multiple tills", args{[]int{5, 3, 4}, 1}, 12},
		{"Multiple tills", args{[]int{10, 2, 3, 3}, 2}, 10},
		{"Multiple tills", args{[]int{2, 3, 10}, 2}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueueTime(tt.args.customers, tt.args.n); got != tt.want {
				t.Errorf("QueueTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
