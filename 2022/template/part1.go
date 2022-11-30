package main

import (
	"log"
)

func part1() {
	cnt := 0
	data := ReadFile()
	for _, value := range data {
		if value == "" {
			continue
		}
		cnt++
	}
	log.Printf("Final count: %d\n", cnt)
}
