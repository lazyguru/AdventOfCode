package main

import (
	"strconv"
	"strings"
)

func part2(data []string) int {
	//exampleData := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	//	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	//	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	//	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	//	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}
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
