package main

import (
	"math"
)

func part2(data []string) int {
	myMap := buildMap(data)
	var start []*Vertex
	for y := 0; y < len(myMap.Vertices); y++ {
		for x := 0; x < len(myMap.Vertices[y]); x++ {
			b := myMap.Vertices[y][x]
			if b.Letter == 'S' || b.Letter == 'a' { // find the starting point
				b.Letter = 'a' // overwrite 'S' with 'a'
				start = append(start, b)
				break
			}
		}
	}
	lowestMoves := math.MaxInt
	for _, s := range start {
		myMap.reset()
		queue := Queue{}
		queue.Add(s)
		for queue.Size() > 0 {
			v := queue.Dequeue()
			if v.Letter == 'E' {
				if v.Move < lowestMoves {
					lowestMoves = v.Move
				}
				break
			}
			adj := myMap.FindAdjacent(v)
			for idx := 0; idx < len(adj); idx++ {
				adj[idx].Visited = true
				adj[idx].Move = v.Move + 1
				queue.Add(adj[idx])
			}
		}
	}
	return lowestMoves
}
