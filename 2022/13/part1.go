package main

import (
	"encoding/json"
)

func part1(data []string) int {
	answer := 0
	pairId := 0
	for idx := 0; idx < len(data)-1; idx += 3 {
		pairId++
		row1 := data[idx]
		row2 := data[idx+1]
		var list1 []any
		var list2 []any
		_ = json.Unmarshal([]byte(row1), &list1)
		_ = json.Unmarshal([]byte(row2), &list2)
		if process(list1, list2) == 1 {
			answer += pairId
		}
	}
	return answer
}
