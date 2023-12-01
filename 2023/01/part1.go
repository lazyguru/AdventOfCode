package main

import (
	"strconv"
)

func part1(data []string) int {
	answer := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		nums := []int{}
		for _, n := range value {
			v, err := strconv.Atoi(string(n))
			if err != nil {
				continue
			}
			nums = append(nums, v)
		}
		answer += (nums[0] * 10) + nums[len(nums)-1]
	}
	return answer
}
