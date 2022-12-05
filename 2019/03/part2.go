package main

import (
	"strings"
)

func part2(data []string) int {
	first := strings.Split(data[0], ",")
	second := strings.Split(data[1], ",")
	pathA := processLine(first)
	pathB := processLine(second)
	path := make(map[Cords]bool)
	for _, point := range pathA {
		path[point] = true
	}
	stepA := 0
	stepB := 0
	for step, point := range pathB {
		if path[point] {
			stepAtCollision := getStep(pathA, point)
			if stepB == 0 {
				stepB = step
				stepA = stepAtCollision
				continue
			}
			if (step + stepAtCollision) < (stepA + stepB) {
				stepB = step
				stepA = stepAtCollision
			}
		}
	}
	return stepA+stepB+2 // +2 to account for 0,0
}

func getStep(p []Cords, c Cords) int {
	for step, point := range p {
		if compare(point, c) {
			return step
		}
	}
	return -1
}

func compare(a Cords, b Cords) bool {
	return a.X == b.X && a.Y == b.Y
}
