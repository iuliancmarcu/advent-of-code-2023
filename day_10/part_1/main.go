package main

import (
	"fmt"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

type Visit struct {
	Step int
	X    int
	Y    int
}

func parseInput() ([]string, Visit) {
	lines := common.ReadFile("day_10/input.txt")

	var start Visit

	// find the start position
	foundStart := false
	for y := range lines {
		for x, char := range lines[y] {
			if char == 'S' {
				start = Visit{Step: 1, X: x, Y: y}
				foundStart = true
			}

			if foundStart {
				break
			}
		}

		if foundStart {
			break
		}
	}

	return lines, start
}

func hasUp(char byte) bool {
	return char == '|' || char == 'L' || char == 'J'
}

func hasDown(char byte) bool {
	return char == '|' || char == 'F' || char == '7'
}

func hasLeft(char byte) bool {
	return char == '-' || char == 'J' || char == '7'
}

func hasRight(char byte) bool {
	return char == '-' || char == 'L' || char == 'F'
}

func main() {
	lines, start := parseInput()

	inBounds := func(y int, x int) bool {
		return y >= 0 && y < len(lines) && x >= 0 && x < len(lines[0])
	}

	steps := make([][]int, len(lines))
	for y := 0; y < len(lines); y++ {
		steps[y] = make([]int, len(lines[y]))

		for x := 0; x < len(lines); x++ {
			steps[y][x] = 0
		}
	}

	stack := make([]Visit, 0)
	stack = append(stack, start)

	for len(stack) > 0 {
		visit := stack[0]

		if steps[visit.Y][visit.X] == 0 {
			steps[visit.Y][visit.X] = visit.Step

			if inBounds(visit.Y-1, visit.X) && hasDown(lines[visit.Y-1][visit.X]) {
				stack = append(stack, Visit{Step: visit.Step + 1, Y: visit.Y - 1, X: visit.X})
			}

			if inBounds(visit.Y+1, visit.X) && hasUp(lines[visit.Y+1][visit.X]) {
				stack = append(stack, Visit{Step: visit.Step + 1, Y: visit.Y + 1, X: visit.X})
			}

			if inBounds(visit.Y, visit.X-1) && hasRight(lines[visit.Y][visit.X-1]) {
				stack = append(stack, Visit{Step: visit.Step + 1, Y: visit.Y, X: visit.X - 1})
			}

			if inBounds(visit.Y, visit.X+1) && hasLeft(lines[visit.Y][visit.X+1]) {
				stack = append(stack, Visit{Step: visit.Step + 1, Y: visit.Y, X: visit.X + 1})
			}
		}

		stack = stack[1:]
	}

	// find max
	max := 0
	for y := 0; y < len(steps); y++ {
		for x := 0; x < len(steps[y]); x++ {
			if steps[y][x] > max {
				max = steps[y][x]
			}
		}
	}

	// -1 because we start from 1 to avoid issues with start position
	fmt.Printf("Max steps is: %v\n", max-1)
}
