package main

import "testing"

func TestPart1(t *testing.T) {
	expectedAnswer := 4361
	exampleData := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	ans := part1(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer)
	}
}
