package main

import "testing"

func TestPart2(t *testing.T) {
	expectedAnswer := 467835
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
	ans := part2(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}
