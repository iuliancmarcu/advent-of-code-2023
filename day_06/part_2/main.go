package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

func parseInput2() (time int, distance int) {
	lines := common.ReadFile("day_06/input.txt")

	timeString := strings.Join(strings.Fields(strings.Split(lines[0], "Time:")[1]), "")
	distanceString := strings.Join(strings.Fields(strings.Split(lines[1], "Distance:")[1]), "")

	timeInt, _ := strconv.Atoi(timeString)
	distanceInt, _ := strconv.Atoi(distanceString)

	time = timeInt
	distance = distanceInt

	return time, distance
}

func main() {
	time, recordDistance := parseInput2()

	waysToBeatTheRace := 0
	for t := 0; t <= time; t++ {
		totalDistance := (time - t) * t

		if totalDistance > recordDistance {
			waysToBeatTheRace++
		}
	}

	fmt.Printf("Ways to beat the race: %v\n", waysToBeatTheRace)
}
