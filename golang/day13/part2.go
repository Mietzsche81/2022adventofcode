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

type Packet struct {
	left    []any
	right   []any
	correct bool
}

var EntryType = reflect.TypeOf(0.0)
var ListType = reflect.TypeOf([]any{})

func (p *Packet) Compare() bool {
	// Set a default value to true, break on contradiction
	p.correct = true
	p.subcompare(p.left, p.right)

	return p.correct
}

func (p *Packet) subcompare(left any, right any) (ret int) {
	// Do not iterate deeper if contradiction found
	if !p.correct {
		return 1
	}

	// Get types
	leftType := reflect.TypeOf(left)
	rightType := reflect.TypeOf(right)
	fmt.Printf("%15T %v\n", left, left)
	fmt.Printf("%15T %v\n\n", right, right)

	if leftType == EntryType && rightType == EntryType {
		ret = intcompare(left.(float64), right.(float64))
		if ret == 1 {
			// left side is higher value, BAD
			p.correct = false
		}
	} else if leftType == ListType && rightType == ListType {
		for i := range left.([]any) {
			if i >= len(right.([]any)) {
				// Ran out of entries on right, BAD
				p.correct = false
				ret = 1
				break
			}
			ret = p.subcompare(left.([]any)[i], right.([]any)[i])
			if ret == -1 {
				// Found the left side is smaller, CORRECT (can leave)
				break
			} else if ret == 0 {
				// Tie, need to continue search
				continue
			} else {
				// Found the left side is larger, BAD (can leave)
				p.correct = false
				break
			}
		}
		// If left side ran out before right side
		if len(left.([]any)) < len(right.([]any)) {
			// left side ran out, GOOD
			ret = -1
		}
	} else if leftType == EntryType && rightType == ListType {
		ret = p.subcompare([]any{left}, right)
	} else if leftType == ListType && rightType == EntryType {
		ret = p.subcompare(left, []any{right})
	} else {
		log.Fatal(fmt.Errorf("Unrecognized type comparisons: '%T' v '%T'", left, right))
	}

	return ret
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

func (p *Packet) intVright(left int, right any) (correct bool) {
	switch val := right.(type) {
	case float64:
		correct = left <= int(val)
	}
	return
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
			data = append(data, Packet{})
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
