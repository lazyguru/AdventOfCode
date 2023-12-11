package main

import (
	"slices"
	"strings"
)

func part1(data []string) int {
	myMap := make(map[int][]string)
	rowsExpanded := 0
	colsToExpand := []int{}
	for idx, value := range data {
		if value == "" {
			continue
		}
		m := strings.Split(value, "")
		if value == strings.Repeat(".", len(value)) {
			myMap[idx+rowsExpanded] = m
			rowsExpanded++
		} else {
			x := -1
			for cnt := strings.Count(value, "#"); cnt > 0; cnt-- {
				x = strings.Index(value[x+1:], "#")
				if !slices.Contains(colsToExpand, x) {
					colsToExpand = append(colsToExpand, x)
				}
			}
		}
		myMap[idx+rowsExpanded] = m
	}
	expandCols(myMap, colsToExpand)
	return findPaths(myMap)
}

func findPaths(myMap map[int][]string) int {
	return 0
}

func expandCols(myMap map[int][]string, colsToExpand []int) map[int][]string {
	for _, x := range colsToExpand {
		for y := 0; y < len(myMap); y++ {
			myMap[y] = slices.Insert(myMap[y], x, ".")
		}
	}
	return myMap
}
