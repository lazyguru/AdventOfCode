package main

import (
	"sort"
	"strconv"
	"strings"
)

func GetSeedsP2(seedLine string) []*Seed {
	seedSlice := strings.Split(seedLine, " ")
	var seeds []*Seed
	for idx := 0; idx < len(seedSlice); idx += 2 {
		j1, _ := strconv.Atoi(seedSlice[idx])
		j2, _ := strconv.Atoi(seedSlice[idx+1])
		for j := j1; j < (j1 + j2); j++ {
			seeds = append(seeds, &Seed{Id: j})
		}
	}
	return seeds
}

func part2(data []string) int {
	_, seedLine, _ := strings.Cut(data[0], ": ")
	seeds := GetSeedsP2(seedLine)
	mapName := ""
	maps := make(map[string]AlmanacMap)
	for _, value := range data[1:] {
		if value == "" {
			continue
		}
		if strings.Contains(value, "map") {
			mapName, _, _ = strings.Cut(value, " ")
			maps[mapName] = AlmanacMap{}
			continue
		}
		nums := GetNums(value)
		PopulateMap(maps, mapName, nums)
	}
	batchSize := 10000
	batches := (len(seeds) / batchSize) + 1
	answer := make(chan int, batches)
	for i := 0; i < batches; i++ {
		start := batchSize * i
		end := start + batchSize
		if start > len(seeds) {
			break
		}
		if end > len(seeds) {
			end = len(seeds)
		}
		go func() { answer <- GetAnswer(seeds[start:end], maps) }()
	}
	var answers []int
	for i := 0; i < batches; i++ {
		t := <-answer
		answers = append(answers, t)
	}
	sort.Ints(answers)
	return answers[0]
}
