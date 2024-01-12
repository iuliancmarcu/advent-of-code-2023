package main

import (
	"fmt"
	"math"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

type Position struct {
	X int
	Y int
}

func parseInput() []string {
	lines := common.ReadFile("day_11/input.txt")
	return lines
}

func transpose(lines []string) []string {
	columns := make([]string, len(lines[0]))

	for y := 0; y < len(lines); y++ {
		line := lines[y]

		for x := 0; x < len(line); x++ {
			columns[x] += string(line[x])
		}
	}

	return columns
}

func manhattan(p1, p2 Position) int {
	return int(math.Abs(float64(p1.X-p2.X))) + int(math.Abs(float64(p1.Y-p2.Y)))
}

func hasOnlySpace(line string) bool {
	for _, char := range line {
		if char != '.' {
			return false
		}
	}
	return true
}

func duplicate(array []string, index int) []string {
	left := array[:index]
	right := array[index:]

	duplicated := make([]string, 0)
	duplicated = append(duplicated, left...)
	duplicated = append(duplicated, right[0])
	duplicated = append(duplicated, right...)
	return duplicated
}

func collectGalaxies(lines []string) []Position {
	columns := transpose(lines)

	galaxies := make([]Position, 0)
	for y, line := range lines {
		for x := range columns {
			char := line[x]
			if char == '#' {
				galaxies = append(galaxies, Position{X: x, Y: y})
			}
		}
	}
	return galaxies
}

func main() {
	universeMap := parseInput()

	fmt.Printf("Lines:\n")
	debugList(universeMap)
	fmt.Printf("\n")

	fmt.Printf("Columns:\n")
	debugList(transpose(universeMap))
	fmt.Printf("\n")

	duplications := 0
	for y, line := range universeMap {
		if hasOnlySpace(line) {
			// duplicate line y
			fmt.Printf("Duplicating line %v\n", y)
			universeMap = duplicate(universeMap, y+duplications)
			duplications++
		}
	}

	universeMap = transpose(universeMap)

	duplications = 0
	for x, column := range universeMap {
		if hasOnlySpace(column) {
			// duplicate column x
			fmt.Printf("Duplicating column %v\n", x)
			universeMap = duplicate(universeMap, x+duplications)
			duplications++
		}
	}

	universeMap = transpose(universeMap)

	fmt.Printf("POST DUPLICATION:\n")

	fmt.Printf("Lines:\n")
	debugList(universeMap)
	fmt.Printf("\n")

	fmt.Printf("Columns:\n")
	debugList(transpose(universeMap))
	fmt.Printf("\n")

	galaxies := collectGalaxies(universeMap)

	sum := 0
	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			if galaxy1 != galaxy2 {
				sum += manhattan(galaxy1, galaxy2)
			}
		}
	}

	fmt.Printf("Sum of shortest paths between each pair of galaxies: %v\n", sum/2)
}

func debugList(list []string) {
	for _, line := range list {
		fmt.Printf("%v\n", line)
	}
}
