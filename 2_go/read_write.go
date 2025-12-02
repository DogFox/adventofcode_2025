package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0, 1024)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			for _, part := range parts {
				part = strings.TrimSpace(part)
				if part != "" {
					lines = append(lines, part)
				}
			}
		} else {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func WriteOutput(path string, value any) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprint(f, value)
	return err
}
