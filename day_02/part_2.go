package main

import (
	"fmt"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

func main() {
	lines := common.ReadFile("day_02/input_2.txt")

	total := 0

	for _, line := range lines {
		var gameId int
		fmt.Sscanf(line, "Game %d:", &gameId)

		restOfLine, found := strings.CutPrefix(line, fmt.Sprintf("Game %d: ", gameId))

		if !found {
			fmt.Printf("Error parsing line \"%v\"\n", line)
			return
		}

		// check if game is valid
		extractions := strings.Split(restOfLine, "; ")

		colors := make(map[string]int)

		for _, ext := range extractions {
			cubes := strings.Split(ext, ", ")

			for _, cube := range cubes {
				var color string
				var value int
				fmt.Sscanf(cube, "%d %s", &value, &color)

				if colors[color] < value {
					colors[color] = value
				}
			}
		}

		total += colors["red"] * colors["green"] * colors["blue"]
	}

	fmt.Printf("Total is: %v\n", total)
}
