package kata

func Zeros(n int) int {
	// your code here
	trailingZeros := 0

	// divide n by 5, 25, 125, 625, ... until n/i < 1
	// add the result to trailingZeros
	for i := 5; n/i >= 1; i *= 5 {
		trailingZeros += n / i
	}

	return trailingZeros
}
