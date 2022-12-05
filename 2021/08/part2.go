package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

func part2(data []string) int {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		line := strings.Split(value, " | ")
		signals := strings.Split(line[0], " ")
		digits := mapDigitsFromSignals(signals)
		outputs := line[1]
		output := strings.Split(outputs, " ")
		sval := ""
		log.Println(digits)
		for idx, digit := range output {
			digit = sortSignal(digit)
			output[idx] = digit
			if digits[digit] == 10 {
				sval += "0"
				continue
			}
			sval += strconv.Itoa(digits[digit])
		}
		log.Println(output, sval)
		val, _ := strconv.Atoi(sval)
		answer += val
	}
	return answer
}

func mapDigitsFromSignals(signals []string) map[string]int {
	digits := make(map[string]int)
	segments := make(map[int]string)
	for idx, signal := range signals {
		signal = sortSignal(signal)
		signals[idx] = signal
	}
	for _, signal := range signals {
		if signal == "" {
			continue
		}
		switch len(signal) {
		case 2:
			digits[signal] = 1
			segments[1] = signal
			continue
		case 3:
			digits[signal] = 7
			segments[7] = signal
			continue
		case 4:
			digits[signal] = 4
			segments[4] = signal
			continue
		case 7:
			digits[signal] = 8
			segments[8] = signal
			continue
		}
	}
	log.Println(digits, segments)
	diff := subtract(segments[4], segments[1])
	for _, signal := range signals {
		if signal == "" {
			continue
		}
		switch len(signal) {
		case 5:
			if contains(signal, diff) {
				digits[signal] = 5
				segments[5] = signal
				continue
			}
			if contains(signal, segments[1]) {
				digits[signal] = 3
				segments[3] = signal
				continue
			}
			digits[signal] = 2
			segments[2] = signal
			break
		case 6:
			if contains(signal, segments[4]) {
				digits[signal] = 9
				segments[9] = signal
				continue
			}
			if contains(signal, segments[7]) {
				digits[signal] = 10
				segments[10] = signal
				continue
			}
			digits[signal] = 6
			segments[6] = signal
			break
		}
	}
	return digits
}

func subtract(s1 string, s2 string) string {
	s := ""
	if s1 == "" || s2 == "" {
		return s1
	}
	for idx := 0; idx < len(s1); idx++ {
		if !strings.Contains(s2, string(s1[idx])) {
			s += string(s1[idx])
		}
	}
	return s
}

func sortSignal(signal string) string {
	s := strings.Split(signal, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func contains(signal string, search string) bool {
	if search == "" || signal == "" {
		return false
	}
	for idx := 0; idx < len(search); idx++ {
		if !strings.Contains(signal, string(search[idx])) {
			return false
		}
	}
	return true
}
