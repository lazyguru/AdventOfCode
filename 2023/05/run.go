package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Destination int
	Source      int
	Length      int
}

type Seed struct {
	Id          int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

type AlmanacMap = map[int]Range

func GetValueFromMap(search int, myMap AlmanacMap) int {
	for sourceStart, s := range myMap {
		sourceEnd := sourceStart + s.Length
		if search >= sourceStart && search <= sourceEnd {
			adj := search - sourceStart
			//log.Println(search, sourceStart, sourceEnd, adj, s.Destination+adj)
			return s.Destination + adj
		}
	}
	return search
}

func GetAnswer(seeds []*Seed, maps map[string]AlmanacMap) int {
	answer := -1
	for _, seed := range seeds {
		//log.Println("Seed", seed.Id)
		seed.Soil = GetValueFromMap(seed.Id, maps["seed-to-soil"])
		//log.Println("Soil", seed.Soil)
		seed.Fertilizer = GetValueFromMap(seed.Soil, maps["soil-to-fertilizer"])
		//log.Println("Fertilizer", seed.Fertilizer)
		seed.Water = GetValueFromMap(seed.Fertilizer, maps["fertilizer-to-water"])
		//log.Println("Water", seed.Water)
		seed.Light = GetValueFromMap(seed.Water, maps["water-to-light"])
		//log.Println("Light", seed.Light)
		seed.Temperature = GetValueFromMap(seed.Light, maps["light-to-temperature"])
		//log.Println("Temperature", seed.Temperature)
		seed.Humidity = GetValueFromMap(seed.Temperature, maps["temperature-to-humidity"])
		//log.Println("Humidity", seed.Humidity)
		seed.Location = GetValueFromMap(seed.Humidity, maps["humidity-to-location"])
		//log.Println("Location", seed.Location)
	}
	for _, seed := range seeds {
		//log.Println("L", seed.Location, answer)
		if answer < 0 || answer > seed.Location {
			answer = seed.Location
		}
	}
	return answer
}

func PopulateMap(maps map[string]AlmanacMap, mapName string, nums []int) {
	maps[mapName][nums[1]] = Range{
		Destination: nums[0],
		Source:      nums[1],
		Length:      nums[2],
	}
}

func GetNums(value string) []int {
	var nums []int
	for _, val := range strings.Split(value, " ") {
		i, _ := strconv.Atoi(val)
		nums = append(nums, i)
	}
	return nums
}

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

func main() {
	filename := "input.txt"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	log.Printf("Reading file: %s\n", filename)
	content := ReadFile(filename)
	//log.Println("Running part 1")
	ans1 := part1(content)
	log.Printf("Part1: %d\n", ans1)
	//log.Println("Running part 2")
	ans2 := part2(content)
	log.Printf("Part2: %d\n", ans2)

}
