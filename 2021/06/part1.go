package main

import (
	"log"
)

func part1() {
	data := ReadFile()
	fish := GetFish(data[0])

	for day := 1; day < 81; day++ {
		fishToday := fish // one time clone to make sure we don't process new fish same day we spawn them
		for idx, f := range fishToday {
			fish[idx]--
			if f == 0 { // need to reset and spawn new fish
				fish[idx] = 6
				fish = append(fish, 8)
			}
		}
	}
	cnt := len(fish)
	log.Printf("Final count: %d\n", cnt)
}
