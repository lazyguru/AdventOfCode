package main

import "testing"

func TestPart2(t *testing.T) {
	expectedAnswer := 31
	exampleData := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	ans := part2(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}
