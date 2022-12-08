package main

import (
	"strings"
)

func part1(data []string) int {
	answer := 0
	dataLen := len(data)
	if data[len(data)-1] == "" {
		dataLen--
	}
	rows := make([][]int, dataLen)
	for idx, value := range data {
		if value == "" {
			continue
		}
		cols := strings.Split(value, "")
		rows[idx] = getCols(cols)
	}
	for a := 0; a < len(rows); a++ {
		cols := rows[a]
		for b := 0; b < len(cols); b++ {
			if checkDepth(b, cols, a, rows) {
				answer += 1 + cols[b]
			}
		}
	}
	return answer
}
