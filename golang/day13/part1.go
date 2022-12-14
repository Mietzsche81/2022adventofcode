package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Packet struct {
	left  []any
	right []any
}

func (p *Packet) Compare() bool {

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

	// out := ProcessData(data)
	//
	// report
	//

	fmt.Println(data)
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
