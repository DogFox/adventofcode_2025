package main

import (
	"strconv"
	"strings"
)

func SolvePart2(lines []string) int {
	totalCount := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)

		maxCountDigit := 12
		lineLength := len(line)

		result := make([]byte, 0, maxCountDigit)
		start := 0

		for maxCountDigit > 0 {
			end := lineLength - maxCountDigit

			bestIndex := start
			for i := start; i <= end; i++ {
				if line[i] > line[bestIndex] {
					bestIndex = i
				}
			}

			result = append(result, line[bestIndex])
			start = bestIndex + 1
			maxCountDigit--
		}

		num, _ := strconv.Atoi(string(result))
		totalCount += num
	}

	return totalCount
}
