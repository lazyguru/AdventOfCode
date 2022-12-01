package main

import (
	"log"
)

func part1(data []string) {
	cnt := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		cnt++
	}
	log.Printf("Final count: %d\n", cnt)
}
