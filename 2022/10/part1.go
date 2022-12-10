package main

import (
	"strconv"
	"strings"
)

func part1(data []string) int {
	x := 1
	signalStrength := 0
	stack := make(map[int]int, len(data))
	lastOp := 0
	for i := 0; i < len(data); i++ {
		value := data[i]
		if value == "" {
			continue
		}
		instructions := strings.Split(value, " ")
		instruction := instructions[0]
		if instruction == "addx" {
			nextAmount, _ := strconv.Atoi(instructions[1])
			lastOp += 2
			stack[lastOp] = nextAmount
		}
		if instruction == "noop" {
			lastOp++
		}
	}
	for i := 0; i <= lastOp; i++ {
		cycle := i + 1
		x += stack[i]
		if cycle == 20 || (cycle-20)%40 == 0 {
			signalStrength += x * cycle
		}
	}
	return signalStrength
}
