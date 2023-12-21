package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

func parseInput() (times []int, distances []int) {
	lines := common.ReadFile("day_06/input_1.txt")

	timesString := strings.Fields(strings.Split(lines[0], "Time:")[1])
	distancesString := strings.Fields(strings.Split(lines[1], "Distance:")[1])

	for i := range timesString {
		time, _ := strconv.Atoi(timesString[i])
		distance, _ := strconv.Atoi(distancesString[i])

		times = append(times, time)
		distances = append(distances, distance)
	}

	return times, distances
}

func main() {
	times, distances := parseInput()

	total := 1

	for i, raceTime := range times {
		recordDistance := distances[i]

		resultsOverRecord := 0
		distanceForTimeHeld := make([]int, raceTime+1)
		for timeHeld := 0; timeHeld <= raceTime; timeHeld++ {
			distanceForTimeHeld[timeHeld] = (raceTime - timeHeld) * timeHeld

			if distanceForTimeHeld[timeHeld] > recordDistance {
				resultsOverRecord++
			}
		}

		total *= resultsOverRecord
	}

	fmt.Printf("Total: %v\n", total)
}
