package main

import (
	"day22/part2"
	"os"
	"strings"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	// board, steps := part2.ParseInput(fileName)
	board, _ := part2.ParseInput(fileName)
	board.PrintOrientation()

	//
	// process
	//

	// final := part2.Process(board, steps)

	//
	// Report
	//

	//score := board.Score(final)
	// fmt.Println(score)
}
