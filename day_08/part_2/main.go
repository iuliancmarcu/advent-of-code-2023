package main

import (
	"fmt"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

type Link struct {
	Origin string
	Left   string
	Right  string
}

func parseInput() (instructions []rune, links map[string]Link) {
	lines := common.ReadFile("day_08/input.txt")

	// first line is the list of instructions
	instructions = []rune(lines[0])

	links = make(map[string]Link)
	for i := 2; i < len(lines); i++ {
		origin, targets, _ := strings.Cut(lines[i], " = ")

		targets = strings.Replace(targets, "(", "", -1)
		targets = strings.Replace(targets, ")", "", -1)

		left, right, _ := strings.Cut(targets, ", ")
		links[origin] = Link{Origin: origin, Left: left, Right: right}
	}

	return instructions, links
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func endsInZ(str string) bool {
	return str[len(str)-1] == 'Z'
}

func allEndInZ(list []string) bool {
	for _, str := range list {
		if str[len(str)-1] != 'Z' {
			return false
		}
	}
	return true
}

func main() {
	instructions, links := parseInput()

	current := make([]string, 0)

	for _, link := range links {
		if link.Origin[len(link.Origin)-1] == 'A' {
			current = append(current, link.Origin)
		}
	}

	// fmt.Printf("Starting from: %v\n", current)
	stepsPerOrigin := make([]int, len(current))

	// for each origin, move through instructions until reaching a node ending in Z
	for i, origin := range current {
		steps := 0
		node := origin
		for !endsInZ(node) {
			inst := instructions[steps%len(instructions)]
			steps++

			link := links[node]
			if inst == 'L' {
				node = link.Left
			} else {
				node = link.Right
			}
		}
		stepsPerOrigin[i] = steps
	}

	// total steps is the LCM of steps for each origin node
	fmt.Printf("Total steps: %v\n", LCM(stepsPerOrigin...))
}
