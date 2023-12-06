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
		_, numbers, _ := strings.Cut(value, ": ")
		winning, held, _ := strings.Cut(strings.TrimSpace(numbers), " | ")
		win := strings.Split(strings.TrimSpace(winning), " ")
		hand := strings.Split(strings.TrimSpace(held), " ")
		cardValue := 0
		var matches []string
		for _, num := range win {
			if strings.TrimSpace(num) == "" {
				continue
			}
			if Contains(hand, num) {
				matches = append(matches, num)
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}
		answer += cardValue
	}
	return answer
}
