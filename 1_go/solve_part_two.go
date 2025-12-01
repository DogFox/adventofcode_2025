package main

import (
	"strconv"
	"strings"
)

func floorDiv(n, d int) int {
	if n >= 0 {
		return n / d
	}
	return -((-n + d - 1) / d)
}

func SolvePart2(lines []string) int {
	position := 50
	timesAtZero := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "L") {
			clicks, _ := strconv.Atoi(strings.TrimPrefix(line, "L"))
			old := position
			position -= clicks

			if position < old {
				timesAtZero += floorDiv(old-1, 100) - floorDiv(position-1, 100)
			}
		} else {
			clicks, _ := strconv.Atoi(strings.TrimPrefix(line, "R"))
			old := position
			position += clicks

			if position > old {
				timesAtZero += floorDiv(position, 100) - floorDiv(old, 100)
			}
		}
	}

	return timesAtZero
}
