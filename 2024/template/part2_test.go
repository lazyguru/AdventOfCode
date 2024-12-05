package main

import "testing"

func TestPart2(t *testing.T) {
	expectedAnswer := 2
	exampleData := []string{}
	ans := part2(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}
