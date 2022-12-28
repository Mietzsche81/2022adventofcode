package part2

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseInput(fileName string) (board Board, steps []Instruction) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan map
	data := make([]string, 0)
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		// Extract
		line := scanner.Text()
		// Load
		if len(strings.TrimSpace(line)) == 0 {
			// Reached end of input, break
			break
		}
		data = append(data, line)
	}

	// Expand board as needed to create cube space
	columns := len(data[0])
	for i, line := range data {
		if len(line) > columns {
			// Need to expand board size with whitespace
			columns := len(line)
			for j, previous := range data {
				if len(previous) < columns {
					data[j] = previous + strings.Repeat(" ", columns-len(previous))
				}
			}
		} else if len(line) < columns {
			data[i] = line + strings.Repeat(" ", columns-len(line))
		}
	}

	// Cube folding
	board.createMetaboard(data)
	board.findNeighbors()

	// Scan instructions
	scanner.Scan()
	pattern := regexp.MustCompile(`(\d+)(R|L)`)
	match := pattern.FindAllStringSubmatch(scanner.Text(), -1)
	steps = make([]Instruction, len(match))
	for i, instruct := range match {
		turn := instruct[2]
		distance, err := strconv.Atoi(instruct[1])
		if err != nil {
			log.Fatal(err)
		}
		steps[i].distance = distance
		steps[i].turn = turn
	}

	return
}
