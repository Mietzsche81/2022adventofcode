package main

import (
	"fmt"
	"os"
	"strings"

	"day22/part1"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	board, steps := part1.ParseInput(fileName)

	//
	// process
	//

	final := part1.Process(board, steps)

	//
	// Report
	//

	score := part1.Score(final)
	fmt.Println(score)
}
