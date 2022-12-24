package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseInput(fileName string) (board []string, steps []Instruction) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan map
	board = make([]string, 0)
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		// Extract
		line := scanner.Text()
		// Load
		if len(strings.TrimSpace(line)) == 0 {
			// Reached end of input, break
			break
		}
		board = append(board, line)
	}

	// Expand board as needed
	columns := len(board[0])
	for i, line := range board {
		if len(line) > columns {
			// Need to expand board size with whitespace
			columns := len(line)
			for j, previous := range board {
				if len(previous) < columns {
					board[j] = previous + strings.Repeat(" ", columns-len(previous))
				}
			}
		} else if len(line) < columns {
			board[i] = line + strings.Repeat(" ", columns-len(line))
		}
	}

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

func Print(board []string, state [3]int) {
	for i, line := range board {
		if state[0] == i {
			fmt.Printf("%s%s%s\n", line[:state[1]], EncodeDirection(state[2]), line[state[1]+1:])
		} else {
			fmt.Println(line)
		}
	}
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

func PrintState(state [3]int) string {
	return fmt.Sprintf("[%d %d %s]", state[0], state[1], EncodeDirection(state[2]))
}
