package main

import "testing"

func TestPart2(t *testing.T) {
	expectedAnswer := 71503
	exampleData := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	ans := part2(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}
