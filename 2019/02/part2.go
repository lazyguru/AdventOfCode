package main

import (
	"log"
	"strconv"
	"strings"
)

func part2(data []string) int {
	answer := -1
	for idx1 := 0; idx1 < 100; idx1++ {
		for idx2 := 0; idx2 < 100; idx2++ {
			input := strings.Split(data[0], ",")
			input[1] = strconv.Itoa(idx1)
			input[2] = strconv.Itoa(idx2)
			answer = intCodeProgram(input)
			if answer != -1 {
				return answer
			}
		}
	}
	return answer
}

func intCodeProgram(input []string) int {
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
			if input[0] == "19690720" {
				noun, _ := strconv.Atoi(input[1])
				verb, _ := strconv.Atoi(input[2])
				return (100 * noun) + verb
			}
			return -1
		default:
			// error
			log.Fatalf("Error occured due to opcode: %s at position %d", input[idx], idx)
		}
	}
	return -1
}
