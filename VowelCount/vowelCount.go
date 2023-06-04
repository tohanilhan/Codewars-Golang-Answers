package main

import "fmt"

func main() {
	fmt.Println(GetCount("abracadabra"))
}

func GetCount(str string) (count int) {
	// Enter solution here

	// set ascii values for vowels
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}

	// loop through string
	for _, v := range str {
		for _, vowel := range vowels {
			if v == vowel {
				count++
			}
		}
	}

	return count
}
