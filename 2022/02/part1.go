package main

import (
	"log"
	"strings"
)

/*
A = Rock
B = Paper
C = Scissors

X = Rock
Y = Paper
Z = Scissors
*/

func part1(data []string) {
	score := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		plays := strings.Split(value, " ")
		switch plays[0] {
		case "A":
			if plays[1] == "Y" {
				score += 2 + 6
			} else if plays[1] == "Z" {
				score += 3 + 0
			} else if plays[1] == "X" {
				score += 1 + 3
			}
		case "B":
			if plays[1] == "Y" {
				score += 2 + 3
			} else if plays[1] == "Z" {
				score += 3 + 6
			} else if plays[1] == "X" {
				score += 1 + 0
			}
		case "C":
			if plays[1] == "Y" {
				score += 2 + 0
			} else if plays[1] == "Z" {
				score += 3 + 3
			} else if plays[1] == "X" {
				score += 1 + 6
			}
		}

	}
	log.Printf("Answer: %d\n", score)
}
