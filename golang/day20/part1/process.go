package part1

import (
	"fmt"
)

func Process(data []int) []int {
	next := make([]*int, len(data))
	for i := range data {
		next[i] = &data[i]
	}

	for i := range data {
		for j := range next {
			if next[j] == &data[i] {
				next = Move(next, j)
				break
			}
		}
	}

	out := make([]int, len(next))
	for i := range out {
		out[i] = *next[i]
	}

	return out
}

func Move(arr []*int, index int) []*int {
	poi := arr[index]
	distance := *poi % len(arr)
	if distance == 0 {
		// Move zero distance, nothing to do
		return arr
	} else if distance < 0 {
		distance = len(arr) + distance - 1
	}

	dest := (index + distance) % len(arr)
	out := make([]*int, 0, len(arr))
	for i, entry := range arr {
		if i == dest {
			out = append(out, entry, poi)
		} else if i == index {
			continue
		} else {
			out = append(out, entry)
		}
	}

	return out
}

func Print(arr []*int) {
	fmt.Printf("[ ")
	for _, entry := range arr {
		fmt.Printf("%4d ", *entry)
	}
	fmt.Printf("]\n")
}

func FindZero(list []int) int {
	for i := range list {
		if list[i] == 0 {
			return i
		}
	}

	return -1
}
