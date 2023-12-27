package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

func parseInput() [][]int {
	lines := common.ReadFile("day_09/input_1.txt")

	sequences := make([][]int, len(lines))

	for i, line := range lines {
		lineVals := strings.Split(line, " ")

		sequence := make([]int, 0)
		for _, valStr := range lineVals {
			val, _ := strconv.Atoi(valStr)
			sequence = append(sequence, val)
		}

		sequences[i] = sequence
	}

	return sequences
}

func main() {
	sequences := parseInput()

	total := 0
	for _, sequence := range sequences {
		// fmt.Printf("Processing sequence: %v\n", sequence)

		seq := sequence
		prediction := seq[len(seq)-1]

		allZeros := false
		for !allZeros {
			allZeros = true
			for i := 1; i < len(seq); i++ {
				seq[i-1] = seq[i] - seq[i-1]

				if seq[i-1] != 0 {
					allZeros = false
				}
			}

			seq = seq[:len(seq)-1]
			prediction += seq[len(seq)-1]
		}

		// fmt.Printf("  prediction: %v\n", prediction)

		total += prediction
	}

	fmt.Printf("Total sum of predictions: %v\n", total)
}
