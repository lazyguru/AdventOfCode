package main

import "testing"

func TestPart1a(t *testing.T) {
	expectedAnswer1 := 4
	exampleData1 := []string{
		"-L|F7",
		"7S-7|",
		"L|7||",
		"-L-J|",
		"L|-JF",
	}
	ans := part1(exampleData1)
	if ans != expectedAnswer1 {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer1)
	}
}
func TestPart1b(t *testing.T) {
	expectedAnswer2 := 8
	exampleData2 := []string{
		"7.F7.",
		".FJ|7",
		"SJ.L7",
		"|F--J",
		"LJ.LJ",
	}
	ans := part1(exampleData2)
	if ans != expectedAnswer2 {
		t.Errorf("part1() = %d, want %d", ans, expectedAnswer2)
	}
}
