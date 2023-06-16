package kata

import (
	"fmt"
	"strings"
)

func High(s string) string {
	// your code here

	highestScore := 0
	highestScoredWord := ""

	// split string into words
	words := strings.Split(s, " ")

	// loop through words
	for _, word := range words {

		// get score for each word
		score := getScore(word)

		// fmt.Println(word, "---->", score)
		// if score is higher than highestScore, set highestScore to score
		if score > highestScore {
			highestScore = score
			highestScoredWord = word
		}

	}

	return highestScoredWord
}

func getScore(word string) int {
	// create a map of letters and their scores
	letterScores := map[string]int{}

	// loop through alphabet and assign scores to letters

	for i := 97; i <= 122; i++ {
		// set the letter to the string representation of i
		letter := fmt.Sprintf("%c", i)

		// convert i to string and assign to map
		letterScores[letter] = i - 96

	}

	// set score to 0
	score := 0

	// loop through word
	for _, letter := range word {
		// add score for each letter to score
		score += letterScores[string(letter)]
	}
	// return score
	return score
}
