package main

import (
	"fmt"
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

	fmt.Println(out)
}
