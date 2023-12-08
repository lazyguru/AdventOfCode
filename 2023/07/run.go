package main

import (
	"log"
	"os"
	"sort"
	"strings"
)

/*
- 7 ** Five of a kind, where all five cards have the same label: AAAAA
- 6 ** Four of a kind, where four cards have the same label and one card has a different label: AA8AA
- 5 ** Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
- 4 ** Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
- 3 ** Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
- 2 ** One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
- 1 ** High card, where all cards' labels are distinct: 23456
*/

const FiveOfKind float64 = 5
const FourOfKind float64 = 4
const FullHouse float64 = 3.5
const ThreeOfKind float64 = 3
const TwoPair float64 = 2
const OnePair float64 = 1
const HighCard float64 = 0

type Game struct {
	Hand  string
	Value float64
	Bid   int
}

type By func(p1, p2 *Game) bool

func (by By) Sort(games []Game) {
	ps := &gameSorter{
		games: games,
		by:    by,
	}
	sort.Sort(ps)
}

type gameSorter struct {
	games []Game
	by    func(p1, p2 *Game) bool
}

func (s *gameSorter) Len() int {
	return len(s.games)
}

func (s *gameSorter) Swap(i, j int) {
	s.games[i], s.games[j] = s.games[j], s.games[i]
}

func (s *gameSorter) Less(i, j int) bool {
	return s.by(&s.games[i], &s.games[j])
}

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

func main() {
	filename := "input.txt"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	log.Printf("Reading file: %s\n", filename)
	content := ReadFile(filename)
	log.Println("Running part 1")
	ans1 := part1(content)
	log.Printf("Part1: %d\n", ans1)
	log.Println("Running part 2")
	ans2 := part2(content)
	log.Printf("Part2: %d\n", ans2)

}
