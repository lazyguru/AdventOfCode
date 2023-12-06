package main

func part2(data []string) int {
	myMap := buildMap(data)
	answer := 0
	for y := 0; y < len(myMap.Vertices); y++ {
		for x := 0; x < len(myMap.Vertices[y]); x++ {
			b := myMap.Vertices[y][x]
			if b.Letter == "*" {
				adj := myMap.FindAdjacent(b)
				if len(adj) < 2 {
					continue
				}
				i := 1
				for _, a := range adj {
					i *= convertToInt(a.Value)
				}
				answer += i
			}
		}
	}
	return answer
}
