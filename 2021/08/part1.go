package main

import (
	"strings"
)

func part1(data []string) int {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		line := strings.Split(value, " | ")
		output := strings.Split(line[1], " ")

		for _, digit := range output {
			switch len(digit) {
			case 2:
				// 1
				answer++
				break
			case 3:
				// 7
				answer++
				break
			case 4:
				// 4
				answer++
				break
			case 7:
				// 8
				answer++
				break
			}
		}
	}
	return answer
}
