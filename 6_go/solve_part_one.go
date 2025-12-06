package main

import (
	"fmt"
	"strconv"
)

func SolvePart1(lines [][]string) int {
	totalCount := 0
	for col := range lines[0] {
		result, _ := strconv.Atoi(lines[0][col])
		op := lines[len(lines)-1][col]

		for row := 1; row < len(lines)-1; row++ {
			el, _ := strconv.Atoi(lines[row][col])

			switch op {
			case "*":
				result *= el
			case "+":
				result += el
			}
		}

		fmt.Printf("col %d (%c) result = %d\n", col, op, result)
		totalCount += result
	}

	return totalCount
}
