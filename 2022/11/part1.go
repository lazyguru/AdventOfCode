package main

import (
	"log"
)

func part1(data []string) int {
	answer := 0
	monkies := processInput(data, DEBUG)

	for round := 1; round < 21; round++ {
		for m := 0; m < len(monkies); m++ {
			monkey := monkies[m]
			monkey.Log("Monkey %d:", m)
			for _, v := range monkey.Items {
				v = monkey.Inspect(v)
				v = v / 3
				monkey.throwToNextMonkey(v, monkies)
				monkey.Items = nil
			}
		}
	}
	highest2 := []int{0, 0}
	for id := 0; id < len(monkies); id++ {
		m := monkies[id]
		m.Log("Monkey %d inspected items %d times.", id, m.OpCount)
		if m.OpCount > highest2[0] {
			highest2[1] = highest2[0]
			highest2[0] = m.OpCount
			continue
		}
		if m.OpCount > highest2[1] {
			highest2[1] = m.OpCount
			continue
		}
	}
	if DEBUG {
		log.Println(highest2)
	}
	answer = highest2[0] * highest2[1]
	return answer
}
