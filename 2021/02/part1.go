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
	hpos := 0
	vpos := 0
	for _, value := range data {
		if value == "" {
			continue
		}
		cmd := strings.Split(value, " ")
		inc, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatal(err)
		}
		switch cmd[0] {
		case "forward":
			hpos += inc
		case "up":
			vpos -= inc
		case "down":
			vpos += inc
		}
	}
	log.Printf("Final count: %d\n", hpos*vpos)
}
