package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

type SpringRow struct {
	Springs string
	Damaged []int
}

func parseInput() []SpringRow {
	lines := common.ReadFile("day_12/input.txt")

	rows := make([]SpringRow, 0)
	for _, line := range lines {
		springs, damagedStr, ok := strings.Cut(line, " ")
		if !ok {
			fmt.Printf("Failed to decode line \"%v\"", line)
			panic(nil)
		}

		damaged := make([]int, 0)
		for _, dmg := range strings.Split(damagedStr, ",") {
			converted, err := strconv.Atoi(dmg)
			if err != nil {
				fmt.Printf("Failed to decode damaged \"%v\" from line \"%v\"", damagedStr, line)
				panic(err)
			}

			damaged = append(damaged, converted)
		}

		rows = append(rows, SpringRow{Springs: springs, Damaged: damaged})
	}

	return rows
}

func numOfFittings(row string, toFit []int, sumOfFittings int) int {
	// fmt.Printf("Checking fittings in \"%v\" for %v\n", row, toFit)
	if len(row) == 0 && len(toFit) == 0 {
		// fmt.Printf("\tFitting found!\n")
		return 1
	}

	if len(row) == 0 || len(row) < sumOfFittings+len(toFit)-1 {
		// fmt.Printf("\tFitting not found!\n")
		return 0
	}

	if row[0] == '.' {
		// fmt.Printf("+Skipping space\n")
		// skip "working" springs
		return numOfFittings(row[1:], toFit, sumOfFittings)
	}

	if row[0] == '#' {
		if len(toFit) < 1 || len(row) < toFit[0] {
			return 0
		}

		size := toFit[0]

		// fmt.Printf("+Checking fit: \"%v\" - %v\n", row, size)
		// check if we can fit
		shouldFitWithSpace := false
		if len(row) > size {
			shouldFitWithSpace = true
		}

		canFit := true
		for i := 0; i < size; i++ {
			if row[i] == '.' {
				canFit = false
			}
		}

		if shouldFitWithSpace {
			if !(row[size] == '.' || row[size] == '?') {
				canFit = false
			}
		}

		if !canFit {
			return 0
		}

		if shouldFitWithSpace {
			return numOfFittings(row[size+1:], toFit[1:], sumOfFittings-toFit[0])
		}

		return numOfFittings(row[size:], toFit[1:], sumOfFittings-toFit[0])
	}

	if row[0] == '?' {
		// fmt.Printf("+Branching\n")
		return numOfFittings(row[1:], toFit, sumOfFittings) +
			numOfFittings("#"+row[1:], toFit, sumOfFittings)
	}

	return 0
}

func main() {
	rows := parseInput()

	// For each row, get total options to fit damaged
	total := 0
	for _, row := range rows {
		totalDamaged := 0
		for _, dmg := range row.Damaged {
			totalDamaged += dmg
		}

		rowTotal := numOfFittings(row.Springs, row.Damaged, totalDamaged)
		fmt.Printf("Total for fittings %v on row \"%v\": %v\n\n", row.Damaged, row.Springs, rowTotal)
		total += rowTotal
	}

	fmt.Printf("Total sum of fitting options: %v\n", total)
}
