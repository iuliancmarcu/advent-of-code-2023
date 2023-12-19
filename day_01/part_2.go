package main

import (
	"fmt"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

var digits = [...]string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

func main() {
	lines := common.ReadFile("day_01/input_2.txt")

	total := 0

	for _, line := range lines {
		var first byte
		var last byte

		for i := range line {
			if first != 0 {
				break
			}

			for j, digit := range digits {
				if i+len(digit) > len(line) {
					continue
				}

				substr := line[i : i+len(digit)]

				if substr == digit {
					first = byte(j%10) + '0'
					break
				}
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if last != 0 {
				break
			}

			for j, digit := range digits {
				if i+len(digit) > len(line) {
					continue
				}

				substr := line[i : i+len(digit)]

				if substr == digit {
					last = byte(j%10) + '0'
					break
				}
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
