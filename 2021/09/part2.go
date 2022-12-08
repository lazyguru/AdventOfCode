package main

import (
	"strings"
)

func part2(data []string) int {
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
	basins := make([]int, 3, 3)
	for a := 0; a < len(rows); a++ {
		cols := rows[a]
		for b := 0; b < len(cols); b++ {
			if checkDepth(b, cols, a, rows) {
				basin := getBasinSize(b, cols, a, rows)
				for k, v := range basins {
					if v < basin {
						basins[k] = basin
					}
				}
			}
		}
	}
	return basins[0] * basins[1] * basins[2]
}

func getBasinSize(b int, cols []int, a int, rows [][]int) int {
	i := 1
	cnt := 1
	for b+i < len(cols) && cols[b+i] != 9 {
		cnt += getBasinSize(b+i, cols, a, rows)
		i++
	}
	i = 1
	for b-i >= 0 && cols[b-i] != 9 {
		cnt += getBasinSize(b-i, cols, a, rows)
		i++
	}
	i = 1
	for a+i < len(rows) && rows[a+i][b] != 9 {
		cnt += getBasinSize(b, cols, a+i, rows)
		i++
	}
	i = 1
	for a-i >= 0 && rows[a-i][b] != 9 {
		cnt += getBasinSize(b, cols, a-i, rows)
		i++
	}
	return cnt
}
