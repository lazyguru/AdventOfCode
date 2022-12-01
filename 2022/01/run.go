package main

import (
	"log"
	"os"
	"strings"
)

func ReadFile() []string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n\n")
}

func main() {
	log.Println("Running part 1")
	part1()
	log.Println("Running part 2")
	part2()
}
