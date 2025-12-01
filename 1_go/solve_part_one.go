package main

import (
	"strconv"
	"strings"
)

func SolvePart1(lines []string) int {
	currentPosition := 50
	timesAtZero := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "L") {
			clicks, _ := strconv.Atoi(strings.TrimPrefix(line, "L"))
			currentPosition -= clicks
		} else {

			clicks, _ := strconv.Atoi(strings.TrimPrefix(line, "R"))
			currentPosition += clicks
		}
		if currentPosition < 0 {
			currentPosition = 100 + (currentPosition % 100)
		}

		if currentPosition >= 100 {
			currentPosition = currentPosition % 100
		}

		if currentPosition == 0 {
			timesAtZero++
		}

	}

	return timesAtZero
}
