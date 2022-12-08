package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
}

func intCodeProgram(input []string, programInput string) int {
	answer := -1
	for idx := 0; idx < len(input); {
		if input[idx] == "" {
			continue
		}
		instructions := getInstructions(input[idx])
		switch instructions[3] {
		case "01": // Addition
			pos1, _ := strconv.Atoi(input[idx+1])
			val1 := pos1
			if instructions[2] == "0" {
				val1, _ = strconv.Atoi(input[pos1])
			}
			pos2, _ := strconv.Atoi(input[idx+2])
			val2 := pos2
			if instructions[1] == "0" {
				val2, _ = strconv.Atoi(input[pos2])
			}
			pos3, _ := strconv.Atoi(input[idx+3])
			input[pos3] = strconv.Itoa(val1 + val2)
			idx += 4
			break
		case "02": // Multiply
			pos1, _ := strconv.Atoi(input[idx+1])
			val1 := pos1
			if instructions[2] == "0" {
				val1, _ = strconv.Atoi(input[pos1])
			}
			pos2, _ := strconv.Atoi(input[idx+2])
			val2 := pos2
			if instructions[1] == "0" {
				val2, _ = strconv.Atoi(input[pos2])
			}
			pos3, _ := strconv.Atoi(input[idx+3])
			input[pos3] = strconv.Itoa(val1 * val2)
			idx += 4
			break
		case "03": // Write input to position
			pos1, _ := strconv.Atoi(input[idx+1])
			input[pos1] = programInput
			idx += 2
			break
		case "04": // Output
			pos1, _ := strconv.Atoi(input[idx+1])
			answer, _ = strconv.Atoi(input[pos1])
			idx += 2
			break
		case "05": // jump-if-true
			pos1, _ := strconv.Atoi(input[idx+1])
			val1 := pos1
			if instructions[2] == "0" {
				val1, _ = strconv.Atoi(input[pos1])
			}
			if val1 != 0 {
				pos2, _ := strconv.Atoi(input[idx+2])
				val2 := pos2
				if instructions[1] == "0" {
					val2, _ = strconv.Atoi(input[pos2])
				}
				idx = val2
			} else {
				idx += 3
			}
			break
		case "06": // jump-if-false
			pos1, _ := strconv.Atoi(input[idx+1])
			val1 := pos1
			if instructions[2] == "0" {
				val1, _ = strconv.Atoi(input[pos1])
			}
			if val1 == 0 {
				pos2, _ := strconv.Atoi(input[idx+2])
				val2 := pos2
				if instructions[1] == "0" {
					val2, _ = strconv.Atoi(input[pos2])
				}
				idx = val2
			} else {
				idx += 3
			}
			break
		case "07": // less-than
			pos1, _ := strconv.Atoi(input[idx+1])
			val1 := pos1
			if instructions[2] == "0" {
				val1, _ = strconv.Atoi(input[pos1])
			}
			pos2, _ := strconv.Atoi(input[idx+2])
			val2 := pos2
			if instructions[1] == "0" {
				val2, _ = strconv.Atoi(input[pos2])
			}
			pos3, _ := strconv.Atoi(input[idx+3])
			if val1 < val2 {
				input[pos3] = "1"
			} else {
				input[pos3] = "0"
			}
			idx += 4
			break
		case "08": // equals
			pos1, _ := strconv.Atoi(input[idx+1])
			val1 := pos1
			if instructions[2] == "0" {
				val1, _ = strconv.Atoi(input[pos1])
			}
			pos2, _ := strconv.Atoi(input[idx+2])
			val2 := pos2
			if instructions[1] == "0" {
				val2, _ = strconv.Atoi(input[pos2])
			}
			pos3, _ := strconv.Atoi(input[idx+3])
			if val1 == val2 {
				input[pos3] = "1"
			} else {
				input[pos3] = "0"
			}
			idx += 4
			break
		case "99":
			// stop
			return answer
		default:
			// error
			log.Fatalf("Error occured due to opcode: %s at position %d", input[idx], idx)
		}
	}
	return answer
}

func getInstructions(input string) []string {
	for len(input) < 5 {
		input = "0" + input
	}
	parts := strings.Split(input, "")
	return []string{
		parts[0],
		parts[1],
		parts[2],
		parts[3] + parts[4],
	}
}

func main() {
	filename := "input.txt"
	if len(os.Args) == 2 {
		filename = os.Args[1]
	}
	log.Printf("Reading file: %s\n", filename)
	content := ReadFile(filename)
	log.Println("Running part 1")
	ans1 := part1(content)
	log.Printf("Part1: %d\n", ans1)
	log.Println("Running part 2")
	ans2 := part2(content)
	log.Printf("Part2: %d\n", ans2)

}
