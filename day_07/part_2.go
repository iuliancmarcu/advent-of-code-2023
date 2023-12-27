package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

var (
	cards = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
	prime = 7919
)

func parseInput() ([]string, map[string]int) {
	hands := make([]string, 0)
	bids := make(map[string]int)

	lines := common.ReadFile("day_07/input_2.txt")

	for _, line := range lines {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bidString := parts[1]

		hands = append(hands, hand)
		bid, err := strconv.Atoi(bidString)

		if err != nil {
			panic(fmt.Sprintf("Cannot parse bid for line: \"%v\"", line))
		}

		bids[hand] = bid
	}

	return hands, bids
}

func cardIndex(card rune) int {
	for i := range cards {
		if cards[i] == card {
			return i
		}
	}

	return -1
}

func maxRight(arr []int) (int, int) {
	index := len(arr) - 1
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[index] {
			index = i
		}
	}

	return index, arr[index]
}

func getHandScore(hand string) float64 {
	handCards := []rune(hand)

	jokers := 0
	counter := make([]int, len(cards))

	for _, card := range handCards {
		if card == 'J' {
			jokers++
			continue
		}

		counter[cardIndex((card))]++
	}

	maxIndex, max := maxRight(counter)
	// fmt.Printf("Hand: %v - Max: %v (%v)\n", hand, val, i)

	counter[maxIndex] += jokers
	max += jokers

	pairCount := 0
	for _, count := range counter {
		if count == 2 {
			pairCount++
		}
	}

	if max == 3 {
		// check for full house
		if pairCount == 1 {
			return 3.5 * float64(prime)
		}
	}

	if pairCount == 2 {
		// double pair
		return 2.5 * float64(prime)
	}

	return float64(max) * float64(prime)
}

func main() {
	hands, bids := parseInput()

	sort.Slice(hands, func(i, j int) bool {
		lhs := hands[i]
		rhs := hands[j]

		lhsScore := getHandScore(lhs)
		rhsScore := getHandScore(rhs)

		if lhsScore == rhsScore {
			// if type (duplicates) is equal, compare place by place
			for place := 0; place < len(lhs); place++ {
				if lhs[place] == rhs[place] {
					continue
				}

				lhsPower := cardIndex(rune(lhs[place])) + 1
				rhsPower := cardIndex(rune(rhs[place])) + 1

				return lhsPower < rhsPower
			}
		}

		return getHandScore(hands[i]) < getHandScore(hands[j])
	})

	// fmt.Printf("Sorted hands:\n")
	// for i := 0; i < len(hands); i++ {
	// 	fmt.Printf("%v %v\n", hands[i], getHandScore(hands[i]))
	// }

	total := uint64(0)

	for i, hand := range hands {
		// fmt.Printf("Hand: %v - Score %v\n", hand, getHandScore(hand))
		total += uint64(i+1) * uint64(bids[hand])
	}

	fmt.Printf("Total score: %v\n", total)
}
