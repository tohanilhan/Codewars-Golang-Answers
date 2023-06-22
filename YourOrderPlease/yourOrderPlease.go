package kata

import (
	"fmt"
	"regexp"
	"strings"
)

func Order(sentence string) string {
	// your code:

	// Split the sentence into an array of strings
	sentenceArr := strings.Split(sentence, " ")
	newArr := make([]string, len(sentenceArr))
	// Iterate through the sentence
	for _, word := range sentenceArr {
		// Check if the word contains the number

		// Extract the number from 1 to 9 from the word
		re := regexp.MustCompile("[1-9]+")
		numStr := re.FindString(word)

		// Convert the number to an integer
		num := 0
		if numStr != "" {
			fmt.Sscanf(numStr, "%d", &num)
		}

		// Iterate through the sentence array
		for i := 0; i < len(sentenceArr); i++ {
			// If the number is equal to the index of the word in the sentence, add the word to the new array
			if num == i+1 {
				newArr[i] = word
			}
		}
	}

	// Join the sentence array into a string
	sentence = strings.Join(newArr, " ")

	return sentence
}
