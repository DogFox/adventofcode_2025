package main

import "sort"

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &UnionFind{parent: parent, size: size}
}

func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	rx, ry := uf.Find(x), uf.Find(y)
	if rx == ry {
		return false
	}

	if uf.size[rx] < uf.size[ry] {
		rx, ry = ry, rx
	}

	uf.parent[ry] = rx
	uf.size[rx] += uf.size[ry]
	return true
}

func (uf *UnionFind) GetComponentSizes() []int {
	roots := make(map[int]struct{})

	for i := range uf.parent {
		root := uf.Find(i)
		roots[root] = struct{}{}
	}

	sizes := make([]int, 0, len(roots))
	for root := range roots {
		sizes = append(sizes, uf.size[root])
	}

	sort.Ints(sizes)
	return sizes
}
