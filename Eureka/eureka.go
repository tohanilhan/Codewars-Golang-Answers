package main

func SumDigPow(a, b uint64) []uint64 {
	var result []uint64

	for i := a; i <= b; i++ {
		if isEurekaNumber(i) {
			result = append(result, i)
		}
	}

	return result
}

func isEurekaNumber(n uint64) bool {
	digits := GetDigits(n)

	var sum uint64
	for i, v := range digits {
		sum += customPow(v, uint64(i+1))
	}

	return sum == n
}

func GetDigits(n uint64) []uint64 {
	var digits []uint64

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	reverseSlice(digits)
	return digits
}

func reverseSlice(slice []uint64) {
	start := 0
	end := len(slice) - 1
	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start++
		end--
	}
}

func customPow(base, exponent uint64) uint64 {
	result := uint64(1)
	for i := uint64(0); i < exponent; i++ {
		result *= base
	}
	return result
}
