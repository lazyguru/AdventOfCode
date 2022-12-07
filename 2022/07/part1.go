package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
	answer := 0
	curDir := []string{"/"}
	dirs := make(map[string]int)
	for _, value := range data {
		if value == "" {
			continue
		}
		if strings.HasPrefix(value, "$ cd") {
			cmd := strings.Split(value, " ")
			if cmd[2] == ".." {
				curDir = curDir[:len(curDir)-1]
			} else if cmd[2] == "/" {
				curDir = []string{}
			} else {
				curDir = append(curDir, cmd[2])
			}
			continue
		}
		if value == "$ ls" {
			continue // nothing to really do here
		}
		if strings.HasPrefix(value, "dir") {
			continue // nothing to really do here
		}
		entry := strings.Split(value, " ")
		size, _ := strconv.Atoi(entry[0])
		updateTree(dirs, curDir, size)
	}
	for _, size := range dirs {
		if size <= 100000 {
			answer += size
		}
	}
	return answer
}
