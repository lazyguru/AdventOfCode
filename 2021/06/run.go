package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile() []string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

func GetFish(data string) []int {
	sFish := strings.Split(data, ",")
	iFish := make([]int, len(sFish))
	for idx, value := range sFish {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		iFish[idx] = i
	}
	return iFish
}

func main() {
	if len(os.Args) == 1 {
		log.Println("Running part 1")
		part1()
		os.Exit(0)
	} else {
		log.Println("Running part 2")
		part2()
		os.Exit(0)
	}
}
