package main

import (
	"fmt"
)

func HumanReadableTime(seconds int) string {
	// your code here

	if seconds > 359999 {
		return "Max time exceeded"
	}

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
