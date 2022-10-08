package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	data := strings.Split(string(content), "\n")
	cols := len(data[0])
	co2 := data
	oxygen := data

	idx := 0
	for idx < cols && len(oxygen) > 1 {
		cnt0, cnt1 := countBits(oxygen, idx)
		common := "0"
		if cnt0 < cnt1 {
			common = "1"
		}
		if cnt0 == cnt1 {
			common = "1"
		}
		oxygen = filterSlice(oxygen, idx, common)
		idx++
	}
	log.Println(oxygen)
	idx = 0
	for idx < cols && len(co2) > 1 {
		cnt0, cnt1 := countBits(co2, idx)
		common := "0"
		if cnt0 > cnt1 { // Looking for least common
			common = "1"
		}
		if cnt0 == cnt1 {
			common = "0"
		}
		co2 = filterSlice(co2, idx, common)
		idx++
	}
	log.Println(co2)
	co2rate := convertBinToInt(co2[0])
	oxygenrate := convertBinToInt(oxygen[0])
	log.Printf("Final count: %d\n", co2rate*oxygenrate)
}
