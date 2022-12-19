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

type Packet struct {
	val []any
}

func Quicksort(data []Packet, left, right int) {
	if left < right {
		p := Partition(data, left, right)
		Quicksort(data, left, p-1)
		Quicksort(data, p+1, right)
	}
}

func Partition(data []Packet, left, right int) int {
	// choose the most right element as our pivot
	pivot := data[right]

	// Swap elements that should go before pivot
	i := left
	for j := left; j < right; j++ {
		if data[j].Before(pivot) {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}

	/*
	 * After swapping every i & j:
	 *  - everything after i is larger than pivot
	 *  - everything before i is smaller than pivot
	 * so, put the pivot right in between.
	 */
	data[right], data[i] = data[i], data[right]

	return i
}

func (left Packet) Before(right Packet) bool {
	// Set a default value to true, break on contradiction
	direction := 0
	PacketCompare(left.val, right.val, &direction)

	if direction < 1 {
		return true
	} else {
		return false
	}
}

func PacketCompare(left any, right any, carry *int) {
	// Do not iterate deeper if contradiction found
	if *carry != 0 {
		return
	}

	// Get types
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)

	if leftType == EntryType && rightType == EntryType {
		*carry = intcompare(left.(float64), right.(float64))
	} else if leftType == ListType && rightType == ListType {
		for i := range left.([]any) {
			if i >= len(right.([]any)) {
				// Ran out of entries on right,  RIGHT goes first
				*carry = 1
				return
			}
			PacketCompare(left.([]any)[i], right.([]any)[i], carry)
			if *carry != 0 {
				// found result, done.
				return
			}
		}
		// If left side ran out before right side
		if len(left.([]any)) < len(right.([]any)) {
			// left side ran out, LEFT goes first
			*carry = -1
			return
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
		// LEFT FIRST
		return -1
	} else if left == right {
		// continue
		return 0
	} else {
		// RIGHT FIRST
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

	// Inject start and stop signal
	start := Packet{
		val: []any{[]any{2.0}},
	}
	stop := Packet{
		val: []any{[]any{6.0}},
	}
	data = append(data, start, stop)
	// Sort
	Quicksort(data, 0, len(data)-1)
	// Score
	score := 1
	for i, entry := range data {
		indicator := 0
		fmt.Printf("%4d %v\n", i, entry)
		if PacketCompare(entry.val, start.val, &indicator); indicator == 0 {
			score *= i + 1
		}
		indicator = 0
		if PacketCompare(entry.val, stop.val, &indicator); indicator == 0 {
			score *= i + 1
		}
	}

	//
	// report
	//
	fmt.Println(score)
}

func ParseInput(fileName string) (data []Packet) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	scanner := bufio.NewScanner(fin)

	data = make([]Packet, 0)
	for scanner.Scan() {
		// Extract
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		// Transform
		packet := Packet{}
		json.Unmarshal([]byte(line), &packet.val)
		data = append(data, packet)
	}
	return
}

func ProcessData(data []string) []string {
	out := make([]string, len(data))

	// Process

	return out
}
