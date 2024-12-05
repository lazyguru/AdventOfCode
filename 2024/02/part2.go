package main

import (
	"slices"
)

func part2(data []string) int {
	answer := 0
	for report, value := range data {
		report += 1
		if value == "" {
			continue
		}
		levels := getLevels(value)
		idx := 0
		safe := false
		for !safe && idx < len(levels) {
			newlevels := slices.Clone(levels)
			newlevels = slices.Delete(newlevels, idx, idx+1)
			safe, _ = evalSafe(newlevels)
			idx++
		}
		if !safe {
			continue
		}
		answer++
	}
	return answer
}
