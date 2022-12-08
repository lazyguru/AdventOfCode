package main

import (
	"strings"
)

func part1(data []string) int {
	answer := -1
	input := strings.Split(data[0], ",")
	answer = intCodeProgram(input, "1")
	return answer
}
