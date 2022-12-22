package part1

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseInput(fileName string) (data []Entry) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	i := 0
	data = make([]Entry, 0)
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
		entry := Entry{
			value:    value,
			position: i,
		}
		if i > 0 {
			entry.front = &data[i-1]
			data[i-1].back = &entry
		}
		data = append(data, entry)
		i++
	}
	data[0].front = &data[i-1]
	data[i-1].back = &data[0]

	return
}
