package main

func part1(data []string) int {
	myMap := buildMap(data)
	var nums []Num
	answer := 0
	for y := 0; y < len(myMap.Vertices); y++ {
		for x := 0; x < len(myMap.Vertices[y]); x++ {
			b := myMap.Vertices[y][x]
			if !IsDigit(b.Letter) && b.Letter != "." {
				adj := myMap.FindAdjacent(b)
				nums = append(nums, adj...)
			}
		}
	}
	for _, v := range nums {
		answer += convertToInt(v.Value)
	}
	return answer
}
