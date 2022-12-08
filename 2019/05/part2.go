package main

import (
	"strings"
)

func part2(data []string) int {
	answer := -1
	input := strings.Split(data[0], ",")
	answer = intCodeProgram(input, "5")
	return answer
}
