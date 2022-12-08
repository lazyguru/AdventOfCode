package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
	answer := 0
	passwordRange := strings.Split(data[0], "-")
	low, _ := strconv.Atoi(passwordRange[0])
	high, _ := strconv.Atoi(passwordRange[1])
	i := low
	for i <= high {
		if !checkForRepeating(i) {
			i++
			continue
		}
		if !checkIncreasing(i) {
			i++
			continue
		}
		i++
		answer++
	}
	return answer
}
