package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"day23/part1"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	iterations, err := strconv.Atoi(strings.TrimSpace(os.Args[2]))
	if err != nil {
		log.Fatal(err)
	}
	data := part1.ParseInput(fileName)

	//
	// process
	//

	out := part1.Process(data, iterations)

	//
	// Report
	//

	score := part1.Score(out)
	fmt.Println(score)
}
