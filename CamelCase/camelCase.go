package main

import (
	"strings"

	"github.com/fatih/color"
)

func main() {
	arr := [...][2]string{
		{"the-stealth-warrior", "theStealthWarrior"},
		{"The_Stealth_Warrior", "TheStealthWarrior"},
		{"the-stealth_warrior", "theStealthWarrior"},
		{"The_Stealth_Warrior", "TheStealthWarrior"},
		{"The_Stealth-Warrior", "TheStealthWarrior"},
		{"A-B-C", "ABC"},
	}

	for _, v := range arr {
		var input = v[0]
		var output = v[1]
		var result = ToCamelCase(input)
		if result == output {
			color.Green("PASS")
		} else {
			color.Red("Expected: %s, got: %s and input was: %s", output, result, input)
		}
	}
}

func ToCamelCase(s string) string {
	// your code

	if len(s) == 0 {
		return ""
	}

	// trim string for underscores and dashes
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Replace(s, "-", " ", -1)

	// split string into words
	words := strings.Split(s, " ")

	// loop through words
	for i, word := range words {
		// if the first word is lowecase, leave it that way
		// otherwise, capitalize the first letter
		if i == 0 && word[0] >= 97 && word[0] <= 122 {
			continue
		} else {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}

	}

	// join words back together
	s = strings.Join(words, "")

	return s
}
