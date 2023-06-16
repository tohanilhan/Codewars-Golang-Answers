package kata

import "testing"

func TestHumanReadableTime(t *testing.T) {
	type args struct {
		seconds int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{0}, "00:00:00"},
		{"59", args{59}, "00:00:59"},
		{"60", args{60}, "00:01:00"},
		{"3599", args{3599}, "00:59:59"},
		{"3600", args{3600}, "01:00:00"},
		{"45296", args{45296}, "12:34:56"},
		{"86399", args{86399}, "23:59:59"},
		{"86400", args{86400}, "24:00:00"},
		{"359999", args{359999}, "99:59:59"},
		{"360000", args{360000}, "Max time exceeded"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HumanReadableTime(tt.args.seconds); got != tt.want {
				t.Errorf("HumanReadableTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
