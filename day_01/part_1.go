package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	inputFile = "day_01/input_1.txt"
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

		var first byte
		var last byte

		for i, c := range line {
			if c >= '0' && c <= '9' {
				first = line[i]
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i] >= '0' && line[i] <= '9' {
				last = line[i]
				break
			}
		}

		// concatenate the 2 digits
		digits := string(first) + string(last)

		// convert to int
		var result int
		fmt.Sscanf(digits, "%d", &result)

		total += result
	}

	fmt.Printf("Total is: %v\n", total)
}
