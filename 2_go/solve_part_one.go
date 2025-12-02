package main

import (
	"strconv"
	"strings"
)

func SolvePart1(lines []string) int {
	totalCount := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)

		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {

			str := strconv.Itoa(i)
			mid := len(str) / 2
			part1 := str[:mid]
			part2 := str[mid:]
			if part1 == part2 {
				totalCount += i
			}
		}
	}

	return totalCount
}
