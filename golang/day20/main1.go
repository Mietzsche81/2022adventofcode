package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"day20/part1"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	data := part1.ParseInput(fileName)

	//
	// process
	//

	out := part1.Process(data)

	//
	// Report
	//

	zero := part1.FindZero(out)
	fmt.Println(out, zero)
	if zero < 0 {
		log.Fatal("could not find zero in output.")
	}
	score := 0
	for _, i := range []int{1000, 2000, 3000} {
		index := (zero + i) % len(out)
		fmt.Println(zero, zero+i, index, out[index])
		score += out[index]
	}
	fmt.Println(score)
}
