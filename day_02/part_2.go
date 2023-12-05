package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputFile = "day_02/input_2.txt"
)

func main() {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error reading file \"%v\"\n", inputFile)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		var gameId int
		fmt.Sscanf(line, "Game %d:", &gameId)

		restOfLine, found := strings.CutPrefix(line, fmt.Sprintf("Game %d: ", gameId))

		if !found {
			fmt.Printf("Error parsing line \"%v\"\n", line)
			return
		}

		// check if game is valid
		extractions := strings.Split(restOfLine, "; ")

		colors := make(map[string]int)

		for _, ext := range extractions {
			cubes := strings.Split(ext, ", ")

			for _, cube := range cubes {
				var color string
				var value int
				fmt.Sscanf(cube, "%d %s", &value, &color)

				if colors[color] < value {
					colors[color] = value
				}
			}
		}

		total += colors["red"] * colors["green"] * colors["blue"]
	}

	fmt.Printf("Total is: %v\n", total)
}
