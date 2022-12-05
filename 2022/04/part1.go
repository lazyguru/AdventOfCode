package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		pairs := strings.Split(value, ",")
		range1 := strings.Split(pairs[0], "-")
		range2 := strings.Split(pairs[1], "-")
		r1min, _ := strconv.Atoi(range1[0])
		r1max, _ := strconv.Atoi(range1[1])
		r2min, _ := strconv.Atoi(range2[0])
		r2max, _ := strconv.Atoi(range2[1])
		if r1min <= r2min && r1max >= r2max {
			answer++
		} else if r1min >= r2min && r1max <= r2max {
			answer++
		}
	}
	return answer
}
