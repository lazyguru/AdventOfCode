package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

type MyMap struct {
	Vertices [][]*Vertex
}

type Vertex struct {
	Visited bool
	X       int
	Y       int
	Letter  rune
	Height  int
	Move    int
}

func (v *Vertex) String() string {
	visited := "Not visited"
	if v.Visited {
		visited = "Visited"
	}
	return fmt.Sprintf("| %d,%d | %s | %s(%d) |", v.X, v.Y, visited, string(v.Letter), v.Height)
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
	diff := n.Height - v.Height
	return !n.Visited && diff <= 1
}

func buildMap(data []string) MyMap {
	myMap := MyMap{}
	for y, value := range data {
		if value == "" {
			continue
		}
		char := []rune(value)
		vertices := make([]*Vertex, 0)
		for x := 0; x < len(char); x++ {
			height := int(char[x]) - int('a')
			if char[x] == 'S' {
				height = 0
			}
			if char[x] == 'E' {
				height = int('z') - int('a')
			}
			v := &Vertex{Height: height, Letter: char[x], Y: y, X: x}
			vertices = append(vertices, v)
		}
		myMap.Vertices = append(myMap.Vertices, vertices)
	}
	return myMap
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
