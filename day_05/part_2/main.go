package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/iuliancmarcu/advent-of-code-2023/common"
)

// seed -> soil -> fertilizer -> water -> light -> temperature -> humidity -> location

var (
	categories = []string{"soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
)

type Interval struct {
	Start uint64
	End   uint64
}

type MappingInterval struct {
	SrcStart uint64
	SrcEnd   uint64
	DstStart uint64
	DstEnd   uint64
}

func (mappingInterval *MappingInterval) OverlapsInterval(interval Interval) (bool, Interval, Interval) {
	overlappingInterval := Interval{}
	overlappingInterval.Start = max(mappingInterval.SrcStart, interval.Start)
	overlappingInterval.End = min(mappingInterval.SrcEnd, interval.End)

	overlaps := overlappingInterval.Start <= overlappingInterval.End

	mappedInterval := Interval{}
	if overlaps {
		offset := mappingInterval.DstStart - mappingInterval.SrcStart

		mappedInterval.Start = overlappingInterval.Start + offset
		mappedInterval.End = overlappingInterval.End + offset
	}

	return overlaps, overlappingInterval, mappedInterval
}

func parseInput() ([]Interval, map[string][]MappingInterval) {
	lines := common.ReadFile("day_05/input.txt")

	var seedIntervals []Interval

	mappings := make(map[string]([]MappingInterval))
	categoryIndex := -1

	// Parse the input
	for i, line := range lines {
		// First line is the seed intervals
		if i == 0 {
			stringSeeds := strings.Split(strings.Split(line, "seeds: ")[1], " ")

			var pair []uint64
			for _, stringSeed := range stringSeeds {
				seed, err := strconv.ParseUint(stringSeed, 10, 64)

				if err != nil {
					panic(err)
				}

				pair = append(pair, seed)

				// When we collected 2 numbers, create interval
				if len(pair) == 2 {
					start := pair[0]
					length := pair[1]

					seedIntervals = append(seedIntervals, Interval{Start: start, End: start + length - 1})
					pair = slices.Delete(pair, 0, 2)
				}
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

		mappings[categories[categoryIndex]] = append(
			mappings[categories[categoryIndex]],
			MappingInterval{
				SrcStart: uintMappings[1],
				SrcEnd:   uintMappings[1] + uintMappings[2] - 1,
				DstStart: uintMappings[0],
				DstEnd:   uintMappings[0] + uintMappings[2] - 1,
			},
		)
	}

	return seedIntervals, mappings
}

func main() {
	seedIntervals, mappings := parseInput()

	intervals := seedIntervals

	// For each category, map the latest intervals
	for _, category := range categories {
		categoryMappings := mappings[category]

		newIntervals := make([]Interval, 0)
		for _, interval := range intervals {
			// fmt.Printf("Mapping interval [%v, %v]\n", interval.Start, interval.End)

			// collect the mapped intervals and the overlapped areas of the current interval
			overlappingsApplied := make([]Interval, 0)
			for _, mappingInterval := range categoryMappings {
				overlaps, overlappingInt, resultInt := mappingInterval.OverlapsInterval(interval)

				if overlaps {
					overlappingsApplied = append(overlappingsApplied, overlappingInt)
					newIntervals = append(newIntervals, resultInt)

					// fmt.Printf("  mapped [%v, %v] to [%v, %v]\n",
					// 	overlappingInt.Start,
					// 	overlappingInt.End,
					// 	resultInt.Start,
					// 	resultInt.End)
				}
			}

			// check if the whole interval has been mapped
			if len(overlappingsApplied) == 1 &&
				overlappingsApplied[0].Start == interval.Start &&
				overlappingsApplied[0].End == interval.End {
				// whole interval was mapped
				continue
			}

			// otherwise find holes in the interval

			// add the interval of everything before and after the current interval
			beforeInterval := Interval{Start: 0, End: interval.Start - 1}
			afterInterval := Interval{Start: interval.End + 1, End: ^uint64(0)}
			overlappingsApplied = append(overlappingsApplied, beforeInterval, afterInterval)

			// sort intervals
			sort.Slice(overlappingsApplied, func(i, j int) bool {
				return overlappingsApplied[i].Start < overlappingsApplied[j].Start
			})

			// figure out holes
			holes := make([]Interval, 0)
			prevInterval := overlappingsApplied[0]
			for _, interv := range overlappingsApplied[1:] {
				if interv.Start-prevInterval.End > 1 {
					hole := Interval{Start: prevInterval.End + 1, End: interv.Start}
					holes = append(holes, hole)
					// fmt.Printf("  hole at [%v, %v]\n", hole.Start, hole.End)
				}

				prevInterval = interv
			}

			// append the holes for the next category mapping
			newIntervals = append(newIntervals, holes...)
		}

		// submit the mapped intervals for the next mapping
		intervals = newIntervals
	}

	// find the smallest interval
	closestLocation := ^uint64(0)
	for _, interval := range intervals {
		if interval.Start < closestLocation {
			closestLocation = interval.Start
		}
	}

	fmt.Printf("Closest location: %v\n", closestLocation)
}
