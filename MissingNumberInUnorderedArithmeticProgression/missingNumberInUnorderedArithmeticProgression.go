package kata

import "sort"

func FindMissingNumber(seq []int) int {
	sort.Ints(seq)
	n := len(seq)

	// Calculate the expected sum based on the arithmetic series formula: n/2 * (a + l)
	expectedSum := (n + 1) * (seq[0] + seq[n-1]) / 2

	// Calculate the actual sum of the elements in the sequence
	actualSum := 0
	for _, v := range seq {
		actualSum += v
	}

	// The missing number is the difference between the expected sum and the actual sum
	missingNumber := expectedSum - actualSum

	return missingNumber
}
