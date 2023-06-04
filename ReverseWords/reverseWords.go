package main

import (
	"strings"
)

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	reversedStr := ""
	for i := range words {
		if len(words[i]) >= 5 {
			reversedStr += Reverse(words[i]) + " "
		} else {
			reversedStr += words[i] + " "
		}

	}

	reversedStr = strings.TrimSuffix(reversedStr, " ")

	return reversedStr
} // SpinWords

// Reverse reverses a string
func Reverse(s string) string {
	chars := []rune(s)
	start := 0
	end := len(chars) - 1
	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}
	return string(chars)
}
