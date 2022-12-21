package main

import (
	"fmt"
	"os"
	"strings"

	"day21/part1"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	data := part1.ParseInput(fileName)
	for _, monkey := range data {
		fmt.Println(monkey.Str())
	}

	//
	// process
	//

	part1.Process(data)

	//
	// Report
	//

	for _, monkey := range data {
		fmt.Println(monkey.Str())
	}
	fmt.Printf("\n\n%s\n\n", data["root"].Str())
}
