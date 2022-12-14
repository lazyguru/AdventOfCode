package main

import (
	"log"
	"os"
	"reflect"
	"strings"
)

func process(row1 []any, row2 []any) int {
	if len(row1) == 0 && len(row2) > 0 {
		return 1
	}
	if len(row2) == 0 && len(row1) > 0 {
		return -1
	}
	for idx := 0; idx < len(row1); idx++ {
		if idx == len(row2) {
			return -1
		}
		row1Type := reflect.TypeOf(row1[idx])
		row2Type := reflect.TypeOf(row2[idx])
		if row1Type.Kind() == reflect.Float64 && row1Type.Kind() == row2Type.Kind() { // both are numbers
			val1 := row1[idx].(float64)
			val2 := row2[idx].(float64)
			if val1 < val2 {
				return 1
			} else if val1 > val2 {
				return -1
			}
		} else if row1Type.Kind() == row2Type.Kind() { // both are same type and not strings, must be slices
			val := process(row1[idx].([]any), row2[idx].([]any))
			if val != 0 {
				return val
			}
		} else { // one is numbers, the other is a slice
			var r1 []any
			var r2 []any
			if row1Type.Kind() == reflect.Float64 {
				r1 = []any{row1[idx]}
				r2 = row2[idx].([]any)
			} else {
				r1 = row1[idx].([]any)
				r2 = []any{row2[idx]}
			}
			val := process(r1, r2)
			if val != 0 {
				return val
			}
		}
		if len(row1) < len(row2) {
			return 1
		}
	}
	return 0
}

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
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
