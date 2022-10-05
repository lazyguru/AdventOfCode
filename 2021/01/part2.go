package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	data := strings.Split(string(content), "\n")
	cnt := 0
	depths := convertToInts(data)
	for idx, _ := range depths {
		if (idx + 3) >= len(depths) {
			break
		}
		if depths[idx+3] > depths[idx] {
			cnt++
		}
	}
	log.Printf("Final count: %d\n", cnt)
}

func convertToInts(data []string) []int {
	var depths []int
	for _, value := range data {
		if value == "" {
			continue
		}
		depth, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Error converting %s to int", value)
		}
		depths = append(depths, depth)
	}
	return depths
}
