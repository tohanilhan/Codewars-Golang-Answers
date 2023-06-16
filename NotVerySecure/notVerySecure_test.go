package kata

import "testing"

func Test_alphanumeric(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// Expect(alphanumeric("a")).To(Equal(true))
		// Expect(alphanumeric("Mazinkaiser")).To(Equal(true))
		// Expect(alphanumeric("hello world_")).To(Equal(false))
		// Expect(alphanumeric("PassW0rd")).To(Equal(true))
		// Expect(alphanumeric("     ")).To(Equal(false))
		// Expect(alphanumeric("")).To(Equal(false))
		// Expect(alphanumeric("\n\t\n")).To(Equal(false))
		// Expect(alphanumeric("ciao\n$$_")).To(Equal(false))
		// Expect(alphanumeric("__ * __")).To(Equal(false))
		// Expect(alphanumeric("&)))(((")).To(Equal(false))
		// Expect(alphanumeric("43534h56jmTHHF3k")).To(Equal(true))
		{".*?", args{".*?"}, false},
		{"a", args{"a"}, true},
		{"Mazinkaiser", args{"Mazinkaiser"}, true},
		{"hello world_", args{"hello world_"}, false},
		{"PassW0rd", args{"PassW0rd"}, true},
		{"     ", args{"     "}, false},
		{"", args{""}, false},
		{"\n\t\n", args{"\n\t\n"}, false},
		{"ciao\n$$_", args{"ciao\n$$_"}, false},
		{"__ * __", args{"__ * __"}, false},
		{"&)))(((", args{"&)))((("}, false},
		{"43534h56jmTHHF3k", args{"43534h56jmTHHF3k"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alphanumeric(tt.args.str); got != tt.want {
				t.Errorf("alphanumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
