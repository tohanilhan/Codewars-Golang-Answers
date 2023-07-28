package kata

func Rot13(msg string) string {
	// Your code here
	newStr := ""
	for _, r := range msg {
		if r >= 'a' && r <= 'z' {
			// Rotate lowercase letters 13 places.
			if r > 'm' {
				newStr += string(r - 13)
			} else {
				newStr += string(r + 13)
			}
		} else if r >= 'A' && r <= 'Z' {
			// Rotate uppercase letters 13 places.
			if r > 'M' {
				newStr += string(r - 13)
			} else {
				newStr += string(r + 13)
			}
		} else {
			newStr += string(r)
		}
	}
	return newStr
}
