package main

import "fmt"

func findBusiestPeriod(data [][]int) int {
	runningCost := 0
	busiestTime := 0
	maxCost := 0

	for time := 0; time < len(data); {
		currentTime := data[time][0]
		currentNoPeople := data[time][1]
		status := data[time][2]

		if status == 1 {
			runningCost += currentNoPeople
		} else {
			runningCost -= currentNoPeople
		}
		time += 1

		if time < len(data) && currentTime == data[time][0] {
			continue
		}

		if runningCost > maxCost {
			maxCost = runningCost
			busiestTime = currentTime
		}
	}
	return busiestTime
}

func main() {

	data := [][]int{
		{1487799425, 14, 1},
		{1487799425, 4, 1},
		{1487799425, 2, 1},
		{1487800378, 10, 1},
		{1487800478, 18, 1},
		{1487901013, 1, 1},
		{1487901211, 7, 1},
		{1487901211, 7, 1},
	}
	// Expected Output : 1487901211
	fmt.Println(findBusiestPeriod(data))
}
