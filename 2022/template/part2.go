package main

import (
	"log"
)

func part2(data []string) {
	cnt := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		cnt++
	}
	log.Printf("Final count: %d\n", cnt)
}
