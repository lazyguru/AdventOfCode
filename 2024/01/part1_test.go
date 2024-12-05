package main

import "testing"

func TestPart1(t *testing.T) {
	expectedAnswer := 11
	exampleData := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	ans := part1(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}
