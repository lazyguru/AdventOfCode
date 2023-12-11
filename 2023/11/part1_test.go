package main

import "testing"

func TestPart1(t *testing.T) {
	expectedAnswer := 374
	exampleData := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	ans := part1(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer)
	}
}
