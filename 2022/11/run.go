package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

type Operation struct {
	Operand   string
	Parameter int
	UseOld    bool
}

type Monkey struct {
	Items     []int
	Operation Operation
	Test      int
	NextTrue  int
	NextFalse int
	OpCount   int
	Debug     bool
}

func (monkey *Monkey) AddItem(item int) {
	monkey.Items = append(monkey.Items, item)
}

func (monkey *Monkey) throwToNextMonkey(v int, monkies map[int]*Monkey) {
	monkey.Log("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", v)
	if v%monkey.Test == 0 {
		monkey.Log("    Current worry level is divisible by %d.", monkey.Test)
		monkey.Log("    Item with worry level %d is thrown to monkey %d.", v, monkey.NextTrue)
		nextMonkey := monkies[monkey.NextTrue]
		nextMonkey.AddItem(v)
	} else {
		monkey.Log("    Current worry level is not divisible by %d.", monkey.Test)
		monkey.Log("    Item with worry level %d is thrown to monkey %d.", v, monkey.NextFalse)
		nextMonkey := monkies[monkey.NextFalse]
		nextMonkey.AddItem(v)
	}
}

func processInput(data []string, debug bool) map[int]*Monkey {
	monkies := make(map[int]*Monkey)
	m := 0
	for idx := 0; idx < len(data); {
		value := data[idx]
		if value == "" {
			continue
		}
		itemLine := strings.Split(data[idx+1], ":")[1]
		itemStr := strings.Split(itemLine, ",")
		monkey := Monkey{Debug: debug}
		for x := range itemStr {
			item, _ := strconv.Atoi(strings.TrimSpace(itemStr[x]))
			monkey.AddItem(item)
		}
		monkey.Operation = parseOp(strings.Split(data[idx+2], ":")[1])
		t := strings.Split(data[idx+3], "by")[1]
		monkey.Test, _ = strconv.Atoi(strings.TrimSpace(t))
		nt := strings.Split(data[idx+4], "monkey")[1]
		monkey.NextTrue, _ = strconv.Atoi(strings.TrimSpace(nt))
		nf := strings.Split(data[idx+5], "monkey")[1]
		monkey.NextFalse, _ = strconv.Atoi(strings.TrimSpace(nf))
		monkies[m] = &monkey
		m++
		idx += 7
	}
	return monkies
}

func (monkey *Monkey) Inspect(v int) int {
	monkey.OpCount++
	monkey.Log("  Monkey inspects an item with a worry level of %d.\n", v)
	switch monkey.Operation.Operand {
	case "+":
		if monkey.Operation.UseOld {
			monkey.Log("    Worry level increases by %d to %d.\n", v, v+v)
			v = v + v
		} else {
			monkey.Log("    Worry level increases by %d to %d.\n", monkey.Operation.Parameter, v+monkey.Operation.Parameter)
			v = v + monkey.Operation.Parameter
		}
		break
	case "*":
		if monkey.Operation.UseOld {
			monkey.Log("    Worry level is multiplied by itself to %d.\n", v*v)
			v = v * v
		} else {
			monkey.Log("    Worry level is multiplied by %d to %d.\n", monkey.Operation.Parameter, v*monkey.Operation.Parameter)
			v = v * monkey.Operation.Parameter
		}
		break
	}
	return v
}

func (monkey *Monkey) InspectItemWithDynamicRelief(v int) int {
	monkey.OpCount++
	monkey.Log("  Monkey inspects an item with a worry level of %d.\n", v)
	switch monkey.Operation.Operand {
	case "+":
		if monkey.Operation.UseOld {
			monkey.Log("    Worry level increases by %d to %d.\n", v, v+v)
			v = v + v
		} else {
			monkey.Log("    Worry level increases by %d to %d.\n", monkey.Operation.Parameter, v+monkey.Operation.Parameter)
			v = v + monkey.Operation.Parameter
		}
		break
	case "*":
		if monkey.Operation.UseOld {
			monkey.Log("    Worry level is multiplied by itself to %d.\n", v*v)
			v = v * v
		} else {
			monkey.Log("    Worry level is multiplied by %d to %d.\n", monkey.Operation.Parameter, v*monkey.Operation.Parameter)
			v = v * monkey.Operation.Parameter
		}
		break
	}
	return v / 3
}

func (monkey *Monkey) Log(format string, v ...any) {
	if monkey.Debug {
		log.Printf(format, v)
	}
}

func parseOp(operation string) Operation {
	op := Operation{
		Operand: "+",
		UseOld:  false,
	}
	if strings.Contains(operation, "+") {
		op.Operand = "+"
	}
	if strings.Contains(operation, "*") {
		op.Operand = "*"
	}
	if operation == fmt.Sprintf(" new = old %s old", op.Operand) {
		op.UseOld = true
		return op
	}
	p := strings.TrimSpace(strings.Split(operation, op.Operand)[1])
	op.Parameter, _ = strconv.Atoi(p)
	return op
}

func ReadFile(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Error when opening input file: ", err)
	}
	return strings.Split(string(content), "\n")
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
