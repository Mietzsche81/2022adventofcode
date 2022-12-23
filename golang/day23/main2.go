package main

import (
	"fmt"
	"os"
	"strings"

	"day23/part2"
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

	_, settlingTime := part2.Process(data)

	//
	// Report
	//

	fmt.Println(settlingTime)
}
