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

func GetPoint(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	return i
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
