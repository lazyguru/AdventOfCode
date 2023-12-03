package main

import "testing"

func TestPart1(t *testing.T) {
	expectedAnswer := 1
	exampleData := []string{}
	ans := part1(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}
