package main

import "testing"

func TestPart1(t *testing.T) {
	expectedAnswer := 6440
	exampleData := []string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	ans := part1(exampleData)
	if ans != expectedAnswer {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer)
	}
}
