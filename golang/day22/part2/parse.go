package part2

import (
	"bufio"
	"fmt"
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

	// Create metaboard for cube folding
	fmt.Println(len(data))
	board.meta = make([][]int, len(data)/50)
	id := 0
	for i := range board.meta {
		board.meta[i] = make([]int, len(data[i])/50)
		for j := range board.meta[i] {
			if data[50*i][50*j] != ' ' {
				board.meta[i][j] = id
				for x := 0; x < 50; x++ {
					board.face[id].value[x] = data[i+x][j : j+50]
				}
				id++
			} else {
				board.meta[i][j] = -1
			}
		}
		fmt.Println(board.meta[i])
	}

	// Use metaboard to create edges

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

func EncodeDirection(dir int) string {
	switch dir {
	case 0:
		return ">"
	case 1:
		return "v"
	case 2:
		return "<"
	case 3:
		return "^"
	}
	log.Fatal(fmt.Errorf("unrecognized direction %d", dir))
	return "?"
}
