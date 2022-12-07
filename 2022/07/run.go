package main

import (
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

func updateTree(dirs map[string]int, curDir []string, size int) {
	wd := ""
	for _, dir := range curDir {
		wd += "/" + dir
		dirs[wd] += size
	}
	dirs["/"] += size
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
