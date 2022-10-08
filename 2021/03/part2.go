package main

import (
	"log"
	"os"
	"strconv"
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
		cnt0 := 0
		for _, value := range oxygen {
			bits := strings.Split(value, "")
			if bits[idx] == "0" {
				cnt0++
			}
		}
		cnt1 := len(oxygen) - cnt0
		common := "0"
		if cnt0 < cnt1 {
			common = "1"
		}
		if cnt0 == cnt1 {
			common = "1"
		}
		newoxygen := []string{}
		for _, value := range oxygen {
			bits := strings.Split(value, "")
			if bits[idx] == common {
				newoxygen = append(newoxygen, value)
			}
		}
		oxygen = newoxygen
		idx++
	}
	log.Println(oxygen)
	idx = 0
	for idx < cols && len(co2) > 1 {
		cnt0 := 0
		for _, value := range co2 {
			bits := strings.Split(value, "")
			if bits[idx] == "0" {
				cnt0++
			}
		}
		cnt1 := len(co2) - cnt0
		common := "0"
		if cnt0 > cnt1 { // Looking for least common
			common = "1"
		}
		if cnt0 == cnt1 {
			common = "0"
		}
		newco2 := []string{}
		for _, value := range co2 {
			bits := strings.Split(value, "")
			if bits[idx] == common {
				newco2 = append(newco2, value)
			}
		}
		co2 = newco2
		idx++
	}
	log.Println(co2)
	co2rate, err := strconv.ParseInt(co2[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	oxygenrate, err := strconv.ParseInt(oxygen[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Final count: %d\n", co2rate*oxygenrate)
}
