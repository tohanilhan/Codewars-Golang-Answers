package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	/*
		fmt.Println(HumanReadableTime(0))      // "00:00:00"
		fmt.Println(HumanReadableTime(59))     // "00:00:59"
		fmt.Println(HumanReadableTime(60))     // "00:01:00"
		fmt.Println(HumanReadableTime(3599))   // "00:59:59"
		fmt.Println(HumanReadableTime(3600))   // "01:00:00"
		fmt.Println(HumanReadableTime(45296))  // "12:34:56"
		fmt.Println(HumanReadableTime(86399))  // "23:59:59"
		fmt.Println(HumanReadableTime(86400))  // "24:00:00"
		fmt.Println(HumanReadableTime(359999)) // "99:59:59"
	*/

	// Test cases
	var arr = []struct {
		input  int
		output string
	}{
		{0, "00:00:00"},
		{59, "00:00:59"},
		{60, "00:01:00"},
		{3599, "00:59:59"},
		{3600, "01:00:00"},
		{45296, "12:34:56"},
		{86399, "23:59:59"},
		{86400, "24:00:00"},
		{359999, "99:59:59"},
	}

	for _, v := range arr {
		var input = v.input
		var output = v.output
		var result = HumanReadableTime(input)
		if result == output {
			color.Green("PASS")
		} else {
			color.Red("Expected:", output)
		}
	}

}

func HumanReadableTime(seconds int) string {
	// your code here
	humanReadableTime := ""

	// Hours
	hours := seconds / 3600
	if hours < 10 {
		humanReadableTime += "0"
	}
	humanReadableTime += fmt.Sprintf("%d", hours) + ":"

	// Minutes
	minutes := (seconds % 3600) / 60
	if minutes < 10 {
		humanReadableTime += "0"
	}

	humanReadableTime += fmt.Sprintf("%d", minutes) + ":"

	// Seconds
	seconds = seconds % 60
	if seconds < 10 {
		humanReadableTime += "0"
	}
	humanReadableTime += fmt.Sprintf("%d", seconds)

	return humanReadableTime
}
