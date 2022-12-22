package part1

import (
	"fmt"
)

func Process(data []int) []int {
	// Initialize index array
	index := make([]int, len(data))
	for i := range data {
		index[i] = i
	}

	for i := range data {
		for j := range index {
			if index[j] == i {
				index = Remove(index, j)
				dest := (j + data[i]) % len(index)
				if dest < 0 {
					dest = len(index) + dest
				}
				index = Insert(index, dest, i)
				break
			}
		}
	}

	out := make([]int, len(data))
	for i := range index {
		out[i] = data[index[i]]
	}

	return out
}

func Remove(arr []int, i int) []int {
	if i == (len(arr) - 1) {
		return arr[:i]
	} else {
		return append(arr[:i], arr[i+1:]...)
	}
}
func Insert(arr []int, i int, val int) []int {
	if i == (len(arr)) {
		return append(arr, val)
	} else {
		out := make([]int, len(arr)+1)
		copy(out, arr)
		out = append(out[:i], val)
		out = append(out, arr[i:]...)
		return out
	}
}

func Print(data []int, order []int) {
	fmt.Printf("[ ")
	for _, entry := range order {
		fmt.Printf("%4d ", data[entry])
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
