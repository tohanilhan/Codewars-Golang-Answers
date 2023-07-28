package kata

import "testing"

func TestRot13(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test-1", args: args{msg: "test"}, want: "grfg"},
		{name: "Test-2", args: args{msg: "Test"}, want: "Grfg"},
		{name: "Test-3", args: args{msg: "This is a test!"}, want: "Guvf vf n grfg!"},
		{name: "Test-4", args: args{msg: "abcdefghijklmnopqrstuvwxyz"}, want: "nopqrstuvwxyzabcdefghijklm"},
		{name: "Test-5", args: args{msg: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"}, want: "NOPQRSTUVWXYZABCDEFGHIJKLM"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rot13(tt.args.msg); got != tt.want {
				t.Errorf("Rot13() = %v, want %v", got, tt.want)
			}
		})
	}
}
