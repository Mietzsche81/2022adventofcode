package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ParseInput(fileName string) (data [][2]int) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	data = make([][2]int, 0)
	i := 0
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		// Extract
		line := strings.TrimSpace(scanner.Text())
		// Load
		for j, char := range line {
			if char == '#' {
				data = append(data, [2]int{i, j})
			}
		}
		// iterate
		i++
	}

	return
}

func Print(state [][2]int) {
	imin, imax := state[0][0], state[0][0]
	jmin, jmax := state[0][1], state[0][1]
	for _, elf := range state {
		if elf[0] < imin {
			imin = elf[0]
		} else if elf[0] > imax {
			imax = elf[0]
		}
		if elf[1] < jmin {
			jmin = elf[1]
		} else if elf[1] > jmax {
			jmax = elf[1]
		}
	}
	for i := imin; i <= imax; i++ {
		for j := jmin; j <= jmax; j++ {
			filled := false
			for _, elf := range state {
				if elf[0] == i && elf[1] == j {
					filled = true
					break
				}
			}
			if filled {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
