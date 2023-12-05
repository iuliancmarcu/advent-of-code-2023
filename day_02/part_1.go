package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	inputFile = "day_02/input_1.txt"
	maxRed    = 12
	maxGreen  = 13
	maxBlue   = 14
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

		if colors["red"] > maxRed || colors["green"] > maxGreen || colors["blue"] > maxBlue {
			continue
		}

		total += gameId
	}

	fmt.Printf("Total is: %v\n", total)
}
