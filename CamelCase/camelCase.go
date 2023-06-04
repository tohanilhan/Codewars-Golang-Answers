package main

import (
	"strings"
)

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
