package kata

func QueueTime(customers []int, n int) int {
	// If there are no customers or no tills, return 0
	if len(customers) == 0 || n == 0 {
		return 0
	}

	// Create an array to track the time remaining for each till
	tills := make([]int, n)

	// Iterate over each customer in the queue
	for _, customer := range customers {
		// Find the till with the minimum remaining time
		minTime := tills[0]
		minIndex := 0
		for i := 1; i < n; i++ {
			if tills[i] < minTime {
				minTime = tills[i]
				minIndex = i
			}
		}

		// Assign the customer's time to the till with the minimum remaining time
		tills[minIndex] += customer
	}

	// Find the maximum time among all the tills
	maxTime := tills[0]
	for i := 1; i < n; i++ {
		if tills[i] > maxTime {
			maxTime = tills[i]
		}
	}

	return maxTime
}

// func QueueTime(customers []int, n int) int {
// 	// your code here

// 	queueTime := 0

// 	// If there are no customers or no tills, return 0
// 	if len(customers) == 0 || n == 0 {
// 		return 0
// 	}

// 	// if there is only one till, return the sum of the times
// 	if n == 1 {
// 		totalTimeForOneTill := 0
// 		for _, v := range customers {
// 			totalTimeForOneTill += v
// 		}

// 		return totalTimeForOneTill
// 	}

// 	// if there are more than one till, return the highest sum of the times
// 	if n > 1 {
// 		// create a slice to store the times for each till
// 		tillTimes := make([]int, n)

// 		// loop through the customers
// 		for _, v := range customers {
// 			// find the till with the lowest time
// 			lowestTime := tillTimes[0]
// 			lowestTimeIndex := 0
// 			for i, v := range tillTimes {
// 				if v < lowestTime {
// 					lowestTime = v
// 					lowestTimeIndex = i
// 				}
// 			}

// 			// add the customer time to the till with the lowest time
// 			tillTimes[lowestTimeIndex] += v
// 		}

// 		fmt.Println(tillTimes)

// 		// find the highest time
// 		highestTime := tillTimes[0]
// 		for _, v := range tillTimes {
// 			if v > highestTime {
// 				highestTime = v
// 			}
// 		}

// 		queueTime = highestTime
// 	}

// 	return queueTime
// }
