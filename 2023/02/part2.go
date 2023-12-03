package main

import (
	"strconv"
	"strings"
)

func part2(data []string) int {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		_, gameData, _ := strings.Cut(value, ":")
		pulls := strings.Split(gameData, ";")
		redCount := 0
		greenCount := 0
		blueCount := 0
		for _, pull := range pulls {
			cubes := strings.Split(pull, ",")
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				count, color, _ := strings.Cut(cube, " ")
				icount, _ := strconv.Atoi(count)
				color = strings.Trim(color, " ")
				if color == "red" && icount > redCount {
					redCount = icount
					continue
				}
				if color == "green" && icount > greenCount {
					greenCount = icount
					continue
				}
				if color == "blue" && icount > blueCount {
					blueCount = icount
					continue
				}
			}
		}
		answer += (redCount * blueCount * greenCount)
	}
	return answer
}
