package kata

import "sort"

func SumOfIntervals(intervals [][2]int) int {
	// your code here

	lenght := 0

	mergedIntervals := merge(intervals)

	for _, interval := range mergedIntervals {
		lenght += interval[1] - interval[0]
	}

	return lenght
}

type intervalsArray [][2]int

func (intA intervalsArray) Len() int {
	return len(intA)
}

func (intA intervalsArray) Swap(i, j int) {
	intA[i], intA[j] = intA[j], intA[i]
}

func (intA intervalsArray) Less(i, j int) bool {
	return intA[i][0] < intA[j][0]
}

func merge(intervals [][2]int) [][2]int {

	intA := intervalsArray(intervals)

	sort.Sort(intA)

	intervalsSorted := [][2]int(intA)
	//fmt.Println(intervalsSorted)

	var output [][2]int
	currentIntervalStart := intervalsSorted[0][0]
	currentIntervalEnd := intervalsSorted[0][1]
	for j := 1; j < len(intervalsSorted); j++ {
		if currentIntervalEnd >= intervalsSorted[j][0] {
			if intervalsSorted[j][1] > currentIntervalEnd {
				currentIntervalEnd = intervalsSorted[j][1]
			}
		} else {
			output = append(output, [2]int{currentIntervalStart, currentIntervalEnd})
			currentIntervalStart = intervalsSorted[j][0]
			currentIntervalEnd = intervalsSorted[j][1]
		}
	}
	output = append(output, [2]int{currentIntervalStart, currentIntervalEnd})
	return output

}
