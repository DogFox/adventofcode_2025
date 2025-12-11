package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Button struct {
	Idx []int
}

func SolvePart1(list [][]string) int {
	totalCount := 0

	for _, row := range list {
		// fmt.Println(row)
		etalon := make([]int, len(row[0])-2)
		for i, val := range row[0] {
			if val == '[' || val == ']' {
				continue
			}
			if val == '#' {
				etalon[i-1] = 1
			}
		}

		swithces := row[1 : len(row)-1]
		var buttons []Button
		for _, str := range swithces {
			str = strings.Replace(str, "(", "", -1)
			str = strings.Replace(str, ")", "", -1)

			parts := strings.Split(str, ",")
			var idx []int
			for _, p := range parts {
				num, _ := strconv.Atoi(strings.TrimSpace(p))
				idx = append(idx, num)
			}
			buttons = append(buttons, Button{Idx: idx})

		}

		fmt.Println(buttons)
		totalCount += switching(etalon, buttons)
	}

	return totalCount
}

func switching(etalon []int, swithces []Button) int {
	n := len(swithces)
	lights := make([]int, len(etalon))

	best := -1

	var dfs func(pos int, presses int)
	dfs = func(pos int, presses int) {
		if pos == n {
			for i := range lights {
				if lights[i] != etalon[i] {
					return
				}
			}
			if best == -1 || presses < best {
				best = presses
			}
			return
		}

		dfs(pos+1, presses)

		for _, id := range swithces[pos].Idx {
			lights[id] ^= 1
		}

		dfs(pos+1, presses+1)

		for _, id := range swithces[pos].Idx {
			lights[id] ^= 1
		}
	}

	dfs(0, 0)
	return best
}
