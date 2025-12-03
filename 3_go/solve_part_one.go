package main

import (
	"strconv"
	"strings"
)

func SolvePart1(lines []string) int {
	totalCount := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)

		firstIndex := -1
		firstDigit := -1
		secondDigit := -1

		for index, runeValue := range line {
			if index+1 != len(line) {
				digit, _ := strconv.Atoi(string(runeValue))

				if digit > firstDigit {
					firstIndex = index
					firstDigit = digit
				}
			}
		}

		part2 := line[firstIndex+1:]
		for _, runeValue := range part2 {
			digit, _ := strconv.Atoi(string(runeValue))

			if digit > secondDigit {
				secondDigit = digit
			}
		}

		if secondDigit == -1 {
			secondDigit = 0
		}

		countStr := strconv.Itoa(firstDigit) + strconv.Itoa(secondDigit)
		count, _ := strconv.Atoi(countStr)
		totalCount += count
	}

	return totalCount
}
