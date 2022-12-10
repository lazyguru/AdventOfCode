package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part2(data []string) string {
	x := 1
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
	spritePosition := strings.Split("###.....................................", "")
	output := ""
	fullOutput := ""
	pos := 0
	for i := 0; i <= lastOp; i++ {
		output += spritePosition[pos]
		//cycle := i + 1
		x += stack[i]
		pos++
		spritePosition = strings.Split("........................................", "")
		if x >= 0 && x < 40 {
			spritePosition[x] = "#"
		}
		if x >= -1 && x < 39 {
			spritePosition[x+1] = "#"
		}
		if x >= -2 && x < 38 {
			spritePosition[x+2] = "#"
		}
		if pos == 40 {
			fullOutput += fmt.Sprintln(output)
			output = ""
			pos = 0
		}
	}
	return fullOutput
}
