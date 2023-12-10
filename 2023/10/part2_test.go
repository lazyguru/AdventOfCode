package main

import "testing"

func TestPart2a(t *testing.T) {
	expectedAnswer1 := 8
	exampleData1 := []string{
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}
	ans := part1(exampleData1)
	if ans != expectedAnswer1 {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer1)
	}
}
func TestPart2b(t *testing.T) {
	expectedAnswer2 := 10
	exampleData2 := []string{
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJIF7FJ-",
		"L---JF-JLJIIIIFJLJJ7",
		"|F|F-JF---7IIIL7L|7|",
		"|FFJF7L7F-JF7IIL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}
	ans := part1(exampleData2)
	if ans != expectedAnswer2 {
		t.Errorf("part2() = %d, want %d", ans, expectedAnswer2)
	}
}
