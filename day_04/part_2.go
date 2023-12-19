package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

func main() {
	lines := common.ReadFile("day_04/input_2.txt")

	cardCopies := make(map[int]int)

	for _, line := range lines {
		cardName, numbers, _ := strings.Cut(line, ": ")

		cardNumber, err := strconv.Atoi(strings.TrimSpace(strings.Split(cardName, "Card ")[1]))

		if err != nil {
			panic(err)
		}

		cardCopies[cardNumber]++

		winningNumbers := strings.Split(strings.Split(numbers, " | ")[0], " ")
		myNumbers := strings.Split(strings.Split(numbers, " | ")[1], " ")

		matches := 0
		for _, winningNumber := range winningNumbers {
			if winningNumber == "" {
				continue
			}

			for _, myNumber := range myNumbers {
				if myNumber == "" {
					continue
				}

				if strings.TrimSpace(winningNumber) == strings.TrimSpace(myNumber) {
					matches++
					break
				}
			}
		}

		for i := 1; i <= matches; i++ {
			cardCopies[cardNumber+i] += cardCopies[cardNumber]
		}
	}

	totalCopies := 0
	for _, value := range cardCopies {
		totalCopies += value
	}

	fmt.Printf("Total copies: %v\n", totalCopies)
}
