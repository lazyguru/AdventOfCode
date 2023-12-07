package main

import (
	"strconv"
	"strings"
)

func part2(data []string) int {
	t, _ := strings.CutPrefix(data[0], "Time:")
	d, _ := strings.CutPrefix(data[1], "Distance:")
	// dist = numSec * (totSec - numSec)
	ti, _ := strconv.Atoi(strings.Replace(t, " ", "", -1))
	di, _ := strconv.Atoi(strings.Replace(d, " ", "", -1))
	scenarios := GetScenarios(ti, di)
	return len(scenarios)
}
