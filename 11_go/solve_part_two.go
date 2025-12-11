package main

import (
	"strings"
)

type stateKey struct {
	node    string
	seenDAC bool
	seenFFT bool
}

func SolvePart2(list []string) int {
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

	memo := make(map[stateKey]int)
	totalCount += dfsMemo(graph, "svr", false, false, memo)
	return totalCount
}

func dfsMemo(
	graph map[string][]string,
	node string,
	seenDAC bool,
	seenFFT bool,
	memo map[stateKey]int,
) int {

	if node == "dac" {
		seenDAC = true
	}
	if node == "fft" {
		seenFFT = true
	}

	if node == "out" {
		if seenDAC && seenFFT {
			return 1
		}
		return 0
	}

	key := stateKey{
		node:    node,
		seenDAC: seenDAC,
		seenFFT: seenFFT,
	}

	if val, ok := memo[key]; ok {
		return val
	}

	total := 0
	for _, next := range graph[node] {
		total += dfsMemo(graph, next, seenDAC, seenFFT, memo)
	}

	memo[key] = total
	return total
}
