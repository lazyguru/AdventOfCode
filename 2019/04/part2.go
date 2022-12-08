package main

import (
	"strconv"
	"strings"
)

func part2(data []string) int {
	answer := 0
	passwordRange := strings.Split(data[0], "-")
	low, _ := strconv.Atoi(passwordRange[0])
	high, _ := strconv.Atoi(passwordRange[1])
	i := low
	for i <= high {
		if !checkForRepeating(i) {
			i++
			continue
		}
		if !checkIncreasing(i) {
			i++
			continue
		}
		if tooManyRepeats(i) {
			i++
			continue
		}
		i++
		answer++
	}
	return answer
}

func tooManyRepeats(i int) bool {
	sliceI := IntToSlice(i, []int{})
	digits := make(map[int]int)
	for k := 1; k < len(sliceI); k++ {
		if sliceI[k] == sliceI[k-1] {
			digits[sliceI[k]]++
		}
	}
	if len(digits) == 1 && digits[0] == 1 {
		return false // only one matching pair of length 2
	}
	for _, digit := range digits {
		if digit == 1 {
			return false
		}
	}
	return true
}
