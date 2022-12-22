package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"day20/part2"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	data := part2.ParseInput(fileName)

	//
	// process
	//

	out := part2.Process(data)

	//
	// Report
	//

	zero := part2.FindZero(out)
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
