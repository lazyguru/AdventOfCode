package main

func part1(data []string) int {
	answer := 0
	markerA := data[0]
	markerB := data[1]
	markerC := data[2]
	markerD := data[3]
	for i := 4; i < len(data); i++ {
		if data[i] == "" {
			continue
		}
		if markerA == markerB ||
			markerA == markerC ||
			markerA == markerD ||
			markerB == markerC ||
			markerB == markerD ||
			markerC == markerD {
			markerA = markerB
			markerB = markerC
			markerC = markerD
			markerD = data[i]
		} else {
			return i
		}
	}
	return answer
}
