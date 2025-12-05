package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	const (
		inputPath  = "input.txt"
		outputPath = "output.txt"
	)

	lines, err := ReadInput(inputPath)
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	ranges, values, err := ParseInput(lines)
	if err != nil {
		log.Fatalf("failed to parse input: %v", err)
	}

	var (
		start  = time.Now()
		result any
	)

	result = SolvePart2(ranges, values)

	elapsed := time.Since(start)

	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Elapsed: %s\n", elapsed)

	if err := WriteOutput(outputPath, result); err != nil {
		log.Fatalf("failed to write output: %v", err)
	}
}
