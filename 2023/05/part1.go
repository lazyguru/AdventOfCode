package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
	_, seedLine, _ := strings.Cut(data[0], ": ")
	seeds := GetSeedsP1(seedLine)
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
	return GetAnswer(seeds, maps)
}

func GetSeedsP1(seedLine string) []*Seed {
	seedSlice := strings.Split(seedLine, " ")
	var seeds []*Seed
	for _, seed := range seedSlice {
		s, _ := strconv.Atoi(seed)
		seeds = append(seeds, &Seed{Id: s})
	}
	return seeds
}
