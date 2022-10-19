package main

import (
	"log"
	"math"
)

func part2() {
	data := ReadFile()
	pos := GetPos(data[0])
	var sum, avg float64
	for _, value := range pos {
		sum += value
	}
	avg = math.Floor(sum / float64(len(pos)))
	var cost float64
	for _, value := range pos {
		dist := math.Abs(avg - value)
		mycost := GetCost(dist)
		cost += mycost
	}
	log.Printf("Final count: %d\n", int(cost))
}

func GetCost(dist float64) float64 {
	var cost, tmp float64
	for tmp = dist; tmp > 0; tmp-- {
		cost += tmp
	}
	return cost
}
