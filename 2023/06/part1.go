package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
	answer := 1
	timeData, _ := strings.CutPrefix(data[0], "Time:")
	times := strings.Fields(timeData)
	distanceData, _ := strings.CutPrefix(data[1], "Distance:")
	distances := strings.Fields(distanceData)
	for idx := 0; idx < len(times); idx++ {
		// dist = numSec * (totSec - numSec)
		t, _ := strconv.Atoi(times[idx])
		d, _ := strconv.Atoi(distances[idx])
		scenarios := GetScenarios(t, d)
		answer *= len(scenarios)
	}
	return answer
}
