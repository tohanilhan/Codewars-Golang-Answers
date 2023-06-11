package main

func alphanumeric(str string) bool {

	// if the string is empty, return false
	if len(str) < 1 {
		return false
	}

	/*
		- 48 - 57 are the ascii values for 0 - 9
		- 65 - 90 are the ascii values for A - Z
		- 97 - 122 are the ascii values for a - z
		- 95 is the ascii value for _
		- 32 is the ascii value for space
	*/

	// loop through the string and check each character
	for _, v := range str {
		// if string contains any non-alphanumeric characters, return false
		if v < 48 || (v > 57 && v < 65) || (v > 90 && v < 97) || v > 122 {
			return false
		}
	}

	return true
}
