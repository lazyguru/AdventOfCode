package main

import (
	"log"
	"strconv"
	"strings"
)

type Slot struct {
	Row   int
	Col   int
	Num   int64
	Found bool
}

type Row struct {
	Slots []Slot
}

type Board struct {
	Id       int
	Rows     []Row
	HasBingo bool
}

func addUnmarked(board Board) int64 {
	var value int64
	for _, row := range board.Rows {
		value += addRow(row)
	}
	return value
}

func addRow(row Row) int64 {
	var value int64
	for _, val := range row.Slots {
		if !val.Found {
			value += val.Num
		}
	}
	return value
}

func buildBoards(data []string) []Board {
	var boards []Board
	currBoard := Board{
		Id:   0,
		Rows: make([]Row, 5),
	}
	currRow := 0
	for _, value := range data[2:] {
		if len(strings.TrimSpace(value)) == 0 || currRow > 4 {
			boards = append(boards, currBoard)
			currBoard = Board{
				Id:   len(boards),
				Rows: make([]Row, 5),
			}
			currRow = 0
			continue
		}
		currBoard.Rows[currRow] = buildMap(value, currRow)
		currRow++
	}
	if currRow > 0 {
		boards = append(boards, currBoard)
	}
	return boards
}

func buildMap(value string, row int) Row {
	newRow := Row{
		Slots: make([]Slot, 5),
	}
	nums := strings.Split(value, " ")
	col := 0
	for _, val := range nums {
		if len(val) == 0 {
			continue
		}
		num, err := strconv.ParseInt(val, 0, 32)
		if err != nil {
			log.Fatal("Error parsing row num: ", err)
		}
		newRow.Slots[col] = Slot{
			Num:   num,
			Found: false,
			Col:   col,
			Row:   row,
		}
		col++
	}
	return newRow
}

func (b *Board) checkNumber(n int64) bool {
	for idxRow, rows := range b.Rows {
		for idxSlot, slot := range rows.Slots {
			if slot.Num == n {
				slot.Found = true
				b.Rows[idxRow].Slots[idxSlot] = slot
				return true
			}
		}
	}
	return false
}

func (b *Board) checkBingo() bool {
	for _, row := range b.Rows {
		if checkRow(row) {
			return true
		}
	}
	for _, col := range []int{0, 1, 2, 3, 4} {
		if b.checkCol(col) {
			return true
		}
	}
	return false
}

func checkRow(row Row) bool {
	for _, val := range row.Slots {
		if !val.Found {
			return false
		}
	}
	return true
}

func (b Board) checkCol(col int) bool {
	for _, row := range b.Rows {
		found := getColValue(row, col)
		if !found {
			return false
		}
	}
	return true
}

func getColValue(row Row, col int) bool {
	for _, val := range row.Slots {
		if val.Col != col {
			continue
		}
		return val.Found
	}
	log.Println("Somehow didn't find column")
	return false
}
