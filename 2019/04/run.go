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

func checkIncreasing(i int) bool {
	sliceI := IntToSlice(i, []int{})
	for k := 1; k < len(sliceI); k++ {
		if sliceI[k] < sliceI[k-1] {
			return false
		}
	}
	return true
}

// IntToSlice Borrowed from: https://stackoverflow.com/a/58231817/564576
func IntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		// sequence = append(sequence, i) // reverse order output
		sequence = append([]int{i}, sequence...)
		return IntToSlice(n/10, sequence)
	}
	return sequence
}

func checkForRepeating(i int) bool {
	s := strconv.Itoa(i)
	for k := 1; k < len(s); k++ {
		if s[k] == s[k-1] {
			return true
		}
	}
	return false
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
