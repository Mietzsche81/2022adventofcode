package main

// Imports
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Main
func main() {

	// Read input

	var fileName string = strings.TrimSpace(os.Args[1])
	fmt.Printf("Processing input from: '%s'\n", fileName)
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line

	scanner := bufio.NewScanner(fin)
	var mostCalories int = 0
	var currentCalories int = 0
	for scanner.Scan() {
		// Get next line
		nextLine := scanner.Text()
		// If line is delimiter (empty line)
		if strings.TrimSpace(nextLine) == "" {
			// Compare against most calories
			if currentCalories > mostCalories {
				// Mark new max and reset running sum
				mostCalories = currentCalories
			}
			currentCalories = 0
		} else {
			// Add to running total
			addValue, err := strconv.Atoi(nextLine)
			if err != nil {
				log.Fatal(err)
			}
			currentCalories += addValue
		}
	}

	// Error handle
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Most calories: %d\n", mostCalories)
}
