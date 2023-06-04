package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	fmt.Println(QueueTime([]int{5, 3, 4}, 1))
	// should return 12
	// because when there is 1 till, the total time is just the sum of the times

	fmt.Println(QueueTime([]int{10, 2, 3, 3}, 2))
	// should return 10
	// because here n=2 and the 2nd, 3rd, and 4th people in the
	// queue finish before the 1st person has finished.

	fmt.Println(QueueTime([]int{2, 3, 10}, 2))
	// should return 12

	if QueueTime([]int{}, 1) == 0 {
		color.Green("PASS")
	} else {
		color.Red("Should return zero for empty queue")
	}
	if QueueTime([]int{1, 2, 3, 4}, 1) == (10) {
		color.Green("PASS")
	} else {
		color.Red("Should return the sum of the times for a single till which is 10 for this test")
	}
	if QueueTime([]int{2, 2, 3, 3, 4, 4}, 2) == (9) {
		color.Green("PASS")
	} else {
		color.Red("Should return the highest sum for a multiple till which is 9 for this test")
	}
	if QueueTime([]int{1, 2, 3, 4, 5}, 100) == (5) {
		color.Green("PASS")
	} else {
		color.Red("Should return the highest sum for a multiple till which is 5 for this test")
	}

}

func QueueTime(customers []int, n int) int {
	// your code here

	queueTime := 0

	if len(customers) == 0 {
		return 0
	}

	// if there is only one till, total time is the sum of the times
	if n == 1 && len(customers) > 0 {
		for _, v := range customers {
			queueTime += v
		}
	}

	// if customer length is less than or equal to the number of tills, the total time is the max of the times
	if len(customers) <= n {
		for _, v := range customers {
			if v > queueTime {
				queueTime = v
			}
		}
	}

	// if there are more than one till, we need to find the shortest queue

	// create

	return queueTime
}
