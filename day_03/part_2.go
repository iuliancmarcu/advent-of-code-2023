package main

import (
	"fmt"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

type PossibleNumber struct {
	Value int
	Row   int
	Start int
	End   int
}

type Symbol struct {
	Value string
	Row   int
	Col   int
}

func main() {
	lines := common.ReadFile("day_03/input_2.txt")

	total := 0

	possibleNumbers := make([]PossibleNumber, 0)
	possibleGears := make([]Symbol, 0)

	for l, line := range lines {
		for i := 0; i < len(line); i++ {
			char := line[i]

			if char != '.' {
				// can be a symbol or a number
				if char >= '0' && char <= '9' {
					// scan until another '.' is found or the end of the line

					j := i
					for j < len(line) && line[j] >= '0' && line[j] <= '9' {
						j++
					}

					start := i
					end := j - 1
					i = end

					// get actual value
					value := 0
					fmt.Sscanf(line[start:end+1], "%d", &value)

					possibleNumbers = append(possibleNumbers, PossibleNumber{
						Value: value,
						Row:   l,
						Start: start,
						End:   end,
					})
				} else {
					if char != '*' {
						continue
					}

					// this is a possible gear
					possibleGears = append(possibleGears, Symbol{
						Value: string(char),
						Row:   l,
						Col:   i,
					})

					fmt.Printf("Possible Gear: %v %v %v\n", string(char), l, i)
				}
			}
		}
	}

	// now we have all the possible numbers and possibleGears
	// we need to figure out which ones are the gears and which numbers are adjacent to them

	for _, symbol := range possibleGears {
		adjacentNumbers := []PossibleNumber{}

		for _, number := range possibleNumbers {
			if symbol.Row == number.Row {
				// same row
				if symbol.Col == number.Start-1 || symbol.Col == number.End+1 {
					adjacentNumbers = append(adjacentNumbers, number)
				}
			} else if symbol.Row == number.Row-1 || symbol.Row == number.Row+1 {
				// adjacent row
				if symbol.Col >= number.Start-1 && symbol.Col <= number.End+1 {
					adjacentNumbers = append(adjacentNumbers, number)
				}
			}
		}

		valid := len(adjacentNumbers) == 2

		if valid {
			total += adjacentNumbers[0].Value * adjacentNumbers[1].Value
		}
	}

	fmt.Printf("Total is: %v\n", total)
}
