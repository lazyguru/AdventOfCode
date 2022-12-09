package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
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
	tailPos := Cords{0, 0}
	headCords := []Cords{}

	for _, v := range vectors {
		headCords = append(headCords, plotLine(v, &headPos)...)
	}
	tailCords := make(map[Cords]int)
	tailPath := []Cords{}
	for _, cords := range headCords {
		if tailPos.ChaseHead(cords) {
			tailCords[tailPos]++
		}
		tailPath = append(tailPath, tailPos)
	}
	answer = len(tailCords) + 1
	return answer
}
