package main

import "testing"

func TestFirstNonRepeating(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"a", args{"a"}, "a"},
		{"stress", args{"stress"}, "t"},
		{"moonmen", args{"moonmen"}, "e"},
		{"abba", args{"abba"}, ""},
		{"aa", args{"aa"}, ""},
		{"~><#~><", args{"~><#~><"}, "#"},
		{"hello world, eh?", args{"hello world, eh?"}, "w"},
		{"sTreSS", args{"sTreSS"}, "T"},
		{"Go hang a salami, I'm a lasagna hog!", args{"Go hang a salami, I'm a lasagna hog!"}, ","},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstNonRepeating(tt.args.str); got != tt.want {
				t.Errorf("FirstNonRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
