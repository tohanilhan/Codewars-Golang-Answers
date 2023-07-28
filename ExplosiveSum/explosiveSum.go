package kata

func ExpSum(n uint64) uint64 {
	// your code here
	// Create a slice to store the number of partitions for each integer up to n.
	partitions := make([]uint64, n+1)
	partitions[0] = 1

	var current uint64
	// Iterate over each integer from 1 to n.
	for current = 1; current <= n; current++ {
		// For each current integer, update the partitions array by considering all possible sums.
		// Starting from current, iterate over each integer up to n.
		for i := current; i <= n; i++ {
			// Calculate the number of partitions by adding the number of partitions for the current integer
			// to the number of partitions for the remaining sum (i - current).
			partitions[i] += partitions[i-current]
		}
	}

	// Return the number of partitions for the given integer n.
	return partitions[n]
}
