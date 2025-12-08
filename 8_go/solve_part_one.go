package main

import (
	"math"
	"sort"
	"strconv"
)

type Box struct {
	x, y, z float64
}

type KV struct {
	i, j int
	dist float64
}

func Distance(a, b Box) float64 {
	dx := b.x - a.x
	dy := b.y - a.y
	dz := b.z - a.z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func SolvePart1(list [][]string, countEdges int) int {
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
	for k := 0; k < countEdges; k++ {
		e := distances[k]
		uf.Union(e.i, e.j)
	}

	sizes := uf.GetComponentSizes()

	sort.Ints(sizes)
	uniqueSizes := make([]int, 0, len(sizes))
	last := -1
	for _, s := range sizes {
		if s != last {
			uniqueSizes = append(uniqueSizes, s)
			last = s
		}
	}

	sort.Slice(uniqueSizes, func(i, j int) bool {
		return uniqueSizes[i] > uniqueSizes[j]
	})

	totalCount := 1
	for i := 0; i < 3; i++ {
		totalCount *= uniqueSizes[i]
	}

	return totalCount
}
