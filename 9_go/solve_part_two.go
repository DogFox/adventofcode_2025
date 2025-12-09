package main

import (
	"strconv"
)

type Point2 struct {
	x, y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pointOnSegment(px, py, x1, y1, x2, y2 int) bool {
	if (px-x1)*(y2-y1) != (py-y1)*(x2-x1) {
		return false
	}
	if px < min(x1, x2) || px > max(x1, x2) {
		return false
	}
	if py < min(y1, y2) || py > max(y1, y2) {
		return false
	}
	return true
}

var insideMemo = map[[2]int]bool{}

func pointInPolygonCached(px, py int, poly []Point2) bool {
	key := [2]int{px, py}

	if v, ok := insideMemo[key]; ok {
		return v
	}

	v := pointInPolygon(px, py, poly)
	insideMemo[key] = v
	return v
}
func pointInPolygon(px, py int, poly []Point2) bool {
	n := len(poly)
	for i := 0; i < n; i++ {
		x1, y1 := poly[i].x, poly[i].y
		x2, y2 := poly[(i+1)%n].x, poly[(i+1)%n].y
		if pointOnSegment(px, py, x1, y1, x2, y2) {
			return true
		}
	}

	cnt := 0
	for i := 0; i < n; i++ {
		x1, y1 := poly[i].x, poly[i].y
		x2, y2 := poly[(i+1)%n].x, poly[(i+1)%n].y

		if (y1 > py) != (y2 > py) {
			xf := float64(x1) + float64(py-y1)*(float64(x2-x1))/float64(y2-y1)
			if xf > float64(px) {
				cnt++
			}
		}
	}

	return cnt%2 == 1
}

func SolvePart2(list [][]string) int {
	points := make([]Point2, 0, len(list))
	for _, row := range list {
		x, _ := strconv.Atoi(row[0])
		y, _ := strconv.Atoi(row[1])
		points = append(points, Point2{x, y})
	}

	if len(points) == 0 {
		return 0
	}

	minX, maxX := points[0].x, points[0].x
	minY, maxY := points[0].y, points[0].y
	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	n := len(points)
	maxArea := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			p1, p2 := points[i], points[j]

			if p1.x == p2.x && p1.y == p2.y {
				continue
			}

			x1 := min(p1.x, p2.x)
			x2 := max(p1.x, p2.x)
			y1 := min(p1.y, p2.y)
			y2 := max(p1.y, p2.y)

			if !rectBorderInside(x1, y1, x2, y2, points) {
				continue
			}

			area := (x2 - x1 + 1) * (y2 - y1 + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func rectBorderInside(x1, y1, x2, y2 int, poly []Point2) bool {
	for x := x1; x <= x2; x++ {
		if !pointInPolygonCached(x, y1, poly) {
			return false
		}
		if !pointInPolygonCached(x, y2, poly) {
			return false
		}
	}

	for y := y1 + 1; y < y2; y++ {
		if !pointInPolygonCached(x1, y, poly) {
			return false
		}
		if !pointInPolygonCached(x2, y, poly) {
			return false
		}
	}

	return true
}
