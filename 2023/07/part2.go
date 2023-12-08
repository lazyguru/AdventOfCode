package main

import (
	"sort"
	"strconv"
	"strings"
)

func part2(data []string) int {
	hand := func(p1, p2 *Game) bool {
		if p1.Value == p2.Value {
			for idx := 0; idx < 5; idx++ {
				if p1.Hand[idx] == p2.Hand[idx] {
					continue
				}
				p1v := GetCardValueP2(string(p1.Hand[idx]))
				p2v := GetCardValueP2(string(p2.Hand[idx]))
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
			Value: CalcValueP2(cards),
		}
		games = append(games, g)
	}
	By(hand).Sort(games)
	for idx := 0; idx < len(games); idx++ {
		answer += games[idx].Bid * (idx + 1)
	}
	return answer
}
func GetCardValueP2(card string) int {
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J": // Joker
		return 1
	case "T":
		return 10
	}
	val, _ := strconv.Atoi(card)
	return val
}

func CalcValueP2(cards string) float64 {
	jokers := strings.Count(cards, "J")
	if jokers == 5 {
		return FiveOfKind
	}
	var scores []float64
	for _, card := range cards {
		if string(card) == "J" {
			continue // We handle jokers separately
		}
		if (jokers + strings.Count(cards, string(card))) > 4 {
			return FiveOfKind
		}
		if (jokers + strings.Count(cards, string(card))) > 3 {
			return FourOfKind
		}
		if (jokers + strings.Count(cards, string(card))) > 2 {
			tmpCards := strings.Replace(cards, "J", "", jokers)
			if CheckFullHouseP2(tmpCards, string(card), strings.Count(cards, string(card))) {
				scores = append(scores, FullHouse)
				continue
			}
			scores = append(scores, ThreeOfKind)
			continue
		}
		if (jokers + strings.Count(cards, string(card))) > 1 {
			tmpCards := strings.Replace(cards, "J", "", jokers)
			if CheckFullHouseP2(tmpCards, string(card), strings.Count(cards, string(card))) {
				scores = append(scores, FullHouse)
				continue
			}
			if CheckTwoPairP2(tmpCards, string(card)) {
				scores = append(scores, TwoPair)
				continue
			}
			scores = append(scores, OnePair)
			continue
		}
	}
	if len(scores) == 0 {
		return HighCard
	}
	sort.Float64s(scores)
	return scores[len(scores)-1]
}

func CheckTwoPairP2(cards string, card string) bool {
	tmp := strings.Replace(cards, card, "", 2)
	if tmp[0] == tmp[1] || tmp[0] == tmp[2] || tmp[1] == tmp[2] {
		return true
	}
	return false
}

func CheckFullHouseP2(cards string, card string, count int) bool {
	tmp := strings.Replace(cards, card, "", count)
	if count == 3 && tmp[0] == tmp[1] {
		return true
	}
	if len(tmp) == 3 && tmp[0] == tmp[1] && tmp[0] == tmp[2] {
		return true
	}
	// Joker + card?
	if len(tmp) == 2 && tmp[0] == tmp[1] {
		return true
	}
	return false
}
