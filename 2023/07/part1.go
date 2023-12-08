package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
	// Closures that order the Game structure.
	hand := func(p1, p2 *Game) bool {
		if p1.Value == p2.Value {
			for idx := 0; idx < 5; idx++ {
				if p1.Hand[idx] == p2.Hand[idx] {
					continue
				}
				p1v := GetCardValueP1(string(p1.Hand[idx]))
				p2v := GetCardValueP1(string(p2.Hand[idx]))
				if p1v == p2v {
					continue
				}
				return p1v < p2v
			}
		}
		return p1.Value < p2.Value
	}
	answer := 0
	var games []Game
	for _, value := range data {
		if value == "" {
			continue
		}
		cards, bid, _ := strings.Cut(value, " ")
		bidAmount, _ := strconv.Atoi(bid)
		g := Game{
			Hand:  cards,
			Bid:   bidAmount,
			Value: CalcValueP1(cards),
		}
		games = append(games, g)
	}
	By(hand).Sort(games)
	for idx := 0; idx < len(games); idx++ {
		answer += games[idx].Bid * (idx + 1)
	}
	return answer
}

func GetCardValueP1(card string) int {
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	}
	val, _ := strconv.Atoi(card)
	return val
}

func CalcValueP1(cards string) float64 {
	for _, card := range cards {
		if strings.Count(cards, string(card)) > 4 {
			return FiveOfKind
		}
		if strings.Count(cards, string(card)) > 3 {
			return FourOfKind
		}
		if strings.Count(cards, string(card)) > 2 {
			if CheckFullHouseP1(cards, string(card), 3) {
				return FullHouse
			}
			return ThreeOfKind
		}
		if strings.Count(cards, string(card)) > 1 {
			if CheckFullHouseP1(cards, string(card), 2) {
				return FullHouse
			}
			if CheckTwoPairP1(cards, string(card)) {
				return TwoPair
			}
			return OnePair
		}
	}
	return HighCard
}

func CheckTwoPairP1(cards string, card string) bool {
	tmp := strings.Replace(cards, card, "", 2)
	if tmp[0] == tmp[1] || tmp[0] == tmp[2] || tmp[1] == tmp[2] {
		return true
	}
	return false
}

func CheckFullHouseP1(cards string, card string, count int) bool {
	tmp := strings.Replace(cards, card, "", count)
	if count == 3 && tmp[0] == tmp[1] {
		return true
	}
	if count == 2 && tmp[0] == tmp[1] && tmp[0] == tmp[2] {
		return true
	}
	return false
}
