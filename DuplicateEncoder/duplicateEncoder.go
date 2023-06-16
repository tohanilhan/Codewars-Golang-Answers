package kata

import "strings"

// DuplicateEncode takes a string and returns a new string where
func DuplicateEncode(word string) string {
	// Enter solution here

	newWord := ""

	word = strings.ToLower(word)

	for _, v := range word {
		// if v is in word more than once, return ")"
		if strings.Count(word, string(v)) > 1 {
			newWord += ")"
		} else {
			// otherwise, return "("
			newWord += "("
		}

	}

	return newWord
}

// This solution is the same as above, but uses a for loop instead of strings.Count

// func DuplicateEncode(word string) string {
// 	// Enter solution here

// 	newWord := ""

// 	word = strings.ToLower(word)

// 	for _, v := range word {
// 		count := 0
// 		// if v is in word more than once, return ")"
// 		// count the number of times v is in word
// 		for _, v2 := range word {
// 			if v == v2 {
// 				count++
// 			}
// 		}
// 		if count > 1 {
// 			newWord += ")"
// 		} else {
// 			// otherwise, return "("
// 			newWord += "("
// 		}
// 	}
// 	return newWord
// }
