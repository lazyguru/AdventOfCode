package main

import (
	"strings"
)

func part1(data []string) int {
	answer := 0
	height := len(data)
	if data[height-1] == "" {
		height--
		data = data[:height]
	}
	width := len(data[0]) - 2 // ignore corners as already counted in height
	// outside edges
	answer += width*2 + height*2
	for y := 1; y < (height - 1); y++ {
		trees := strings.Split(data[y], "")
		for x := 1; x < (len(trees) - 1); x++ {
			if checkLeft(trees, x) {
				answer++
				continue
			}
			if checkRight(trees, x) {
				answer++
				continue
			}
			if checkAbove(trees, x, y, data) {
				answer++
				continue
			}
			if checkBelow(trees, x, y, data) {
				answer++
				continue
			}
		}
	}
	return answer
}

func checkAbove(trees []string, x int, y int, data []string) bool {
	for i := 0; i < y; i++ {
		treesAbove := strings.Split(data[i], "")
		if treesAbove[x] >= trees[x] {
			return false
		}
	}
	return true
}

func checkBelow(trees []string, x int, y int, data []string) bool {
	for i := y + 1; i < len(data); i++ {
		treesBelow := strings.Split(data[i], "")
		if treesBelow[x] >= trees[x] {
			return false
		}
	}
	return true
}

func checkRight(trees []string, x int) bool {
	for y := x + 1; y < len(trees); y++ {
		if trees[y] >= trees[x] {
			return false
		}
	}
	return true
}

func checkLeft(trees []string, x int) bool {
	for y := 0; y < x; y++ {
		if trees[y] >= trees[x] {
			return false
		}
	}
	return true
}
