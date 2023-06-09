package main

func IsValidWalk(walk []rune) bool {
	// Enter solution here

	northCount := 0
	southCount := 0
	eastCount := 0
	westCount := 0

	if len(walk) != 10 {
		// if the length of the walk is not 10, return false
		// because it is not a 10 minute walk
		return false
	}

	// loop through the walk and count each direction step
	for _, v := range walk {
		switch v {
		case 'n':
			northCount++
		case 's':
			southCount++
		case 'e':
			eastCount++
		case 'w':
			westCount++
		}
	}

	// if the number of north steps is equal to the number of south steps
	// and the number of east steps is equal to the number of west steps
	// then the walk is valid
	// which means the person ends up where they started
	if northCount == southCount && eastCount == westCount {
		return true
	}

	return false
}
