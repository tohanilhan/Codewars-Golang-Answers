package kata

import "sort"

func RemovNb(m uint64) [][2]uint64 {
	var sum uint64 = (m * (m + 1)) / 2

	var result [][2]uint64

	// Iterate through all possible values of a
	for a := uint64(1); a <= m; a++ {
		// Calculate the corresponding b value
		b := (sum - a) / (a + 1)

		// Check if b is less than m and satisfies the equation
		if b <= m && (sum-a-b) == (a*b) && a < b {
			result = append(result, [2]uint64{a, b}, [2]uint64{b, a})
		}
	}

	// Sort the result if it is not empty
	if len(result) > 0 {
		sort.Slice(result, func(i, j int) bool {
			return result[i][0] < result[j][0]
		})
	} else {
		return nil
	}

	return result
}
