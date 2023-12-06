package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type MyMap struct {
	Vertices [][]*Vertex
}

type Vertex struct {
	X      int
	Y      int
	Letter string
	Move   int
}

type Num struct {
	Row      int
	ColStart int
	ColEnd   int
	Value    string
	Symbol   string
}

func convertToInt(r string) int {
	i, err := strconv.Atoi(r)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func IsDigit(r string) bool {
	_, err := strconv.Atoi(r)
	if err != nil {
		return false
	}
	return true
}

func findNumber(n *Vertex, m [][]*Vertex, r string) Num {
	num := Num{
		Row:      n.Y,
		ColStart: n.X,
		ColEnd:   n.X,
		Value:    n.Letter,
		Symbol:   r,
	}
	// Check left
	for i := 1; i <= n.X; i++ {
		if IsDigit(m[n.Y][n.X-i].Letter) {
			num.Value = m[n.Y][n.X-i].Letter + num.Value
			num.ColStart--
			continue
		}
		break
	}
	// Check right
	for i := 1; i < len(m[n.Y])-n.X; i++ {
		if IsDigit(m[n.Y][n.X+i].Letter) {
			num.Value += m[n.Y][n.X+i].Letter
			num.ColEnd++
			continue
		}
		break
	}
	return num
}

func (m *MyMap) FindAdjacent(v *Vertex) []Num {
	var adj []Num
	if v.Y < len(m.Vertices)-1 { // Can we go down?
		n := m.Vertices[v.Y+1][v.X]
		if IsDigit(n.Letter) { // If this matches, then left and right don't matter as they will be part of this number
			adj = append(adj, findNumber(n, m.Vertices, v.Letter))
		} else { // It's possible for the below 2 to match and be separate numbers
			if v.X < (len(m.Vertices[v.Y]) - 1) { // Can we go down & right?
				n := m.Vertices[v.Y+1][v.X+1]
				if IsDigit(n.Letter) {
					adj = append(adj, findNumber(n, m.Vertices, v.Letter))
				}
			}
			if v.X > 0 { // Can we go down & left?
				n := m.Vertices[v.Y+1][v.X-1]
				if IsDigit(n.Letter) {
					adj = append(adj, findNumber(n, m.Vertices, v.Letter))
				}
			}
		}
	}
	if v.Y > 0 { // Can we go up?
		n := m.Vertices[v.Y-1][v.X]
		if IsDigit(n.Letter) { // If this matches, then left and right don't matter as they will be part of this number
			adj = append(adj, findNumber(n, m.Vertices, v.Letter))
		} else { // It's possible for the below 2 to match and be separate numbers
			if v.X < (len(m.Vertices[v.Y-1]) - 1) { // Can we go up & right?
				n := m.Vertices[v.Y-1][v.X+1]
				if IsDigit(n.Letter) {
					adj = append(adj, findNumber(n, m.Vertices, v.Letter))
				}
			}
			if v.X > 0 { // Can we go up & left?
				n := m.Vertices[v.Y-1][v.X-1]
				if IsDigit(n.Letter) {
					adj = append(adj, findNumber(n, m.Vertices, v.Letter))
				}
			}
		}
	}
	if v.X > 0 { // Can we go left?
		n := m.Vertices[v.Y][v.X-1]
		if IsDigit(n.Letter) {
			adj = append(adj, findNumber(n, m.Vertices, v.Letter))
		}
	}
	if v.X < len(m.Vertices[v.Y]) { // Can we go right?
		n := m.Vertices[v.Y][v.X+1]
		if IsDigit(n.Letter) {
			adj = append(adj, findNumber(n, m.Vertices, v.Letter))
		}
	}
	return adj
}

func buildMap(data []string) MyMap {
	myMap := MyMap{
		Vertices: make([][]*Vertex, 0, len(data)),
	}
	for y, value := range data {
		if value == "" {
			continue
		}
		vertices := make([]*Vertex, 0, len(value))
		for x, val := range value {
			v := &Vertex{Letter: string(val), Y: y, X: x}
			vertices = append(vertices, v)
		}
		myMap.Vertices = append(myMap.Vertices, vertices)
	}
	return myMap
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
