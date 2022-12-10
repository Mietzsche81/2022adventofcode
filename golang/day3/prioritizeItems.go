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
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	scanner := bufio.NewScanner(fin)

	//
	// Process data
	//
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		doubles := FindDouble(line)
		// Get rid of duplicates
		results := make(map[rune]int)
		for _, char := range doubles {
			results[char] = RtoI(char)
		}
		// Add to sum
		for _, val := range results {
			total += val
		}
	}

	//
	// Report Output
	//
	fmt.Println(total)
}

func RtoI(r rune) int {
	if (r-'a' < 26) && (r-'a' >= 0) {
		return int(r - 'a' + 1)
	} else {
		return int(r - 'A' + 27)
	}
}

func FindDouble(line string) (doubles []rune) {
	var center int = len(line) / 2
	for i, char := range line[:center] {
		if strings.Contains(line[center:], line[i:i+1]) {
			doubles = append(doubles, char)
		}
	}
	return doubles
}
