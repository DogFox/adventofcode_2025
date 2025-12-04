package main

import (
	"strings"
)

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func SolvePart1(lines []string) int {
	totalCount := 0
	for rowIndex, line := range lines {
		line = strings.TrimSpace(line)
		for runeIndex, runeValue := range line {
			if string(runeValue) == "@" {
				count := 0
				for _, d := range directions {

					newX := rowIndex + d[0]
					newY := runeIndex + d[1]

					if newX >= len(lines) || newY >= len(lines[0]) || newX < 0 || newY < 0 {
						continue
					}
					if string(lines[newX][newY]) == "@" {
						count++
					}
				}

				if count < 4 {
					// fmt.Println("x ", rowIndex, " y ", runeIndex)
					totalCount++
				}
			}

		}
	}

	return totalCount
}
