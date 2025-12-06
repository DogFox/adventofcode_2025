package main

import (
	"strconv"
	"strings"
	"unicode"
)

func readColumnsGrouped(matrix [][]rune) [][]int {
	maxWidth := 0
	for _, row := range matrix {
		if len(row) > maxWidth {
			maxWidth = len(row)
		}
	}

	var result [][]int
	var currentGroup []int

	for j := 0; j < maxWidth; j++ {
		var digits []rune
		hasDigit := false

		for i := 0; i < len(matrix); i++ {
			var char rune
			if j < len(matrix[i]) {
				char = matrix[i][j]
			} else {
				char = ' '
			}
			if unicode.IsDigit(char) {
				digits = append(digits, char)
				hasDigit = true
			}
		}

		if hasDigit {
			numStr := string(digits)
			if num, err := strconv.Atoi(numStr); err == nil {
				currentGroup = append(currentGroup, num)
			}
		} else {
			if len(currentGroup) > 0 {
				result = append(result, currentGroup)
				currentGroup = nil
			}
		}
	}
	if len(currentGroup) > 0 {
		result = append(result, currentGroup)
	}

	return result
}

func SolvePart2(lines [][]rune) int {
	totalCount := 0
	opLineStr := string(lines[len(lines)-1])
	ops := strings.Fields(opLineStr)

	groups := readColumnsGrouped(lines)

	for ind, terms := range groups {
		result := terms[0]
		op := ops[ind]

		for i := 1; i < len(terms); i++ {
			el := terms[i]

			switch op {
			case "*":
				result *= el
			case "+":
				result += el
			}
		}

		totalCount += result
	}

	return totalCount
}
