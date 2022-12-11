package main

import (
	"log"
)

func part2(data []string) int {
	answer := 0
	monkies := processInput(data, DEBUG)
	mod := 1
	for m := 0; m < len(monkies); m++ {
		mod *= monkies[m].Test
	}
	for round := 1; round < 10001; round++ {
		for m := 0; m < len(monkies); m++ {
			monkey := monkies[m]
			for _, v := range monkey.Items {
				v = monkey.Inspect(v)
				v = v % mod
				monkey.throwToNextMonkey(v, monkies)
				monkey.Items = nil
			}
		}
		if round == 1 || round == 20 || round%1000 == 0 {
			if DEBUG {
				log.Printf("== After round %d ==\n", round)
			}
			checkHighest(monkies)
		}
	}
	if DEBUG {
		log.Printf("== After all rounds ==\n")
	}
	highest1, highest2 := checkHighest(monkies)
	if DEBUG {
		log.Println(highest1, highest2)
	}
	answer = highest1 * highest2
	return answer
}

func checkHighest(monkies map[int]*Monkey) (int, int) {
	highest1 := 0
	highest2 := 0
	for id := 0; id < len(monkies); id++ {
		m := monkies[id]
		m.Log("Monkey %d inspected items %d times.", id, m.OpCount)
		if m.OpCount > highest1 {
			highest2 = highest1
			highest1 = m.OpCount
			continue
		}
		if m.OpCount > highest2 {
			highest2 = m.OpCount
			continue
		}
	}
	return highest1, highest2
}
