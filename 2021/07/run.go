package main

import (
	"log"
	"os"
	"sort"
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

func GetPos(data string) []float64 {
	str := strings.Split(data, ",")
	i := make([]float64, len(str))
	for idx, value := range str {
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}
		i[idx] = val
	}
	sort.Slice(i, func(a int, b int) bool {
		return i[a] < i[b]
	})
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
