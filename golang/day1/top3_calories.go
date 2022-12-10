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

	//
	// Read input
	//

	var fileName string = strings.TrimSpace(os.Args[1])
	fmt.Printf("Processing input from: '%s'\n", fileName)
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	//
	// Scan to read line by line
	//

	scanner := bufio.NewScanner(fin)
	scanner.Split(bufio.ScanLines)
	mostCalories := []int{0, 0, 0}
	atEOF := false
	currentCalories := 0
	for !atEOF {
		// Get next group
		currentCalories, atEOF = sum_until_blank(scanner)

		// Compare against 3 largest
		for i, value := range mostCalories {
			if currentCalories > value {
				// Mark new max, slide others down list
				insert_truncate(mostCalories, currentCalories, i)
				break
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//
	// Report results
	//

	fmt.Println("Most calories: ")
	fmt.Println(mostCalories)
	fmt.Printf("Total: %d", sum(mostCalories...))
}

func sum_until_blank(scanner *bufio.Scanner) (total int, atEOF bool) {
	total = 0
	for {
		atEOF = !scanner.Scan()
		nextLine := scanner.Text()
		if strings.TrimSpace(nextLine) != "" {
			// Add to running total
			add, err := strconv.Atoi(nextLine)
			if err != nil {
				log.Fatal(err)
			}
			total += add
		} else {
			// empty line detected
			break
		}
		if atEOF {
			break
		}
	}
	return total, atEOF
}

func insert_truncate(array []int, insert int, i int) []int {
	return append(array[:i], append([]int{insert}, array[i:len(array)-1]...)...)
}

func sum(nums ...int) int {
	ret := 0
	for _, n := range nums {
		ret += n
	}
	return ret
}
