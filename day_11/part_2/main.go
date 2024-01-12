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

func manhattan(p1, p2 Position) uint {
	return uint(math.Abs(float64(p1.X-p2.X))) + uint(math.Abs(float64(p1.Y-p2.Y)))
}

func hasOnlySpace(line string) bool {
	for _, char := range line {
		if char != '.' {
			return false
		}
	}
	return true
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

func sumExpansionsBetween(a, b int, expansions []int) uint {
	lower := a
	higher := b
	if a > b {
		lower = b
		higher = a
	}

	sum := uint(0)
	for _, exp := range expansions {
		if exp > lower && exp < higher {
			sum += 999999
		}
	}

	return sum
}

func main() {
	universeMap := parseInput()

	// Collect all expanded lines and columns
	lineExpansions := make([]int, 0)
	columnExpansions := make([]int, 0)

	for y, line := range universeMap {
		if hasOnlySpace(line) {
			lineExpansions = append(lineExpansions, y)
		}
	}

	universeMap = transpose(universeMap)

	for x, column := range universeMap {
		if hasOnlySpace(column) {
			columnExpansions = append(columnExpansions, x)
		}
	}

	universeMap = transpose(universeMap)

	// Collect galaxies
	galaxies := collectGalaxies(universeMap)

	// Sum up all paths between pairs of galaxies
	sum := uint(0)
	for _, galaxy1 := range galaxies {
		for _, galaxy2 := range galaxies {
			if galaxy1 != galaxy2 {
				normalDistance := manhattan(galaxy1, galaxy2)

				hExpansionSum := sumExpansionsBetween(galaxy1.X, galaxy2.X, columnExpansions)
				vExpansionSum := sumExpansionsBetween(galaxy1.Y, galaxy2.Y, lineExpansions)

				sum += normalDistance + hExpansionSum + vExpansionSum
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
