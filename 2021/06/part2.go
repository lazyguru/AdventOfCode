package main

import (
	"log"
	"strconv"
	"strings"
)

// Part2 essentially inspired by https://github.com/c-kk/aoc/blob/master/2021-go/day06/solve.go
// I think I would have eventually gotten here on my own, but in searching for a hint I ruined the puzzle for the day

func part2() {
	data := ReadFile()
	fish := GetFishV2(data[0])

	for day := 1; day < 257; day++ {
		todayFish := make([]int, 9)
		for idx, f := range fish {
			if idx == 0 { // need to reset and spawn new fish
				todayFish[6] += f
				todayFish[8] += f
			} else {
				todayFish[idx-1] += f
			}
		}
		fish = todayFish
	}
	cnt := 0
	for _, f := range fish {
		cnt += f
	}
	log.Printf("Final count: %d\n", cnt)
}

func GetFishV2(data string) []int {
	sFish := strings.Split(data, ",")
	iFish := make([]int, 9)
	for _, value := range sFish {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		iFish[i]++
	}
	return iFish
}
