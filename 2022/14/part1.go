package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Cords struct {
	X int
	Y int
}

func (c *Cords) String() string {
	return fmt.Sprintf("%d, %d", c.X, c.Y)
}

func part1(data []string) int {
	answer := 0
	rocks := []*Cords{}
	for _, value := range data {
		if value == "" {
			continue
		}
		c := strings.Split(value, " -> ")
		cords := []*Cords{}
		for _, v := range c {
			xy := strings.Split(v, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			cords = append(cords, &Cords{X: x, Y: y})
		}
		path := getPath(cords)
		rocks = append(rocks, path...)
	}
	return answer
}

func getPath(cords []*Cords) []*Cords {
	path := []*Cords{}
	var lastCord *Cords
	for idx := 0; idx < len(cords); idx++ {
		c := cords[idx]
		if lastCord == nil {
			lastCord = c
		}

		dist := getDist(c, lastCord)
		for dist > 0 {
			xStep := 0
			yStep := 0
			if lastCord.X > c.X {
				xStep = -1
			} else if lastCord.X < c.X {
				xStep = 1
			}
			if lastCord.Y > c.Y {
				yStep = -1
			} else if lastCord.Y < c.Y {
				yStep = 1
			}
			newCord := &Cords{X: lastCord.X + xStep, Y: lastCord.Y + yStep}
			path = append(path, newCord)
			dist = getDist(newCord, lastCord)
			lastCord = newCord
		}
		lastCord = c
		path = append(path, c)
	}
	log.Println(path)
	return path
}

func getDist(a *Cords, b *Cords) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}
