package main

import (
	"unicode"
)

func FirstNonRepeating(str string) string {
	//your code here

	// Create a map to store the count of each character
	charCount := make(map[rune]int)

	// create a slice to store the order of the characters
	var charOrder []rune

	// Loop through the string
	for _, v := range str {
		// If the character is not in the map, add it
		// doesn't matter if it is lower or upper case, it will be added
		if _, ok := charCount[unicode.ToLower(v)]; !ok {
			charCount[unicode.ToLower(v)] = 1
			charOrder = append(charOrder, v)
		} else {
			// If the character is in the map, increment the count
			charCount[unicode.ToLower(v)]++
		}
	}

	// Loop through the charOrder slice
	for _, v := range charOrder {
		// If the count of the character is 1, return it
		if charCount[unicode.ToLower(v)] == 1 {
			return string(v)
		}
	}

	return ""

}
