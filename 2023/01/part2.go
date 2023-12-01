package main

import (
	"sort"
	"strings"
)

func part2(data []string) int {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		nums := make(map[int]int)
		snums := []string{
			"zero",
			"one",
			"two",
			"three",
			"four",
			"five",
			"six",
			"seven",
			"eight",
			"nine",
			"0",
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
		}
		for idx, r := range snums {
			fpos := strings.Index(value, r)
			lpos := strings.LastIndex(value, r)
			if fpos != -1 {
				nums[fpos] = idx
				if idx > 10 {
					nums[fpos] = idx - 10
				}
			}
			if lpos != -1 {
				nums[lpos] = idx
				if idx > 10 {
					nums[lpos] = idx - 10
				}
			}
		}
		ikeys := []int{}
		for key, _ := range nums {
			ikeys = append(ikeys, key)
		}
		sort.Ints(ikeys)
		first, _ := nums[ikeys[0]]
		last, _ := nums[ikeys[len(ikeys)-1]]
		val := first*10 + last
		answer += val
	}
	return answer
}
