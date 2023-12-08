package main

import "testing"

func TestPart1Example1(t *testing.T) {
	expectedAnswer := 2
	ans := part1(exampleData1)
	if ans != expectedAnswer {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer)
	}
}

func TestPart1Example2(t *testing.T) {
	expectedAnswer := 6
	ans := part1(exampleData2)
	if ans != expectedAnswer {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer)
	}
}
