package main

import (
	"strings"
)

func part2(data []string) int {
	answer := 1
	if data[len(data)-1] == "" {
		data = data[:len(data)-1]
	}
	for y := 1; y < (len(data) - 1); y++ {
		trees := strings.Split(data[y], "")
		for x := 1; x < (len(trees) - 1); x++ {
			myAns := 1
			count := 0
			if count = countLeft(trees, x); count != 0 {
				myAns *= count
			}
			if count = countRight(trees, x); count != 0 {
				myAns *= count
			}
			if count = countAbove(trees, x, y, data); count != 0 {
				myAns *= count
			}
			if count = countBelow(trees, x, y, data); count != 0 {
				myAns *= count
			}
			if myAns > answer {
				answer = myAns
			}
		}
	}
	return answer
}

func countAbove(trees []string, x int, y int, data []string) int {
	treeCount := 0
	for i := y - 1; i >= 0; i-- {
		treeCount++
		treesAbove := strings.Split(data[i], "")
		if treesAbove[x] >= trees[x] {
			break
		}
	}
	return treeCount
}

func countBelow(trees []string, x int, y int, data []string) int {
	treeCount := 0
	for i := y + 1; i < len(data); i++ {
		treeCount++
		treesBelow := strings.Split(data[i], "")
		if treesBelow[x] >= trees[x] {
			break
		}
	}
	return treeCount
}

func countRight(trees []string, x int) int {
	treeCount := 0
	for y := x + 1; y < len(trees); y++ {
		treeCount++
		if trees[y] >= trees[x] {
			break
		}
	}
	return treeCount
}

func countLeft(trees []string, x int) int {
	treeCount := 0
	for y := x - 1; y >= 0; y-- {
		treeCount++
		if trees[y] >= trees[x] {
			break
		}
	}
	return treeCount
}
