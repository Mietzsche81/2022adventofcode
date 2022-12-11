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

type MonkeyOperation func(m *Monkey, i int)

type Monkey struct {
	items       []int
	operation   MonkeyOperation
	test        int
	dest        [2]int
	inspections int
}

func (m Monkey) Test(i int) int {
	if (m.items[i] % m.test) == 0 {
		return m.dest[0]
	} else {
		return m.dest[1]
	}
}

func (m *Monkey) Operate(i int) {
	m.operation(m, i)
}

func (m *Monkey) CreateOperation(s string) {
	reOp := regexp.MustCompile(`new = old ([\+\-\/\*]) (\d+|old)`)
	match := reOp.FindStringSubmatch(s)
	if match[2] == "old" {
		switch match[1] {
		case "*":
			m.operation = func(m *Monkey, i int) {
				m.items[i] *= m.items[i]
			}
		case "+":
			m.operation = func(m *Monkey, i int) {
				m.items[i] += m.items[i]
			}
		default:
			log.Fatal(fmt.Errorf("MakeOperation: Unknown operand %s", match[1]))
		}
	} else {
		incrementor, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}
		/*
			Use 'create' factory to create function pointer with "enclosed"
			incrementor value
		*/
		switch match[1] {
		case "*":
			create := func(incrementor int) MonkeyOperation {
				return func(m *Monkey, i int) {
					m.items[i] *= incrementor
				}
			}
			m.operation = create(incrementor)
		case "+":
			create := func(incrementor int) MonkeyOperation {
				return func(m *Monkey, i int) {
					m.items[i] += incrementor
				}
			}
			m.operation = create(incrementor)
		default:
			log.Fatal(fmt.Errorf("MakeOperation: Unknown operand %s", match[1]))
		}
	}

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
	fmt.Println(data)
	Simulate(data, 10000)
	fmt.Println(data)

	//
	// report
	//

	PrintInspections(data)
	fmt.Println(CalculateScore(data))
}

func ParseInput(fileName string) (data []Monkey) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	scanner := bufio.NewScanner(fin)
	data = make([]Monkey, 0)
	reNum := regexp.MustCompile(`\d+`)
	for scanner.Scan() {
		// Monkey #:
		newMonkey := Monkey{}
		line := scanner.Text()
		// Starting Items: #, #
		scanner.Scan()
		line = scanner.Text()
		items := reNum.FindAllString(line, -1)
		newMonkey.items = make([]int, len(items))
		for i, entry := range items {
			item, err := strconv.Atoi(entry)
			if err != nil {
				log.Fatal(err)
			}
			newMonkey.items[i] = item
		}
		// Operation: new = old ? ##
		scanner.Scan()
		line = scanner.Text()
		newMonkey.CreateOperation(line)
		// Test: divisible by ##
		scanner.Scan()
		num, err := strconv.Atoi(reNum.FindString(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		newMonkey.test = num
		// If true:
		scanner.Scan()
		num, err = strconv.Atoi(reNum.FindString(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		newMonkey.dest[0] = num
		// If false:
		scanner.Scan()
		num, err = strconv.Atoi(reNum.FindString(scanner.Text()))
		if err != nil {
			log.Fatal(err)
		}
		newMonkey.dest[1] = num
		// Load
		data = append(data, newMonkey)
		// Empty line
		scanner.Scan()
	}
	return
}

func Simulate(monkeys []Monkey, rounds int) {
	// Find common denominator among division tests
	commonDenominator := 1
	for j := range monkeys {
		commonDenominator *= monkeys[j].test
	}
	for i := 0; i < rounds; i++ {
		// Update
		for j := range monkeys {
			UpdateMonkey(monkeys, j)
		}
		// Scale by common denominator to prevent overflow
		for j := range monkeys {
			for k := range monkeys[j].items {
				monkeys[j].items[k] %= commonDenominator
			}
		}
	}
}

func UpdateMonkey(monkeys []Monkey, j int) {
	m := &monkeys[j]
	reps := len((*m).items)
	// All operations on item 0, NOT i, because throw items when done
	for i := 0; i < reps; i++ {
		// Inspect
		m.inspections += 1
		m.Operate(0)
		// Bored
		// (*m).items[0] /= 3
		// Test
		dest := m.Test(0)
		// Throw to dest
		value := (*m).items[0]
		(*m).items = remove((*m).items, 0)
		monkeys[dest].items = append(monkeys[dest].items, value)
	}
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func PrintInspections(monkeys []Monkey) {
	for _, m := range monkeys {
		fmt.Println(m.inspections)
	}
}

func CalculateScore(monkeys []Monkey) int {
	mostActive := make([]int, 2)
	for _, m := range monkeys {
		// Compare against 3 largest
		for i, value := range mostActive {
			if m.inspections > value {
				// Mark new max, slide others down list
				insert_truncate(mostActive, m.inspections, i)
				break
			}
		}
	}
	return mostActive[0] * mostActive[1]
}

func insert_truncate(array []int, insert int, i int) []int {
	return append(array[:i], append([]int{insert}, array[i:len(array)-1]...)...)
}
