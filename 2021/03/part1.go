package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	data := strings.Split(string(content), "\n")
	cols := len(data[0])
	gamma := 0
	epsilon := 0
	idx := 0
	for idx < cols {
		cnt0 := 0
		for _, value := range data {
			bits := strings.Split(value, "")
			if bits[idx] == "0" {
				cnt0++
			}
		}
		cnt1 := len(data) - cnt0
		gamma *= 2
		epsilon *= 2
		if cnt0 < cnt1 {
			gamma += 1
		} else {
			epsilon += 1
		}
		idx++
	}
	log.Printf("Final count: %d\n", gamma*epsilon)
}
