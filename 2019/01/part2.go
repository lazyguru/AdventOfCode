package main

import (
	"log"
	"math"
	"strconv"
)

func part2(data []string) {
	// Formula: roundDown(mass / 3) - 2
	answer := 0.0
	for _, value := range data {
		if value == "" {
			continue
		}
		mass, _ := strconv.ParseFloat(value, 32)
		for mass > 0 {
			fuel := calcFuel(mass)
			answer += fuel
			mass = fuel
		}
	}
	log.Printf("Answer: %0.f\n", answer)
}

func calcFuel(mass float64) float64 {
	if mass <= 0 {
		return 0
	}
	fuel := math.Round((mass/3)-0.5) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel
}
