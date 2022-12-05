package main

import (
	"math"
	"strconv"
	"strings"
)

func part2(data []string) string {
	answer := ""
	stacks := make(map[int][]string)
	for _, value := range data {
		if strings.Contains(value, " 1   2   3 ") {
			break
		}
		cols := strings.Split(value, "")
		cnt := int(math.Round(float64(len(value))/4.0) + 1.0)
		for i := 1; i < cnt; i++ {
			pos := i*4 - 3
			if cols[pos] != " " {
				stacks[i] = append(stacks[i], cols[pos])
			}
		}
	}
	for idx, stack := range stacks {
		stacks[idx] = doReverse(stack)
	}
	for _, value := range data {
		if !strings.Contains(value, "move") {
			continue
		}
		tokens := strings.Split(value, " ")
		count, _ := strconv.Atoi(tokens[1])
		from, _ := strconv.Atoi(tokens[3])
		to, _ := strconv.Atoi(tokens[5])
		stackA := stacks[from]
		stackB := stacks[to]
		pos := len(stackA) - count
		item := stackA[pos : pos+count]
		stackB = append(stackB, item...)
		stackA = stackA[:pos]
		stacks[from] = stackA
		stacks[to] = stackB
	}
	for i := 0; i < len(stacks); i++ {
		idx := i + 1
		stack := stacks[idx]
		answer += stack[len(stack)-1]
	}
	return answer
}
