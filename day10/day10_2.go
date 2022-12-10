package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	cycleNumber := ProcessData(data)
	signal := DisplaySignal(cycleNumber)
	//
	// report
	//

	fmt.Println(signal)
}

func ParseInput(fileName string) (data []int) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	scanner := bufio.NewScanner(fin)
	data = make([]int, 0)
	for scanner.Scan() {
		// Extract
		line := strings.Split(scanner.Text(), " ")
		// Transform
		if line[0] == "noop" {
			// Single cycle, no movement
			data = append(data, 0)
		} else {
			// Assume addx, grab increment
			incr, err := strconv.Atoi(line[1])
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, 0)
			data = append(data, incr)
		}
	}
	return
}

func ProcessData(data []int) []int {
	out := make([]int, len(data)+1)
	// Initialize cycle number
	out[0] = 1
	for i, incr := range data {
		// Apply any addx operations
		out[i+1] = out[i] + incr
	}
	return out
}

func DisplaySignal(data []int) string {
	// Apply initial cycle number
	score := "#"
	// How wide is sprite (center + radius)
	cursorRadius := 1
	// How many pixels to display before new line
	screenLength := 40
	for i := 1; i < len(data); i++ {
		// Normalize sprite signal to screen domain
		sprite := data[i] % screenLength
		// Normalize CRT's draw cursor to screen domain
		crt := i % screenLength
		// Find distance between draw cursor and sprite
		distance := sprite - crt
		if crt == 0 {
			// If cursor reset, add a new line
			score += "\n"
		}
		if (-cursorRadius <= distance) && (distance <= cursorRadius) {
			// Sprite within draw cursor
			score += "#"
		} else {
			// Sprite outside of draw cursor
			score += "."
		}
	}
	return score
}
