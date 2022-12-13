package main

func part1(data []string) int {
	myMap := buildMap(data)
	var start = &Vertex{Visited: true}
	var end = &Vertex{}
	for y := 0; y < len(myMap.Vertices); y++ {
		for x := 0; x < len(myMap.Vertices[y]); x++ {
			b := myMap.Vertices[y][x]
			if b.Letter == 'S' { // find the starting point
				start = b
				start.Visited = true
				break
			}
			if b.Letter == 'E' { // find the end point
				end = b
				break
			}
		}
	}
	queue := Queue{}
	queue.Add(start)
	for queue.Size() > 0 {
		v := queue.Dequeue()
		if v.Letter == 'E' {
			return v.Move
		}
		adj := myMap.FindAdjacent(v)
		for idx := 0; idx < len(adj); idx++ {
			adj[idx].Visited = true
			adj[idx].Move = v.Move + 1
			queue.Add(adj[idx])
		}
	}
	return end.Move
}
