package kata

import "testing"

func TestSpinWords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Hey fellow warriors", args{"Hey fellow warriors"}, "Hey wollef sroirraw"},
		{"This is a test", args{"This is a test"}, "This is a test"},
		{"This is another test", args{"This is another test"}, "This is rehtona test"},
		{"You are almost to the last test", args{"You are almost to the last test"}, "You are tsomla to the last test"},
		{"Just kidding there is still one more", args{"Just kidding there is still one more"}, "Just gniddik ereht is llits one more"},
		{"Seriously this is the last one", args{"Seriously this is the last one"}, "ylsuoireS this is the last one"},
		{"Welcome", args{"Welcome"}, "emocleW"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpinWords(tt.args.str); got != tt.want {
				t.Errorf("SpinWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
