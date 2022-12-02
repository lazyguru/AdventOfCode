package download

import "fmt"

func getDay(day int) string {
	if day > 9 {
		return fmt.Sprintf("%d", day)
	}
	return fmt.Sprintf("0%d", day)
}

func downloadInput(year, day int) {
	sday := getDay(day)
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%s/input", year, sday)
}

func downloadPuzzle(year, day) string {
	sday := getDay(day)
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%s", year, sday)
	return ""
}

func SetupPuzzle(year, day, part int) {
	downloadInput(year, day)
	puzzle := downloadPuzzle(year, day)
}
