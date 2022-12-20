package part1

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseInput(fileName string) (blueprints []BlueprintMap) {
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
		m := make(BlueprintMap)
		for _, robot := range match {
			m[robot[1]] = initializeComponent()
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
