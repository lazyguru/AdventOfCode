package main

import (
	"log"
	"math"
	"os"
	"strings"
)

type Vect struct {
	Dir      string
	Distance int
}
type Cords struct {
	X int
	Y int
}

func plotLine(v Vect, curCords *Cords) []Cords {
	var path []Cords
	distance := v.Distance
	Xstep := 0
	Ystep := 0
	switch v.Dir {
	case "D":
		Ystep = -1
		break
	case "U":
		Ystep = 1
		break
	case "L":
		Xstep = -1
		break
	case "R":
		Xstep = 1
		break
	}
	for i := distance; i > 0; i-- {
		curCords.X += Xstep
		curCords.Y += Ystep
		path = append(path, *curCords)
	}
	return path
}

func (c *Cords) ChaseHead(head Cords) bool {
	x := head.X - c.X
	y := head.Y - c.Y
	if isTouching(x, y) {
		return false
	}
	if x == 0 && math.Abs(float64(y)) > 1 {
		if y < 0 {
			c.Y += y + 1
		} else {
			c.Y += y - 1
		}
		return true
	}
	if y == 0 && math.Abs(float64(x)) > 1 {
		if x < 0 {
			c.X += x + 1
		} else {
			c.X += x - 1
		}
		return true
	}
	if math.Abs(float64(x)) > 1 && math.Abs(float64(y)) == 1 {
		c.Y += y
		if x < 0 {
			c.X += x + 1
		} else {
			c.X += x - 1
		}
		return true
	}
	if math.Abs(float64(x)) == 1 && math.Abs(float64(y)) > 1 {
		c.X += x
		if y < 0 {
			c.Y += y + 1
		} else {
			c.Y += y - 1
		}
		return true
	}
	if math.Abs(float64(x)) > 1 && math.Abs(float64(y)) > 1 {
		if x < 0 {
			c.X += x + 1
		} else {
			c.X += x - 1
		}
		if y < 0 {
			c.Y += y + 1
		} else {
			c.Y += y - 1
		}
		return true
	}
	return false
}

func isTouching(a int, b int) bool {
	x := int(math.Abs(float64(a)))
	y := int(math.Abs(float64(b)))
	return x == 0 && y == 0 ||
		x == 1 && y == 0 ||
		x == 0 && y == 1 ||
		x == 1 && y == 1
}

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

func main() {
	filename := "input.txt"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	log.Printf("Reading file: %s\n", filename)
	content := ReadFile(filename)
	log.Println("Running part 1")
	ans1 := part1(content)
	log.Printf("Part1: %d\n", ans1)
	log.Println("Running part 2")
	ans2 := part2(content)
	log.Printf("Part2: %d\n", ans2)

}
