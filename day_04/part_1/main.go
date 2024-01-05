package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

func main() {
	lines := common.ReadFile("day_04/input.txt")

	totalPoints := 0.0

	for i, line := range lines {
		numbers := strings.Split(line, ": ")[1]

		winningNumbers := strings.Split(strings.Split(numbers, " | ")[0], " ")
		myNumbers := strings.Split(strings.Split(numbers, " | ")[1], " ")

		matches := 0
		for _, winningNumber := range winningNumbers {
			if winningNumber == "" {
				continue
			}

			for _, myNumber := range myNumbers {
				if myNumber == "" {
					continue
				}

				if strings.TrimSpace(winningNumber) == strings.TrimSpace(myNumber) {
					matches++
					break
				}
			}
		}

		if matches > 0 {
			fmt.Printf("Card %v: %v\n", i+1, matches)
			totalPoints += math.Pow(2, float64(matches-1))
		}
	}

	fmt.Printf("Total points: %v\n", int(totalPoints))
}
