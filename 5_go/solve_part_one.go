package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func ParseInput(lines []string) ([]Range, []int, error) {
	ranges := make([]Range, 0)
	values := make([]int, 0)

	inValuesSection := false

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			inValuesSection = true
			continue
		}

		if !inValuesSection {
			parts := strings.Split(line, "-")

			start, err := strconv.Atoi(parts[0])
			if err != nil {
				return nil, nil, fmt.Errorf("err when convert start: %s", err)
			}

			end, err := strconv.Atoi(parts[1])
			if err != nil {
				return nil, nil, fmt.Errorf("err when convert end: %s", err)
			}

			ranges = append(ranges, Range{Start: start, End: end})
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, fmt.Errorf("err when convert value: %s", err)
			}

			values = append(values, value)
		}
	}

	mergedRanges := mergeRanges(ranges)

	return mergedRanges, values, nil
}

func mergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := make([]Range, 0)
	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]
		if current.End >= next.Start-1 {
			if next.End > current.End {
				current.End = next.End
			}
		} else {
			merged = append(merged, current)
			current = next
		}
	}
	merged = append(merged, current)

	return merged
}

func SolvePart1(ranges []Range, values []int) int {
	totalCount := 0
	for _, value := range values {
		for _, r := range ranges {
			if value >= r.Start && value <= r.End {
				totalCount++
				break
			}
		}
	}

	return totalCount
}
