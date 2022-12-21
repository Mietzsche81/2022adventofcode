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

func ParseInput(fileName string) (data map[string]*Monkey) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	data = make(map[string]*Monkey)
	pattern := regexp.MustCompile(`(\w+):\s*(\d+)?(?:(\w+)\s*([+-\/*])\s*(\w+))?`)
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		// Extract
		line := strings.TrimSpace(scanner.Text())
		match := pattern.FindStringSubmatch(line)
		name := match[1]
		sValue := match[2]
		operand1 := match[3]
		sOperation := match[4]
		operand2 := match[5]
		// Transform
		// Make sure the monkey of interest exists
		if _, exists := data[name]; !exists {
			data[name] = &Monkey{
				name: name,
			}
		}
		if len(sValue) > 0 {
			// Given value
			value, err := strconv.Atoi(sValue)
			if err != nil {
				log.Fatal(err)
			}
			data[name].value = &value
		} else if len(operand1) > 0 {
			// Given operation
			for i, monkeyName := range []string{operand1, operand2} {
				// Make sure the monkey of interest exists
				if _, exists := data[monkeyName]; !exists {
					data[monkeyName] = &Monkey{
						name: monkeyName,
					}
				}
				// Assign operands
				data[name].operands[i] = data[monkeyName]
			}
			// Assign operation
			data[name].AssignOperation(sOperation)
		} else {
			log.Fatal(fmt.Errorf("failed to parse line: %s", line))
		}
	}
	// Assign root operation
	data["root"].operation = eql
	return
}

func (m *Monkey) AssignOperation(op string) {
	switch op {
	case "+":
		m.operation = add
	case "-":
		m.operation = sub
	case "*":
		m.operation = mul
	case "/":
		m.operation = div
	default:
		log.Fatal(fmt.Errorf("unrecognized monkey operation '%s'", op))
	}
}

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}

func eql(a int, b int) int {
	if a == b {
		return 1
	} else {
		return 0
	}
}
