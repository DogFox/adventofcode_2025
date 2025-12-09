package main

import (
	"math"
	"strconv"
)

type Point struct {
	x, y float64
}

func Square(a, b Point) float64 {
	dx := b.x - a.x
	dy := b.y - a.y
	return math.Sqrt(dx*dx + dy*dy)
}

func SolvePart1(list [][]string) int {
	points := make([]Point, 0, len(list))
	for _, row := range list {
		x, _ := strconv.ParseFloat(row[0], 64)
		y, _ := strconv.ParseFloat(row[1], 64)
		points = append(points, Point{x: x, y: y})
	}

	totalCount := 0
	var bestPair [2]Point

	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1 := points[i]
			p2 := points[j]

			if p1.x == p2.x && p1.y == p2.y {
				continue
			}

			area := int(math.Abs(float64(p1.x-p2.x)+1) * math.Abs(float64(p1.y-p2.y)+1))
			// fmt.Println(p1, p2, area)
			// fmt.Println(p1.x-p2.x, p1.y-p2.y)
			if area > totalCount {
				totalCount = area
				bestPair[0] = p1
				bestPair[1] = p2
			}
		}
	}

	return totalCount
}
