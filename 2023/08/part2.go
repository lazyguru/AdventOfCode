package main

import (
	"sort"
	"strings"
)

var answers []int

func part2(data []string) int {
	answers = []int{}
	instructions := data[0]
	var currOps []string
	m := GetMap(data[2:])
	for key, _ := range m {
		if strings.HasSuffix(key, "A") {
			currOps = append(currOps, key)
		}
	}
	for _, currOp := range currOps {
		var answer int
		for !checkEndsWithZ([]string{currOp}) {
			var newOps []string
			answer, newOps = processData(instructions, m, []string{currOp}, answer)
			currOp = newOps[0]
		}
		answers = append(answers, answer)
	}
	return findAnswer()
}
func processData(instructions string, m map[string][]string, currOps []string, answer int) (int, []string) {
	for _, op := range instructions {
		if checkEndsWithZ(currOps) {
			break
		}
		answer++
		for key, currOp := range currOps {
			currOps[key] = GetNextOp(op, m, currOp)
		}
	}
	return answer, currOps
}

func findAnswer() int {
	if len(answers) == 1 {
		return answers[0]
	}
	sort.Ints(answers)
	lcm := 1
	for idx := 0; idx < len(answers); idx++ {
		a2 := answers[idx]
		lcm = (lcm * a2) / gcd(lcm, a2)
	}
	return lcm
}
func gcd(a1 int, a2 int) int {
	for a2%a1 != 0 {
		a3 := a1
		a1 = a2 % a1
		a2 = a3
	}
	return a1
}

func checkEndsWithZ(ops []string) bool {
	for _, op := range ops {
		if !strings.HasSuffix(op, "Z") {
			return false
		}
	}
	return true
}
