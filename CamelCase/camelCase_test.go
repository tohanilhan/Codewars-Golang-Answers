package kata

import "testing"

func TestToCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{""}, ""},
		{"the", args{"the"}, "the"},
		{"The", args{"The"}, "The"},
		{"theStealthWarrior", args{"theStealthWarrior"}, "theStealthWarrior"},
		{"TheStealthWarrior", args{"TheStealthWarrior"}, "TheStealthWarrior"},
		{"the_stealth_warrior", args{"the_stealth_warrior"}, "theStealthWarrior"},
		{"The_Stealth_Warrior", args{"The_Stealth_Warrior"}, "TheStealthWarrior"},
		{"the-stealth_warrior", args{"the-stealth_warrior"}, "theStealthWarrior"},
		{"The-Stealth-Warrior", args{"The-Stealth-Warrior"}, "TheStealthWarrior"},
		{"A-B-C", args{"A-B-C"}, "ABC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args.s); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
