package main

import (
	"strings"
)

func SolvePart2(lines []string) int {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		grid[i] = []rune(line)
	}

	totalCount := 0
	changed := true

	for changed {
		changed = false
		toChange := make([][]int, 0)

		for rowIndex := 0; rowIndex < len(grid); rowIndex++ {
			for colIndex := 0; colIndex < len(grid[rowIndex]); colIndex++ {
				if grid[rowIndex][colIndex] == '@' {
					count := 0
					for _, d := range directions {
						newX := rowIndex + d[0]
						newY := colIndex + d[1]

						if newX >= len(grid) || newY >= len(grid[0]) || newX < 0 || newY < 0 {
							continue
						}
						if grid[newX][newY] == '@' {
							count++
						}
					}

					if count < 4 {
						toChange = append(toChange, []int{rowIndex, colIndex})
					}
				}
			}
		}

		for _, pos := range toChange {
			grid[pos[0]][pos[1]] = '.'
			totalCount++
			changed = true
		}
	}

	return totalCount
}
