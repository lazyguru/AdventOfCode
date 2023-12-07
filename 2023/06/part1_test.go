package main

import "testing"

func TestPart1(t *testing.T) {
	expectedAnswer := 288
	exampleData := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	ans := part1(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer)
	}
}
