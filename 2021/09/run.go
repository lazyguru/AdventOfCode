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

func getCols(cols []string) []int {
	icols := make([]int, len(cols))
	for idx, col := range cols {
		icol, _ := strconv.Atoi(col)
		icols[idx] = icol
	}
	return icols
}

func checkDepth(b int, cols []int, a int, rows [][]int) bool {
	if b > 0 && cols[b] >= cols[b-1] {
		return false
	}
	if b+1 < len(cols) && cols[b] >= cols[b+1] {
		return false
	}
	if a > 0 && cols[b] >= rows[a-1][b] {
		return false
	}
	if a+1 < len(rows) && cols[b] >= rows[a+1][b] {
		return false
	}
	return true
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
