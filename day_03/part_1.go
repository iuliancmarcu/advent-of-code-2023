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
	lines := common.ReadFile("day_03/input_1.txt")

	total := 0

	possibleNumbers := make([]PossibleNumber, 0)
	symbols := make([]Symbol, 0)

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
					// this is a symbol
					symbols = append(symbols, Symbol{
						Value: string(char),
						Row:   l,
						Col:   i,
					})

					fmt.Printf("Symbol: %v %v %v\n", string(char), l, i)
				}
			}
		}
	}

	// now we have all the possible numbers and symbols
	// we need to figure out which numbers are valid (adjacent to symbols)
	// and which are not
	for _, number := range possibleNumbers {
		valid := false

		for _, symbol := range symbols {
			if symbol.Row == number.Row {
				// same row
				if symbol.Col == number.Start-1 || symbol.Col == number.End+1 {
					valid = true
					break
				}
			} else if symbol.Row == number.Row-1 || symbol.Row == number.Row+1 {
				// adjacent row
				if symbol.Col >= number.Start-1 && symbol.Col <= number.End+1 {
					valid = true
					break
				}
			}
		}

		if valid {
			fmt.Printf("Valid number: %v %v %v\n", number.Value, number.Start, number.End)
			total += number.Value
		}
	}

	fmt.Printf("Total is: %v\n", total)
}
