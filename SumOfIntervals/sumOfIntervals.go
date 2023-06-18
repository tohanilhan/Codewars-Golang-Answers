package kata

import "sort"

func SumOfIntervals(intervals [][2]int) int {
	// your code here

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	length := 0
	lowest := intervals[0][0]
	for _, v := range intervals {
		if v[1] >= lowest {
			length += v[1]
			if v[0] > lowest {
				length -= v[0]
			} else {
				length -= lowest
			}
			lowest = v[1]
		}
	}

	return length

}
