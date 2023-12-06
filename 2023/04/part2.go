package main

import (
	"strconv"
	"strings"
)

func part2(data []string) int {
	answer := 0
	cardMultipliers := make(map[int]int)
	for _, value := range data {
		if value == "" {
			continue
		}
		cardId, numbers, _ := strings.Cut(value, ": ")
		_, cardNum, _ := strings.Cut(cardId, " ")
		cardN, _ := strconv.Atoi(strings.TrimSpace(cardNum))
		cardMultipliers[cardN] += 1
		winning, held, _ := strings.Cut(strings.TrimSpace(numbers), " | ")
		win := strings.Split(strings.TrimSpace(winning), " ")
		hand := strings.Split(strings.TrimSpace(held), " ")
		var matches []string
		for _, num := range win {
			if strings.TrimSpace(num) == "" {
				continue
			}
			if Contains(hand, num) {
				matches = append(matches, num)
			}
		}
		for i := 0; i < len(matches); i++ {
			cardMultipliers[cardN+i+1] += cardMultipliers[cardN]
		}
		answer += cardMultipliers[cardN]
	}
	return answer
}
