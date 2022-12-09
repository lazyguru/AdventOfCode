package main

import (
	"strconv"
	"strings"
)

func part2(data []string) int {
	answer := 0
	var vectors []Vect
	for _, value := range data {
		if value == "" {
			continue
		}
		movement := strings.Split(value, " ")
		dist, _ := strconv.Atoi(movement[1])
		vect := Vect{
			movement[0],
			dist,
		}
		vectors = append(vectors, vect)
	}
	headPos := Cords{0, 0}
	segmentPos := make([]Cords, 9)
	headCords := []Cords{}
	for _, v := range vectors {
		headCords = append(headCords, plotLine(v, &headPos)...)
	}
	tailCords := followLeader(headCords, segmentPos)
	answer = len(tailCords) + 1
	return answer
}

func followLeader(headCords []Cords, segments []Cords) map[Cords]int {
	curHead := headCords
	tailCords := make(map[Cords]int)
	for _, segment := range segments {
		tailPath := []Cords{}
		tailCords = make(map[Cords]int)
		for _, cords := range curHead {
			if segment.ChaseHead(cords) {
				tailCords[segment]++
			}
			tailPath = append(tailPath, segment)
		}
		curHead = tailPath
	}
	return tailCords
}
