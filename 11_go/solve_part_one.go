package main

import "strings"

func SolvePart1(list []string) int {
	totalCount := 0
	graph := make(map[string][]string)

	for _, line := range list {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, ":")
		name := strings.TrimSpace(parts[0])
		var targets []string
		targets = append(targets, strings.Fields(parts[1])...)
		graph[name] = targets
	}

	paths := findAllPaths(graph, "you", "out")
	totalCount += len(paths)

	return totalCount
}

func findAllPaths(graph map[string][]string, start, end string) [][]string {
	var result [][]string
	var dfs func(node string, path []string, visited map[string]bool)

	dfs = func(node string, path []string, visited map[string]bool) {
		if node == end {
			cp := append([]string{}, path...)
			result = append(result, cp)
			return
		}

		for _, next := range graph[node] {
			if visited[next] {
				continue
			}

			visited[next] = true
			dfs(next, append(path, next), visited)
			delete(visited, next)
		}
	}

	visited := map[string]bool{start: true}
	dfs(start, []string{start}, visited)

	return result
}
