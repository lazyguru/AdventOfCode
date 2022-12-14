package main

import (
	"encoding/json"
	"strings"
)

type SortedSlice struct {
	Data []string
}

func (sorted *SortedSlice) Insert(idx int, val string) {
	if len(sorted.Data) == idx {
		sorted.Data = append(sorted.Data, val)
	}
	sorted.Data = append(sorted.Data[:idx+1], sorted.Data[idx:]...)
	sorted.Data[idx] = val
}

func (sorted *SortedSlice) Append(val string) {
	sorted.Data = append(sorted.Data, val)
}

func (sorted *SortedSlice) First() string {
	if len(sorted.Data) == 0 {
		return ""
	}
	return sorted.Data[0]
}

func (sorted *SortedSlice) Size() int {
	return len(sorted.Data)
}

func (sorted *SortedSlice) Get(i int) string {
	return sorted.Data[i]
}

func (sorted *SortedSlice) Find(s string) int {
	for idx := 0; idx < len(sorted.Data); idx++ {
		if sorted.Data[idx] == s {
			return idx
		}
	}
	return -1
}

func (sorted *SortedSlice) String() string {
	return strings.Join(sorted.Data, "\n")
}

func part2(data []string) int {
	var sorted SortedSlice
	sorted.Append("[[2]]")
	sorted.Append("[[6]]")
	for _, value := range data {
		if value == "" {
			continue
		}
		var list1 []any
		var list2 []any
		_ = json.Unmarshal([]byte(value), &list1)
		for i := 0; i < sorted.Size(); i++ {
			_ = json.Unmarshal([]byte(sorted.Get(i)), &list2)
			state := process(list1, list2)
			if state == -1 && i == (sorted.Size()-1) {
				sorted.Append(value)
				break
			} else if state == 1 {
				sorted.Insert(i, value)
				break
			} else if state == 0 {
				sorted.Insert(i, value)
				break
			}
		}
	}
	beginIdx := sorted.Find("[[2]]") + 1
	endIdx := sorted.Find("[[6]]") + 1
	return beginIdx * endIdx
}
