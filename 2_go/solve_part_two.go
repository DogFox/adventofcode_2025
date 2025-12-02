package main

import (
	"strconv"
	"strings"
)

func isRepeatingPattern(s string) bool {
	if len(s) < 2 {
		return false
	}

	for patternLen := 1; patternLen <= len(s)/2; patternLen++ {
		if len(s)%patternLen != 0 {
			continue
		}

		pattern := s[:patternLen]
		repeats := len(s) / patternLen

		isValid := true
		for i := 0; i < repeats; i++ {
			start := i * patternLen
			end := start + patternLen
			if s[start:end] != pattern {
				isValid = false
				break
			}
		}

		if isValid && repeats >= 2 {
			return true
		}
	}

	return false
}

func SolvePart2(lines []string) int {
	totalCount := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)

		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			if isRepeatingPattern(str) {
				totalCount += i
			}
		}
	}

	return totalCount
}
