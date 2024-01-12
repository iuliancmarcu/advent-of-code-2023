package main

import (
	"fmt"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

func parseInput() []string {
	lines := common.ReadFile("day_00/input.txt")
	return lines
}

func main() {
	lines := parseInput()

	fmt.Printf("Number of lines: %v\n", len(lines))
}
