package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

// seed -> soil -> fertilizer -> water -> light -> temperature -> humidity -> location

var (
	categories = []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
)

func main() {
	lines := common.ReadFile("day_05/input_1.txt")

	var seeds []uint64
	mappings := make(map[string]([][]uint64))
	categoryIndex := -1

	// Parse the input
	for i, line := range lines {
		// First line is the seeds
		if i == 0 {
			stringSeeds := strings.Split(strings.Split(line, "seeds: ")[1], " ")

			for _, stringSeed := range stringSeeds {
				seed, err := strconv.ParseUint(stringSeed, 10, 64)

				if err != nil {
					panic(err)
				}

				seeds = append(seeds, seed)
			}

			continue
		}

		// Skip empty lines
		if line == "" {
			continue
		}

		// Skip the map: line and increment the category index
		if strings.Contains(line, "map:") {
			categoryIndex++
			continue
		}

		// Parse the mappings
		stringMappings := strings.Split(line, " ")
		uintMappings := []uint64{}

		// fmt.Printf("In category: %v\n", categories[categoryIndex])
		// fmt.Printf("  found string mappings: %v\n", stringMappings)

		for _, stringMapping := range stringMappings {
			mapping, err := strconv.ParseUint(stringMapping, 10, 64)

			if err != nil {
				panic(err)
			}

			uintMappings = append(uintMappings, mapping)
		}

		// fmt.Printf("  found uint mappings: %v\n", uintMappings)

		mappings[categories[categoryIndex]] = append(mappings[categories[categoryIndex]], uintMappings)
	}

	finalMappings := make(map[uint64]uint64)

	for _, seed := range seeds {
		finalMappings[seed] = seed

		for _, category := range categories {
			for _, categoryMappings := range mappings[category] {
				currentMappedValue := finalMappings[seed]

				destStart := categoryMappings[0]
				srcStart := categoryMappings[1]
				length := categoryMappings[2]

				// fmt.Printf("Seed: %v, Category: %v, CurrentMappedValue: %v, SrcStart: %v, DestStart: %v, Length: %v\n", seed, category, currentMappedValue, srcStart, destStart, length)

				if currentMappedValue >= srcStart && currentMappedValue < srcStart+length {
					// fmt.Printf("Seed %v: %v -> %v\n", seed, currentMappedValue, destStart+(currentMappedValue-srcStart))
					finalMappings[seed] = destStart + (currentMappedValue - srcStart)
					break
				}
			}
		}
	}

	closestLocation := ^uint64(0)

	for _, seed := range seeds {
		if finalMappings[seed] < closestLocation {
			closestLocation = finalMappings[seed]
		}
	}

	fmt.Printf("Closest location: %v\n", closestLocation)
}
