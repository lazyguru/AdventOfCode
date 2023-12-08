package main

import "testing"

func TestPart2Example1(t *testing.T) {
	expectedAnswer := 2
	ans := part2(exampleData1)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}

func TestPart2Example2(t *testing.T) {
	expectedAnswer := 6
	ans := part2(exampleData2)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}

func TestPart2Example3(t *testing.T) {
	expectedAnswer := 6
	ans := part2(exampleData3)
	if ans != expectedAnswer {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer)
	}
}
