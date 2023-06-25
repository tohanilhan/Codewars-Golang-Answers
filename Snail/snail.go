package kata

func Snail(snailMap [][]int) []int {

	result := []int{}

	if len(snailMap) == 0 {
		return result
	}

	for len(snailMap) > 0 {

		// Add first row to result
		result = append(result, snailMap[0]...)

		// Remove first row
		snailMap = snailMap[1:]

		// if there is no more rows, break
		if len(snailMap) == 0 {
			break
		}
		// Rotate 90 degrees backwards
		snailMap = rotate(snailMap)
	}
	return result
}

func rotate(snailMap [][]int) [][]int {

	// Create new map
	rotatedMap := make([][]int, len(snailMap[0]))

	// Rotate 90 degrees backwards
	for i := 0; i < len(snailMap); i++ {
		for j := 0; j < len(snailMap[i]); j++ {
			rotatedMap[len(snailMap[i])-j-1] = append(rotatedMap[len(snailMap[i])-j-1], snailMap[i][j])
		}
	}

	return rotatedMap
}
