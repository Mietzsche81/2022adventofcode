package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var Stockpile = map[string]int{
	"ore":      0,
	"clay":     0,
	"obsidian": 0,
	"geode":    0,
}

var Producers = map[string]int{
	"ore":      0,
	"clay":     0,
	"obsidian": 0,
	"geode":    0,
}

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
		line := scanner.Text()
		// Transform
		line = strings.TrimSpace(line)
		// Load
		data = append(data, line)
	}
	return
}

func ProcessData(data []string) []string {
	out := make([]string, len(data))

	// Process

	return out
}
