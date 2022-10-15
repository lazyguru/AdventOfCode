package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func part2() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	data := strings.Split(string(content), "\n")
	nums := strings.Split(data[0], ",")
	boards := buildBoards(data)
	for _, num := range nums {
		iNum, err := strconv.ParseInt(num, 0, 32)
		if err != nil {
			log.Fatal("Error parsing num: ", err)
		}
		for idxBoard, board := range boards {
			if board.checkNumber(iNum) && !board.HasBingo && board.checkBingo() {
				if !stillPlayableBoards(boards) {
					unmarked := addUnmarked(board)
					log.Printf("Final count: %d\n", (unmarked * iNum))
					os.Exit(0)
				}
				boards[idxBoard].HasBingo = true
			}
		}
	}
}

func stillPlayableBoards(boards []Board) bool {
	playable := 0
	for _, board := range boards {
		if !board.HasBingo {
			playable++
		}
		if playable > 1 {
			return true
		}
	}
	log.Println(playable)
	return (playable > 1)
}
