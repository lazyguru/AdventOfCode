package main

import (
	"log"
	"strings"
)

func part2() {
	cnt := 0
	data := ReadFile()
	cols := 1000
	rows := 1000
	grid := make([]int, cols*rows)
	for _, value := range data {
		if value == "" {
			continue
		}
		points := strings.Split(value, " -> ")
		start := strings.Split(points[0], ",")
		end := strings.Split(points[1], ",")
		x1 := GetPoint(start[0])
		y1 := GetPoint(start[1])
		x2 := GetPoint(end[0])
		y2 := GetPoint(end[1])

		if x1 == x2 { // vertical line
			if y1 < y2 {
				for i := y1; i <= y2; i++ {
					grid[x1*cols+i]++
				}
			} else {
				for i := y2; i <= y1; i++ {
					grid[x1*cols+i]++
				}
			}
		} else if y1 == y2 { // horizontal line
			if x1 < x2 {
				for i := x1; i <= x2; i++ {
					grid[i*cols+y1]++
				}
			} else {
				for i := x2; i <= x1; i++ {
					grid[i*cols+y1]++
				}
			}
		} else { // diagonal
			i := x1
			xstep := 1
			if x1 > x2 {
				xstep = -1
			}
			n := y1
			ystep := 1
			if y1 > y2 {
				ystep = -1
			}
			for i != (x2 + xstep) {
				grid[i*cols+n]++
				i += xstep
				n += ystep
			}
		}
	}
	for _, value := range grid {
		if value > 1 {
			cnt++
		}
	}
	log.Printf("Answer: %d\n", cnt)
}
