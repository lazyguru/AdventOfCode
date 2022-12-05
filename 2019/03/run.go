package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cords struct {
	X int
	Y int
}

type Vect struct {
	Dir string
	Distance int
}

func getDistance(point Cords) int {
	return int(math.Abs(float64(point.X)) + math.Abs(float64(point.Y)))
}

func parseVect(value string) Vect {
	direction := value[0:1]
	distance, _ := strconv.Atoi(value[1:])
	return Vect{direction, distance}
}

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

func processLine(second []string) []Cords {
	var path []Cords
	c := Cords{0, 0} // reset origin
	for _, value := range second {
		if value == "" {
			continue
		}
		v := parseVect(value)
		path = append(path, plotLine(v, c)...)
		c = path[len(path)-1] // Grab last point
	}
	return path
}

func plotLine(v Vect, c Cords) []Cords {
	var path []Cords
	curCords := Cords{c.X, c.Y}
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
		if c.X == 0 && c.Y == 0 && i == 0{
			continue
		}
		curCords.X+=Xstep
		curCords.Y+=Ystep
		path = append(path, curCords)
	}
	return path
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
