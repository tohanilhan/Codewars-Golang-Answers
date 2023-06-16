package kata

func DigitalRoot(n int) int {
	// ...

	// if n bigger than or equal to 10 , sum the digits
	for n >= 10 {
		sum := 0
		// sum the digits
		for n > 0 {
			// add the last digit to the sum
			sum += n % 10
			// remove the last digit
			n /= 10
		}
		// assign the sum to n
		n = sum

		// if n is smaller than 10, return n
	}
	return n
}
