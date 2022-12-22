package part2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseInput(fileName string) (data []int) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	key := 811589153
	data = make([]int, 0)
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		// Extract
		line := strings.TrimSpace(scanner.Text())
		// Transform
		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		// Load
		data = append(data, value*key)
	}

	return
}
