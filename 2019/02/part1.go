package main

import (
	"log"
	"strconv"
	"strings"
)

func part1(data []string) int {
	answer := 0
	input := strings.Split(data[0], ",")
	input[1] = "12"
	input[2] = "2"
	for idx := 0; idx < len(input); {
		if input[idx] == "" {
			continue
		}
		switch input[idx] {
		case "1":
			pos1, _ := strconv.Atoi(input[idx+1])
			val1, _ := strconv.Atoi(input[pos1])
			pos2, _ := strconv.Atoi(input[idx+2])
			val2, _ := strconv.Atoi(input[pos2])
			pos3, _ := strconv.Atoi(input[idx+3])
			input[pos3] = strconv.Itoa(val1 + val2)
			idx += 4
			break
		case "2":
			pos1, _ := strconv.Atoi(input[idx+1])
			val1, _ := strconv.Atoi(input[pos1])
			pos2, _ := strconv.Atoi(input[idx+2])
			val2, _ := strconv.Atoi(input[pos2])
			pos3, _ := strconv.Atoi(input[idx+3])
			input[pos3] = strconv.Itoa(val1 * val2)
			idx += 4
			break
		case "99":
			// stop
			answer, _ = strconv.Atoi(input[0])
			return answer
		default:
			// error
			log.Fatalf("Error occured due to opcode: %s at position %d", input[idx], idx)
		}
	}
	return answer
}
