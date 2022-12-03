package main

func part2(data []string) int {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		answer++
	}
	return answer
}
