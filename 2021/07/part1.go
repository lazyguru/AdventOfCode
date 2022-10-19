package main

import (
	"log"
	"math"
)

func part1() {
	data := ReadFile()
	pos := GetPos(data[0])
	median := pos[(len(pos)+1)/2]
	var cnt float64
	for _, value := range pos {
		cnt += math.Abs(median - value)
	}
	log.Printf("Final count: %d\n", int(cnt))
}
