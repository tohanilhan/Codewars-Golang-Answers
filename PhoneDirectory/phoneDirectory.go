package kata

import (
	"regexp"
	"strings"
)

func Phone(dir, num string) string {
	// your code:

	r := regexp.MustCompile(`\d{1,2}-\d{3}-\d{3}-\d{4}`)

	if !r.MatchString(num) {
		return "Error => Not found: " + num
	}

	// count the number of times the number appears in the directory
	count := strings.Count(dir, num)

	// if the number appears more than once, return an error
	if count > 1 {
		return "Error => Too many people: " + num
	}

	// Split the directory into an array of strings
	dirArr := strings.Split(dir, "\n")

	// Iterate through the directory
	for _, entry := range dirArr {
		// Check if the entry contains the number
		if regexp.MustCompile(num).MatchString(entry) {

			// Extract the name
			name := regexp.MustCompile(`<(.+)>`).FindStringSubmatch(entry)[1]

			// Remove the name from the entry
			entry = regexp.MustCompile(`<(.+)>`).ReplaceAllString(entry, "")

			// remove the number from the entry
			entry = regexp.MustCompile(num).ReplaceAllString(entry, "")

			// extract the address
			address := regexp.MustCompile(`[^\w\s.-]`).ReplaceAllString(entry, " ")

			// remove underscores
			address = regexp.MustCompile(`_`).ReplaceAllString(address, " ")

			// remove the extra spaces
			address = regexp.MustCompile(`\s{2,}`).ReplaceAllString(address, " ")

			// remove the leading and trailing spaces
			address = strings.TrimSpace(address)

			// Return the result
			return "Phone => " + num + ", Name => " + name + ", Address => " + address
		}
	}

	return "Error => Not found: " + num
}
