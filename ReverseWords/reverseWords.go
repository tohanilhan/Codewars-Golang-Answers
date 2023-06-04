package main

import (
	"strings"

	"github.com/fatih/color"
)

func main() {
	arr := [...][2]string{
		{"Hey fellow warriors", "Hey wollef sroirraw"},
		{"This is a test", "This is a test"},
		{"This is another test", "This is rehtona test"},
		{"You are almost to the last test", "You are tsomla to the last test"},
		{"Just kidding there is still one more", "Just gniddik ereht is llits one more"},
		{"Seriously this is the last one", "ylsuoireS this is the last one"},
	}

	for _, v := range arr {
		var input = v[0]
		var output = v[1]
		var result = SpinWords(input)
		if result == output {
			color.Green("PASS")
		} else {
			color.Red("Expected:", output)
		}
	}
}

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
