package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

func getDir(a, b int) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}

func getLevels(value string) []int {
	levelsStr := strings.Split(value, " ")
	levels := make([]int, len(levelsStr))
	for i := range levelsStr {
		val, err := strconv.Atoi(levelsStr[i])
		if err != nil {
			log.Fatal("Error when converting to int: ", err)
		}
		levels[i] = val
	}
	return levels
}

func evalSafe(levels []int) (bool, int) {
	prevValue := levels[0]
	dir := getDir(levels[0], levels[1])
	for i := 1; i < len(levels); i++ {
		if levels[i] == 0 { // Special case
			continue
		}
		if levels[i] == prevValue {
			return false, i
		}
		if dir == 1 && levels[i] > prevValue+3 {
			return false, i
		}
		if dir == 1 && levels[i] < prevValue {
			return false, i
		}
		if dir == -1 && levels[i] < prevValue-3 {
			return false, i
		}
		if dir == -1 && levels[i] > prevValue {
			return false, i
		}
		prevValue = levels[i]
	}
	return true, -1
}

func main() {
	filename := "input.txt"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	log.Printf("Reading file: %s\n", filename)
	content := ReadFile(filename)
	log.Println("Running part 1")
	ans1 := part1(content)
	log.Printf("Part1: %d\n", ans1)
	log.Println("Running part 2")
	ans2 := part2(content)
	log.Printf("Part2: %d\n", ans2)
}
