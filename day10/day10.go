package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	data := ParseInput(fileName)

	//
	// process
	//

	out := ProcessData(data)

	//
	// report
	//

	fmt.Println(out)
}

func ParseInput(fileName string) (data []string) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	scanner := bufio.NewScanner(fin)
	data = make([]string, 0)
	for scanner.Scan() {
		// Extract
		nextLine := scanner.Text()
		// Transform

		// Load
		data = append(data, nextLine)
	}
	return
}

func ProcessData(data []string) []string {
	out := make([]string, 0)

	return out
}
