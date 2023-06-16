package kata

import "testing"

func TestGetCount(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"abracadabra", args{"abracadabra"}, 5},
		{"aeiou", args{"aeiou"}, 5},
		{"sdfghjkl", args{"sdfghjkl"}, 0},
		{"", args{""}, 0},
		{"banana", args{"banana"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := GetCount(tt.args.str); gotCount != tt.wantCount {
				t.Errorf("GetCount() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
