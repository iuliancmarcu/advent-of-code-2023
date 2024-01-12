package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

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

func unfoldRows(rows []SpringRow) []SpringRow {
	newRows := make([]SpringRow, 0)

	for _, row := range rows {
		springsArr := make([]string, 0)
		for i := 0; i < 5; i++ {
			springsArr = append(springsArr, row.Springs)
		}

		springs := strings.Join(springsArr, "?")

		damaged := make([]int, 0)
		for i := 0; i < 5; i++ {
			damaged = append(damaged, row.Damaged...)
		}

		newRows = append(newRows, SpringRow{
			Springs: springs,
			Damaged: damaged,
		})
	}

	return newRows
}

func hashCall(row string, toFit []int) string {
	str := row + "_"
	for _, fit := range toFit {
		str += strconv.Itoa(fit) + ","
	}
	return str
}

func numOfFittings(row string, toFit []int, sumOfFittings int, callCache map[string]uint) uint {
	cacheHash := hashCall(row, toFit)

	cache, ok := callCache[cacheHash]
	if ok {
		return cache
	}

	// fmt.Printf("Checking fittings in \"%v\" for %v\n", row, toFit)
	if len(row) == 0 && len(toFit) == 0 {
		// fmt.Printf("\tFitting found!\n")
		callCache[cacheHash] = 1
		return 1
	}

	if len(row) == 0 || len(row) < sumOfFittings+len(toFit)-1 {
		// fmt.Printf("\tFitting not found!\n")
		callCache[cacheHash] = 0
		return 0
	}

	switch row[0] {
	case '.':
		// fmt.Printf("+Skipping space\n")
		result := numOfFittings(row[1:], toFit, sumOfFittings, callCache)
		callCache[cacheHash] = result
		return result
	case '#':
		if len(toFit) < 1 || len(row) < toFit[0] {
			callCache[cacheHash] = 0
			return 0
		}

		size := toFit[0]

		// fmt.Printf("+Checking fit: \"%v\" - %v\n", row, size)
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
			callCache[cacheHash] = 0
			return 0
		}

		if shouldFitWithSpace {
			result := numOfFittings(row[size+1:], toFit[1:], sumOfFittings-toFit[0], callCache)
			callCache[cacheHash] = result
			return result
		}

		result := numOfFittings(row[size:], toFit[1:], sumOfFittings-toFit[0], callCache)
		callCache[cacheHash] = result
		return result
	case '?':
		// fmt.Printf("+Branching\n")
		result := numOfFittings(row[1:], toFit, sumOfFittings, callCache) +
			numOfFittings("#"+row[1:], toFit, sumOfFittings, callCache)
		callCache[cacheHash] = result
		return result
	}

	callCache[hashCall(row, toFit)] = 0
	return 0
}

func main() {
	rows := parseInput()

	rows = unfoldRows(rows)

	fittingsChannel := make(chan uint, len(rows))
	var wg sync.WaitGroup

	// For each row, get total options to fit damaged
	for _, row := range rows {
		wg.Add(1)
		go func(row SpringRow) {
			defer wg.Done()

			totalDamaged := 0
			for _, dmg := range row.Damaged {
				totalDamaged += dmg
			}

			callCache := make(map[string]uint)
			rowTotal := numOfFittings(row.Springs, row.Damaged, totalDamaged, callCache)
			// fmt.Printf("Total for fittings %v on row \"%v\": %v\n\n", row.Damaged, row.Springs, rowTotal)

			fittingsChannel <- rowTotal
		}(row)
	}

	wg.Wait()
	close(fittingsChannel)

	// Read from sampleChan and put into a slice
	total := uint(0)
	for s := range fittingsChannel {
		total += s
	}

	fmt.Printf("Total sum of fitting options: %v\n", total)
}
