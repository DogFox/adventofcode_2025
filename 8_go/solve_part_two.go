package main

import (
	"fmt"
	"sort"
	"strconv"
)

func SolvePart2(list [][]string) int {
	boxes := make([]Box, 0, len(list))
	for _, row := range list {
		x, _ := strconv.ParseFloat(row[0], 64)
		y, _ := strconv.ParseFloat(row[1], 64)
		z, _ := strconv.ParseFloat(row[2], 64)
		boxes = append(boxes, Box{x: x, y: y, z: z})
	}

	n := len(boxes)

	capacity := n * (n - 1) / 2
	distances := make([]KV, 0, capacity)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := Distance(boxes[i], boxes[j])
			distances = append(distances, KV{i, j, dist})
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].dist < distances[j].dist
	})

	uf := NewUnionFind(n)
	components := n
	var lastEdge KV
	for _, edge := range distances {
		if components == 1 {
			break
		}
		if uf.Union(edge.i, edge.j) {
			components--
			if components == 1 {
				lastEdge = edge
				break
			}
		}
	}

	totalCount := int(boxes[lastEdge.i].x) * int(boxes[lastEdge.j].x)

	fmt.Println(lastEdge.i, lastEdge.j, boxes[lastEdge.i], boxes[lastEdge.j], totalCount)
	return totalCount
}
