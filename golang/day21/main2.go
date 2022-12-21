package main

import (
	"day21/part2"
	"fmt"
	"os"
	"strings"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])

	//
	// process
	//

	humn := part2.GradientDescent(fileName)

	//
	// Report
	//

	fmt.Println(humn)
}
