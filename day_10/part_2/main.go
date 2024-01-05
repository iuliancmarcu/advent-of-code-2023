package main

import (
	"fmt"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

type Visit struct {
	Y int
	X int
}

func parseInput() ([]string, Visit) {
	lines := common.ReadFile("day_10/input.txt")

	var start Visit

	// find the start position
	foundStart := false
	for y := range lines {
		for x, char := range lines[y] {
			if char == 'S' {
				start = Visit{Y: y, X: x}
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
	return char == '|' || char == 'L' || char == 'J' || char == 'S'
}

func hasDown(char byte) bool {
	return char == '|' || char == 'F' || char == '7' || char == 'S'
}

func hasLeft(char byte) bool {
	return char == '-' || char == 'J' || char == '7' || char == 'S'
}

func hasRight(char byte) bool {
	return char == '-' || char == 'L' || char == 'F' || char == 'S'
}

func main() {
	lines, start := parseInput()

	walls := buildWalls(&lines, start)

	// fmt.Printf("Initial:\n")
	// print(&walls)
	// fmt.Printf("\n")

	expandedWalls := expandWalls(&walls)

	// fmt.Printf("Expanded:\n")
	// print(&expandedWalls)
	// fmt.Printf("\n")

	floodFill(&expandedWalls, 15)

	// fmt.Printf("Flooded:\n")
	// print(&expandedWalls)
	// fmt.Printf("\n")

	compactedWalls := compactWalls(&expandedWalls)

	fmt.Printf("Compacted:\n")
	print(&compactedWalls)
	fmt.Printf("\n")

	fmt.Printf("Total inside cells: %v\n", countZeros(&compactedWalls))
}

func floodFill(walls *[][]int, value int) {
	inBounds := func(y int, x int) bool {
		return y >= 0 && y < len((*walls)) && x >= 0 && x < len((*walls)[y])
	}

	stack := make([]Visit, 0)
	stack = append(stack, Visit{Y: 0, X: 0})
	for len(stack) > 0 {
		curr := stack[0]

		if inBounds(curr.Y, curr.X) && (*walls)[curr.Y][curr.X] == 0 {
			(*walls)[curr.Y][curr.X] = value

			stack = append(stack, Visit{Y: curr.Y - 1, X: curr.X})
			stack = append(stack, Visit{Y: curr.Y + 1, X: curr.X})
			stack = append(stack, Visit{Y: curr.Y, X: curr.X - 1})
			stack = append(stack, Visit{Y: curr.Y, X: curr.X + 1})
		}

		stack = stack[1:]
	}
}

func buildWalls(lines *[]string, start Visit) [][]int {
	inBounds := func(y int, x int) bool {
		return y >= 0 && y < len((*lines)) && x >= 0 && x < len((*lines)[y])
	}

	walls := make([][]int, len((*lines)))
	for y := 0; y < len((*lines)); y++ {
		walls[y] = make([]int, len((*lines)[y]))

		for x := 0; x < len(walls[y]); x++ {
			walls[y][x] = 0
		}
	}

	stack := make([]Visit, 0)
	stack = append(stack, start)

	for len(stack) > 0 {
		visit := stack[0]

		if walls[visit.Y][visit.X] == 0 {
			wall := 0

			if hasUp((*lines)[visit.Y][visit.X]) && inBounds(visit.Y-1, visit.X) && hasDown((*lines)[visit.Y-1][visit.X]) {
				// connected to up
				wall |= 1 << 0
				stack = append(stack, Visit{Y: visit.Y - 1, X: visit.X})
			}

			if hasDown((*lines)[visit.Y][visit.X]) && inBounds(visit.Y+1, visit.X) && hasUp((*lines)[visit.Y+1][visit.X]) {
				// connected to down
				wall |= 1 << 1
				stack = append(stack, Visit{Y: visit.Y + 1, X: visit.X})
			}

			if hasLeft((*lines)[visit.Y][visit.X]) && inBounds(visit.Y, visit.X-1) && hasRight((*lines)[visit.Y][visit.X-1]) {
				// connected to left
				wall |= 1 << 2
				stack = append(stack, Visit{Y: visit.Y, X: visit.X - 1})
			}

			if hasRight((*lines)[visit.Y][visit.X]) && inBounds(visit.Y, visit.X+1) && hasLeft((*lines)[visit.Y][visit.X+1]) {
				// connected to right
				wall |= 1 << 3
				stack = append(stack, Visit{Y: visit.Y, X: visit.X + 1})
			}

			walls[visit.Y][visit.X] = wall
		}

		stack = stack[1:]
	}

	return walls
}

func projectWallTo3x3(wall int) [][]int {
	return [][]int{
		{0, wall & (1 << 0), 0},
		{wall & (1 << 2), wall, wall & (1 << 3)},
		{0, wall & (1 << 1), 0},
	}
}

func expandWalls(walls *[][]int) [][]int {
	newWalls := make([][]int, len((*walls))*3)
	for y := range newWalls {
		newWalls[y] = make([]int, len((*walls)[0])*3)
	}

	for y := range *walls {
		for x := range (*walls)[y] {
			projectedWall := projectWallTo3x3((*walls)[y][x])

			for py := range projectedWall {
				for px := range projectedWall[py] {
					newWalls[y*3+py][x*3+px] = projectedWall[py][px]
				}
			}
		}
	}

	return newWalls
}

func wallFrom3x3Kern(kern *[][]int) int {
	for y := range *kern {
		for x := range (*kern)[y] {
			if (*kern)[y][x] != 0 {
				return 1
			}
		}
	}

	return 0
}

func compactWalls(walls *[][]int) [][]int {
	newWalls := make([][]int, len(*walls)/3)
	for y := range newWalls {
		newWalls[y] = make([]int, len((*walls)[0])/3)
	}

	wallAt := func(y, x int) int {
		return (*walls)[y][x]
	}

	kern3x3At := func(y, x int) *[][]int {
		return &[][]int{
			{wallAt(y, x), wallAt(y, x+1), wallAt(y, x+2)},
			{wallAt(y+1, x), wallAt(y+1, x+1), wallAt(y+1, x+2)},
			{wallAt(y+2, x), wallAt(y+2, x+1), wallAt(y+2, x+2)},
		}
	}

	for y := range newWalls {
		for x := range newWalls[y] {
			newWalls[y][x] = wallFrom3x3Kern(kern3x3At(y*3, x*3))
		}
	}

	return newWalls
}

func print(array *[][]int) {
	for y := 0; y < len((*array)); y++ {
		for x := 0; x < len((*array)[y]); x++ {
			fmt.Printf("%x", (*array)[y][x])
		}
		fmt.Print("\n")
	}
}

func countZeros(array *[][]int) int {
	count := 0
	for y := range *array {
		for x := range (*array)[y] {
			if (*array)[y][x] == 0 {
				count++
			}
		}
	}
	return count
}
