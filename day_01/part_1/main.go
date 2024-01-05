package main

import (
	"fmt"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

func main() {
	lines := common.ReadFile("day_01/input.txt")

	total := 0

	for _, line := range lines {
		var first byte
		var last byte

		for i, c := range line {
			if c >= '0' && c <= '9' {
				first = line[i]
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = line[i]
				break
			}
		}

		// concatenate the 2 digits
		digits := string(first) + string(last)

		// convert to int
		var result int
		fmt.Sscanf(digits, "%d", &result)

		total += result
	}

	fmt.Printf("Total is: %v\n", total)
}
