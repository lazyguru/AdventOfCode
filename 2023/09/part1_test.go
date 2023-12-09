package main

import "testing"

func TestPart1(t *testing.T) {
	expectedAnswer := 114
	exampleData := []string{
		"0 3 6 9 12 15",
		"1 3 6 10 15 21",
		"10 13 16 21 30 45",
	}
	ans := part1(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer)
	}
}
