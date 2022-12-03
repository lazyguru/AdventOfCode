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
		half := len(value) / 2
		half1 := value[:half]
		half2 := value[half:]
		var match string
		for idx, _ := range half1 {
			if strings.Contains(half2, string(half1[idx])) {
				match = string(half1[idx])
				answer += priorities[match]
				break
			}
		}
	}
	return answer
}
