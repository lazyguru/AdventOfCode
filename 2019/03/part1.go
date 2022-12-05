package main

import (
	"strings"
)

func part1(data []string) int {
	first := strings.Split(data[0], ",")
	second := strings.Split(data[1], ",")
	pathA := processLine(first)
	pathB := processLine(second)
	path := make(map[Cords]bool)
	for _, point := range pathA {
		path[point] = true
	}
	lastDistance := 0
	for _, point := range pathB {
		if path[point] {
			distance := getDistance(point)
			if lastDistance == 0 || distance < lastDistance {
				lastDistance = distance
			}
		}
	}
	return lastDistance
}
