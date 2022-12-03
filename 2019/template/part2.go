package main

import (
	"log"
)

func part2(data []string) {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		answer++
	}
	log.Printf("Answer: %d\n", answer)
}
