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
	lastDepth := 0
	cnt := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		depth, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Error converting %s to int", value)
		}
		if depth == 0 {
			continue
		}
		if depth > lastDepth && lastDepth != 0 {
			cnt++
		}
		lastDepth = depth
	}
	log.Printf("Final count: %d\n", cnt)
}
