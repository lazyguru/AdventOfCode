package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
*/

const PIPE_NS = '|'
const PIPE_EW = '-'
const PIPE_90_NE = 'L'
const PIPE_90_NW = 'J'
const PIPE_90_SW = '7'
const PIPE_90_SE = 'F'
const GROUND = '.'
const START = 'S'

type MyMap struct {
	Vertices [][]*Vertex
}

type Vertex struct {
	Visited bool
	X       int
	Y       int
	Letter  rune
	Move    int
}

func (v *Vertex) String() string {
	visited := "Not visited"
	if v.Visited {
		visited = "Visited"
	}
	return fmt.Sprintf("| %d,%d | %s | %s(%d) |", v.X, v.Y, visited, string(v.Letter), v.Move)
}

func (m *MyMap) FindAdjacent(v *Vertex) []*Vertex {
	var adj []*Vertex
	if v.Y < len(m.Vertices)-1 { // Can we go down?
		n := m.Vertices[v.Y+1][v.X]
		if isValid(n, v) {
			adj = append(adj, n)
		}
	}
	if v.Y > 0 { // Can we go up?
		n := m.Vertices[v.Y-1][v.X]
		if isValid(n, v) {
			adj = append(adj, n)
		}
	}
	if v.X > 0 { // Can we go left?
		n := m.Vertices[v.Y][v.X-1]
		if isValid(n, v) {
			adj = append(adj, n)
		}
	}
	if v.X < len(m.Vertices[v.Y])-1 { // Can we go right?
		n := m.Vertices[v.Y][v.X+1]
		if isValid(n, v) {
			adj = append(adj, n)
		}
	}
	return adj
}

func (m *MyMap) reset() {
	for _, row := range m.Vertices {
		for _, v := range row {
			v.Move = 0
			v.Visited = false
		}
	}
}

func isValid(n *Vertex, v *Vertex) bool {
	if n.Visited || n.Letter == GROUND {
		return false
	}
	switch v.Letter {
	case START:
		if n.X > v.X && (n.Letter == PIPE_EW || n.Letter == PIPE_90_NW || n.Letter == PIPE_90_SW) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.X < v.X && (n.Letter == PIPE_EW || n.Letter == PIPE_90_NE || n.Letter == PIPE_90_SE) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.Y < v.Y && (n.Letter == PIPE_NS || n.Letter == PIPE_90_SE || n.Letter == PIPE_90_SW) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.Y > v.Y && (n.Letter == PIPE_NS || n.Letter == PIPE_90_NE || n.Letter == PIPE_90_NW) {
			return true //!n.Visited || n.Move > v.Move
		}
		return false
	case GROUND:
		return false
	case PIPE_90_NE:
		if n.X > v.X && (n.Letter == PIPE_EW || n.Letter == PIPE_90_SW || n.Letter == PIPE_90_NW) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.Y < v.Y && (n.Letter == PIPE_NS || n.Letter == PIPE_90_SE || n.Letter == PIPE_90_SW) {
			return true //!n.Visited || n.Move > v.Move
		}
		return false
	case PIPE_90_NW:
		if n.X < v.X && (n.Letter == PIPE_EW || n.Letter == PIPE_90_SE || n.Letter == PIPE_90_NE) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.Y < v.Y && (n.Letter == PIPE_NS || n.Letter == PIPE_90_SE || n.Letter == PIPE_90_SW) {
			return true //!n.Visited || n.Move > v.Move
		}
		return false
	case PIPE_90_SE:
		if n.X > v.X && (n.Letter == PIPE_EW || n.Letter == PIPE_90_SW || n.Letter == PIPE_90_NW) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.Y > v.Y && (n.Letter == PIPE_NS || n.Letter == PIPE_90_NE || n.Letter == PIPE_90_NW) {
			return true //!n.Visited || n.Move > v.Move
		}
		return false
	case PIPE_90_SW:
		if n.X < v.X && (n.Letter == PIPE_EW || n.Letter == PIPE_90_SE || n.Letter == PIPE_90_NE) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.Y > v.Y && (n.Letter == PIPE_NS || n.Letter == PIPE_90_NE || n.Letter == PIPE_90_NW) {
			return true //!n.Visited || n.Move > v.Move
		}
		return false
	case PIPE_EW:
		if n.X > v.X && (n.Letter == PIPE_EW || n.Letter == PIPE_90_SW || n.Letter == PIPE_90_NW) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.X < v.X && (n.Letter == PIPE_EW || n.Letter == PIPE_90_SE || n.Letter == PIPE_90_NE) {
			return true //!n.Visited || n.Move > v.Move
		}
		return false
	case PIPE_NS:
		if n.Y < v.Y && (n.Letter == PIPE_NS || n.Letter == PIPE_90_SW || n.Letter == PIPE_90_SE) {
			return true //!n.Visited || n.Move > v.Move
		}
		if n.Y > v.Y && (n.Letter == PIPE_NS || n.Letter == PIPE_90_NW || n.Letter == PIPE_90_NE) {
			return true //!n.Visited || n.Move > v.Move
		}
		return false
	}
	return false
}

func buildMap(data []string) *MyMap {
	myMap := MyMap{}
	for y, value := range data {
		if value == "" {
			continue
		}
		char := []rune(value)
		vertices := make([]*Vertex, 0)
		for x := 0; x < len(char); x++ {
			v := &Vertex{Letter: char[x], Y: y, X: x}
			vertices = append(vertices, v)
		}
		myMap.Vertices = append(myMap.Vertices, vertices)
	}
	return &myMap
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
