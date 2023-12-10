package main

func part1(data []string) int {
	myMap := buildMap(data)
	var start = &Vertex{Visited: true}
	for y := 0; y < len(myMap.Vertices); y++ {
		for x := 0; x < len(myMap.Vertices[y]); x++ {
			b := myMap.Vertices[y][x]
			if b.Letter == START { // find the starting point
				start = b
				start.Visited = true
				break
			}
		}
	}
	queue := Queue{}
	queue.Add(start)
	highestMoves := 0
	for queue.Size() > 0 {
		v := queue.Dequeue()
		adj := myMap.FindAdjacent(v)
		for idx := 0; idx < len(adj); idx++ {
			adj[idx].Visited = true
			adj[idx].Move = v.Move + 1
			if highestMoves < adj[idx].Move {
				highestMoves = adj[idx].Move
			}
			queue.Add(adj[idx])
		}
	}
	return highestMoves
}
