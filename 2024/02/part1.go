package main

func part1(data []string) int {
	answer := 0
	for report, value := range data {
		report += 1
		if value == "" {
			continue
		}
		levels := getLevels(value)
		safe, _ := evalSafe(levels)
		if safe {
			answer++
		}
	}
	return answer
}
