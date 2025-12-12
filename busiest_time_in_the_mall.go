package main

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
