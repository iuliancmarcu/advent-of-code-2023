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
	lines := common.ReadFile("day_08/input_1.txt")

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

func main() {
	instructions, links := parseInput()

	steps := 0
	current := "AAA"
	for current != "ZZZ" {
		inst := instructions[steps%len(instructions)]
		link := links[current]

		steps++
		if inst == 'L' {
			current = link.Left
		} else {
			current = link.Right
		}
	}

	fmt.Printf("Total steps: %v\n", steps)
}
