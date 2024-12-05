package main

import (
	"slices"
	"strconv"
	"strings"
)

func part2(data []string) int {
	answer := 0
	listLeft := make([]int, 0)
	listRight := make([]int, 0)
	for _, value := range data {
		if value == "" {
			continue
		}
		strs := strings.Split(value, "   ")
		left, _ := strconv.Atoi(strs[0])
		right, _ := strconv.Atoi(strs[1])
		listLeft = append(listLeft, left)
		listRight = append(listRight, right)
	}
	slices.Sort(listLeft)
	slices.Sort(listRight)
	for i := 0; i < len(listLeft); i++ {
		idx := slices.Index(listRight, listLeft[i])
		cnt := 0
		for idx != -1 && listRight[idx] == listLeft[i] {
			cnt += 1
			idx += 1
		}
		answer += cnt * listLeft[i]
	}
	return answer
}
