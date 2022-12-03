package main

import (
	"log"
	"math"
	"strconv"
)

func part1(data []string) {
	// Formula: roundDown(mass / 3) - 2
	answer := 0.0
	for _, value := range data {
		if value == "" {
			continue
		}
		mass, _ := strconv.ParseFloat(value, 32)
		answer += math.Round((mass/3)-0.5) - 2
	}
	log.Printf("Answer: %0.f\n", answer)
}
