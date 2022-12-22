package part1

import "fmt"

func Process(data *[]Entry) []int {
	out := make([]*int, len(data))
	for i, val := range data {
		out[i] = &val
	}
	size := copy(out, data)

	for i := 0; i < size; i++ {
		move(&out, i)
	}

	return out
}

func move(arr *[]int, index int) {
	if (arr)
	distance := (*arr)[index]
	size := len(*arr)
	fmt.Println(size % distance)
}
