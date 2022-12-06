package main

import (
	"strings"
)

func part2(data []string) int {
	answer := 0
	startOfPacket := part1(data)
	for i := startOfPacket; i < len(data); i++ {
		if data[i] == "" {
			continue
		}
		if checkUnique(data[i : i+14]) {
			return i + 14
		}
	}
	return answer
}

func checkUnique(data []string) bool {
	temp := strings.Join(data, "")
	for _, value := range data {
		if strings.Count(temp, value) > 1 {
			return false
		}
	}
	return true
}
