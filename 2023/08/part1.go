package main

func part1(data []string) int {
	instructions := data[0]
	answer := 0
	currOp := "AAA"
	m := GetMap(data[2:])
	for currOp != "ZZZ" {
		for _, op := range instructions {
			if currOp == "ZZZ" {
				break
			}
			answer++
			currOp = GetNextOp(op, m, currOp)
		}
	}
	return answer
}
