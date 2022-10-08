package main

import (
	"log"
	"strconv"
	"strings"
)

func convertBinToInt(value string) int {
	rate, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(rate)
}

func countBits(slice []string, idx int) (int, int) {
	cnt0 := 0
	for _, value := range slice {
		bits := strings.Split(value, "")
		if bits[idx] == "0" {
			cnt0++
		}
	}
	cnt1 := len(slice) - cnt0
	return cnt0, cnt1
}

func filterSlice(slice []string, idx int, common string) []string {
	newslice := []string{}
	for _, value := range slice {
		bits := strings.Split(value, "")
		if bits[idx] == common {
			newslice = append(newslice, value)
		}
	}
	return newslice
}
