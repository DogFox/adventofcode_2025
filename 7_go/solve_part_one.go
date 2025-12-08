package main

import (
	"fmt"
	"os"
)

var directions = [][]int{
	{1, -1},
	{1, 1},
}

type Start struct {
	x int
	y int
}

func FindStart(matrix [][]string) Start {
	start := Start{x: 0, y: 0}
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == "S" {
				start.x = row
				start.y = col
			}
		}
	}

	return start
}

func showBeam(matrix [][]string, start Start) {
	newX := start.x
	newY := start.y

	if newX >= len(matrix) || newY >= len(matrix[0]) || newX < 0 || newY < 0 {
		return
	}

	if matrix[newX][newY] == "." || matrix[newX][newY] == "S" {
		matrix[newX][newY] = "|"
		showBeam(matrix, Start{x: start.x + 1, y: start.y})
	}

	if matrix[newX][newY] == "^" {
		showBeam(matrix, Start{x: start.x, y: start.y - 1})
		showBeam(matrix, Start{x: start.x, y: start.y + 1})
	}
}

func SolvePart1(matrix [][]string) int {
	totalCount := 0
	start := FindStart(matrix)

	showBeam(matrix, start)
	for indR, row := range matrix {
		for intC, cell := range row {
			if cell == "^" && matrix[indR-1][intC] == "|" {
				totalCount++
			}
		}
	}

	if err := writeMatrixToFile(matrix, "output.txt"); err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	}

	return totalCount
}

func writeMatrixToFile(matrix [][]string, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, row := range matrix {
		for _, cell := range row {
			_, _ = f.WriteString(cell)
		}
		_, _ = f.WriteString("\n")
	}

	return nil
}
