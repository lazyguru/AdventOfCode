package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	chunks := ReadFile()
	elves := make(map[int]int)
	keys := make([]int, 0, len(elves))
	for idx, chunk := range chunks {
		for _, value := range strings.Split(chunk, "\n") {
			if value == "" {
				continue
			}
			cals, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			elves[idx+1] += cals
		}
		keys = append(keys, idx+1)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return elves[keys[i]] > elves[keys[j]]
	})
	log.Printf("Answer: %d\n", elves[keys[0]])
}
