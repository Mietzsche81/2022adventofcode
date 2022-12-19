package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type State struct {
	time      int
	stockpile map[string]int
	producers map[string]int
	blueprint map[string]map[string]int
}

func (s *State) Initialize() {
	s.time = 0
	s.stockpile = InitializeComponent()
	s.producers = InitializeComponent()
	for key, _ := range s.producers {
		s.blueprint[key] = InitializeComponent()
	}
	// default single ore miner
	s.producers["ore"] = 1
	//
}

func (src *State) copy() State {
	dst := State{
		time:      src.time,
		blueprint: src.blueprint,
	}
	dst.stockpile = CopyComponent(src.stockpile)
	dst.producers = CopyComponent(src.producers)

	return dst
}

func InitializeComponent() map[string]int {
	return map[string]int{
		"ore":      0,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
}

func CopyComponent(src map[string]int) map[string]int {
	dst := make(map[string]int)

	for key, val := range src {
		dst[key] = val
	}

	return dst
}

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	data := ParseInput(fileName)
	PrintBluePrint(data[28])

	//
	// process
	//
	for i, bp := range data {
		s := State{}
		s.Initialize()
		s.blueprint = bp
	}

	// out := ProcessData(data)
	//
	// report
	//

	// fmt.Println(out)
}

func ParseInput(fileName string) (blueprints []map[string]map[string]int) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	pattern := regexp.MustCompile(`Each (\w+) robot costs (\d+) (\w+)(?: and (\d+) (\w+))*`)
	scanner := bufio.NewScanner(fin)
	// data = make([]map[string]int)
	for scanner.Scan() {
		// Extract
		line := strings.TrimSpace(scanner.Text())
		match := pattern.FindAllStringSubmatch(line, -1)
		// Transform
		m := make(map[string]map[string]int)
		for _, robot := range match {
			m[robot[1]] = InitializeComponent()
			for i := 2; i < len(robot); i += 2 {
				cost, err := strconv.Atoi(robot[i])
				resource := robot[i+1]
				if err != nil {
					continue
				}
				m[robot[1]][resource] = cost
			}
		}
		// Load
		blueprints = append(blueprints, m)
	}
	return
}

func PrintBluePrint(bp map[string]map[string]int) {
	for robot, costs := range bp {
		fmt.Printf("%s robot costs:\n", robot)
		for resource, cost := range costs {
			fmt.Printf("\t%3d %s\n", cost, resource)
		}
	}
}

func ProcessData(data []string) []string {
	out := make([]string, len(data))

	// Process

	return out
}
