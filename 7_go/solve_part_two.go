package main

import "fmt"

func SolvePart2(matrix [][]string) int {
	totalCount := 0
	start := FindStart(matrix)

	showBeam(matrix, start)

	if err := writeMatrixToFile(matrix, "output.txt"); err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	}

	totalCount = CountTimelines(matrix, start)

	return totalCount
}

func CountTimelines(matrix [][]string, start Start) int {
	h := len(matrix)
	w := len(matrix[0])

	valid := func(r, c int) bool {
		if r < 0 || c < 0 || r >= h || c >= w {
			return false
		}
		cell := matrix[r][c]
		return cell == "|" || cell == "^"
	}

	memo := make([][]int, h)
	seen := make([][]bool, h)
	for i := 0; i < h; i++ {
		memo[i] = make([]int, w)
		seen[i] = make([]bool, w)
	}

	return dfs(start.x, start.y, matrix, valid, memo, seen)
}

func dfs(r, c int, matrix [][]string, valid func(int, int) bool,
	memo [][]int, seen [][]bool) int {

	h := len(matrix)
	if r == h-1 {
		return 1
	}

	if seen[r][c] {
		return memo[r][c]
	}
	seen[r][c] = true

	cell := matrix[r][c]
	total := 0

	switch cell {
	case "|":
		nr, nc := r+1, c
		if valid(nr, nc) {
			total += dfs(nr, nc, matrix, valid, memo, seen)
		}

	case "^":
		nr := r + 1
		left, right := c-1, c+1

		if valid(nr, left) {
			total += dfs(nr, left, matrix, valid, memo, seen)
		}
		if valid(nr, right) {
			total += dfs(nr, right, matrix, valid, memo, seen)
		}
	}

	memo[r][c] = total
	return total
}
