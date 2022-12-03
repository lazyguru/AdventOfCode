package main

import (
	"strings"
)

func part2(data []string) int {
	answer := 0
	for i := 0; i < len(data); i += 3 {
		if data[i] == "" {
			continue
		}
		row1 := data[i]
		row2 := data[i+1]
		row3 := data[i+2]
		var match string
		for idx, _ := range row1 {
			if strings.Contains(row2, string(row1[idx])) &&
				strings.Contains(row3, string(row1[idx])) {
				match = string(row1[idx])
				answer += priorities[match]
				break
			}
		}
	}
	return answer
}
