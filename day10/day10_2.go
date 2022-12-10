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
	signal := ScoreData(cycleNumber)
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
			data = append(data, 0)
		} else {
			mod, err := strconv.Atoi(line[1])
			if err != nil {
				log.Fatal(err)
			}
			data = append(data, 0)
			data = append(data, mod)
		}
	}
	return
}

func ProcessData(data []int) []int {
	out := make([]int, len(data)+1)
	out[0] = 1
	for i, incr := range data {
		out[i+1] = out[i] + incr
	}
	return out
}

func ScoreData(data []int) string {
	score := "#"
	cursorRadius := 1
	for i := 1; i < len(data); i++ {
		sprite := data[i] % 40
		crt := i % 40
		distance := sprite - crt
		if crt == 0 {
			score += "\n"
		}
		if (-cursorRadius <= distance) && (distance <= cursorRadius) {
			score += "#"
		} else {
			score += "."
		}
	}
	return score
}
