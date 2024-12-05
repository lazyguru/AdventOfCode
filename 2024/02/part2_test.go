package main

import "testing"

func TestPart2(t *testing.T) {
	expectedAnswer := 4
	exampleData := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}
	ans := part2(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}
