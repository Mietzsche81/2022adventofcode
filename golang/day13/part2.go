package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

var EntryType = reflect.TypeOf(0.0)
var ListType = reflect.TypeOf([]any{})

type Packet []any

func Quicksort(data []Packet, lo int, hi int) {

}

func Partition(data []Packet, lo int, hi int) int {
	// Default pivot: highest point
	pivotValue := data[hi]

	// Swap partitioned elements about the pivot
	i := lo - 1
	for j := lo; j < hi; j++ {
		if !data[j].GreaterThan(pivotValue) {
			i++
			tmp := &data[i]
			// TODO
		}
	}

	return i
}

func (left Packet) GreaterThan(right Packet) bool {
	// Set a default value to true, break on contradiction
	direction := 0
	PacketCompare(left, right, &direction)

	if direction == 1 {
		return true
	} else {
		return false
	}
}

func PacketCompare(left any, right any, carry *int) {
	// Do not iterate deeper if contradiction found
	if *carry == 1 {
		return
	}

	// Get types
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	fmt.Printf("%15T %v\n", left, left)
	fmt.Printf("%15T %v\n\n", right, right)

	if leftType == EntryType && rightType == EntryType {
		*carry = intcompare(left.(float64), right.(float64))
	} else if leftType == ListType && rightType == ListType {
		for i := range left.([]any) {
			if i >= len(right.([]any)) {
				// Ran out of entries on right, BAD
				*carry = 1
				break
			}
			PacketCompare(left.([]any)[i], right.([]any)[i], carry)
			if *carry == -1 {
				// Found the left side is smaller, CORRECT (can leave)
				break
			} else if *carry == 0 {
				// Tie, need to continue search
				continue
			} else {
				// Found the left side is larger, BAD (can leave)
				break
			}
		}
		// If left side ran out before right side
		if len(left.([]any)) < len(right.([]any)) {
			// left side ran out, GOOD
			*carry = -1
		}
	} else if leftType == EntryType && rightType == ListType {
		PacketCompare([]any{left}, right, carry)
	} else if leftType == ListType && rightType == EntryType {
		PacketCompare(left, []any{right}, carry)
	} else {
		log.Fatal(fmt.Errorf("Unrecognized type comparisons: '%T' v '%T'", left, right))
	}

	return
}

func intcompare(left float64, right float64) int {
	if left < right {
		return -1
	} else if left == right {
		return 0
	} else {
		return 1
	}
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

	// Inject
	score := 0
	for i, packet := range data[0:] {
		if packet.Compare() {
			score += i + 1
			fmt.Printf("++++++ Packet %d is CORRECT\n", i+1)
		} else {
			fmt.Printf("------ Packet %d is WRONG\n", i+1)
		}
	}

	//
	// report
	//

	fmt.Println(score)
}

func ParseInput(fileName string) (data []PacketSorter) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	scanner := bufio.NewScanner(fin)

	data = make([]PacketSorter, 0)
	for i := 0; scanner.Scan(); i++ {
		// Extract
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			// Decrement counter and skip this line
			i--
			continue
		}
		// Transform
		if i%2 == 0 {
			data = append(data, PacketSorter{})
			json.Unmarshal([]byte(line), &(data[i/2].left))
		} else {
			json.Unmarshal([]byte(line), &(data[i/2].right))
		}
	}
	return
}

func ProcessData(data []string) []string {
	out := make([]string, len(data))

	// Process

	return out
}
