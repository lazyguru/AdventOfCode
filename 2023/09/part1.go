package main

import (
	"log"
	"strconv"
	"strings"
)

func part1(data []string) int {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		strs := strings.Split(value, " ")
		var nums []int
		for _, str := range strs {
			ivalue, _ := strconv.Atoi(str)
			nums = append(nums, ivalue)
		}
		answer += CalcDiff(nums)
	}
	return answer
}

func CalcDiff(nums []int) int {
	run := make(map[int][]int)
	diffCount := -1
	r := 0
	run[r] = nums
	for diffCount != 0 {
		r++
		diffCount = 0
		diffs := run[r-1]
		for i := 0; i < len(diffs)-1; i++ {
			diffCount += diffs[i+1] - diffs[i]
			run[r] = append(run[r], diffs[i+1]-diffs[i])
		}
	}
	for i := r; i >= 0; i-- {
		if i == 0 {
			log.Println("final", run, "value", run[i][len(run[i])-1])
			return run[i][len(run[i])-1] // New value should have already been added
		}
		if run[i][0] == 0 && run[i][len(run[i])-1] == 0 {
			run[i] = append(run[i], run[i][len(run[i])-1])
			run[i-1] = append(run[i-1], run[i-1][len(run[i-1])-1])
			continue
		}
		run[i-1] = append(run[i-1], run[i-1][len(run[i-1])-1]+run[i][len(run[i])-1])
	}
	log.Panic("should not get here")
	return 0
}
