package main

func SolvePart2(ranges []Range, _ []int) int {
	totalCount := 0
	for _, r := range ranges {
		totalCount += r.End - r.Start + 1
	}

	return totalCount
}
