package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
	reds := 12
	greens := 13
	blues := 14
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		id, gameData, _ := strings.Cut(value, ":")
		_, num, _ := strings.Cut(id, " ")
		gameId, _ := strconv.Atoi(num)
		pulls := strings.Split(gameData, ";")
		skip := false
		for _, pull := range pulls {
			cubes := strings.Split(pull, ",")
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				count, color, _ := strings.Cut(cube, " ")
				icount, _ := strconv.Atoi(count)
				color = strings.Trim(color, " ")
				if color == "red" && icount > reds {
					skip = true
					continue
				}
				if color == "green" && icount > greens {
					skip = true
					continue
				}
				if color == "blue" && icount > blues {
					skip = true
					continue
				}
			}
		}
		if skip == false {
			answer += gameId
		}
	}
	return answer
}
