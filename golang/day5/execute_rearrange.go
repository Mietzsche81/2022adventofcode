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

type Command struct {
	quantity int
	source   int
	dest     int
}

func main() {

	//
	// Read input
	//

	machineID, err := strconv.Atoi(strings.TrimSpace(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	fileName := strings.TrimSpace(os.Args[2])
	stack := ParseGrid(fileName)
	steps := ParseProcedure(fileName)

	//
	// Execute Procedure
	//
	PrintStacks(&stack)
	if machineID == 9000 {
		ExecuteProcedure9000(&stack, &steps)
	} else {
		ExecuteProcedure9001(&stack, &steps)

	}
	PrintStacks(&stack)

	//
	// Report
	//
	PrintTop(&stack)
}

func PrintStacks(stack *[]string) {
	for i := range *stack {
		fmt.Printf("%d %s\n", i+1, (*stack)[i])
	}
}

func PrintTop(stack *[]string) {
	for _, s := range *stack {
		fmt.Printf("%c", s[len(s)-1])
	}
	fmt.Printf("\n")
}

func PrintSteps(steps *[]Command) {
	for _, com := range *steps {
		fmt.Printf("%2dx : %2d -> %2d\n", com.quantity, com.source, com.dest)
	}
}

func ParseGrid(fileName string) (stack []string) {
	// Open stream
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	scanner := bufio.NewScanner(fin)

	// Step by lines
	pattern := regexp.MustCompile(`(\s{3,4})|\[([A-Z])\]\s?|\s(\d+)\s{0,2}`)
	grid := make([][]string, 0)
	for scanner.Scan() {
		match := pattern.FindAllStringSubmatch(scanner.Text(), -1)
		// Break when reaching blank line
		if len(match) == 0 {
			break
		}
		// Parse crates
		row := make([]string, len(match))
		for i, group := range match {
			row[i] = group[2]
		}
		// Append row
		grid = append(grid, row)
	}
	// Transpose grid to stacks
	stack = make([]string, len(grid[0]))
	for i := range grid {
		for j := range grid[i] {
			stack[j] = grid[i][j] + stack[j]
		}
	}
	return stack
}

func ParseProcedure(fileName string) (steps []Command) {
	// Open stream
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	scanner := bufio.NewScanner(fin)

	// Step by lines
	pattern := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for scanner.Scan() {
		match := pattern.FindAllStringSubmatch(scanner.Text(), -1)
		if len(match) == 0 {
			// Skip blank lines
			continue
		}
		// Append commands
		quantity, err := strconv.Atoi(match[0][1])
		if err != nil {
			log.Fatal(err)
		}
		source, err := strconv.Atoi(match[0][2])
		if err != nil {
			log.Fatal(err)
		}
		dest, err := strconv.Atoi(match[0][3])
		if err != nil {
			log.Fatal(err)
		}
		steps = append(steps,
			Command{
				quantity: quantity,
				source:   source,
				dest:     dest,
			},
		)
	}
	return steps
}

func ExecuteProcedure9000(stack *[]string, steps *[]Command) {
	for _, command := range *steps {
		for i := 0; i < command.quantity; i++ {
			popPush(stack, command.source, command.dest)
		}
	}
}

func ExecuteProcedure9001(stack *[]string, steps *[]Command) {
	for _, command := range *steps {
		multipopPush(stack, command.source, command.dest, command.quantity)
	}
}

func multipopPush(stack *[]string, src int, dst int, x int) {
	src--
	dst--
	cut := len((*stack)[src]) - x
	move := (*stack)[src][cut:]
	(*stack)[src] = (*stack)[src][:cut]
	(*stack)[dst] += move
}

func popPush(stack *[]string, src int, dst int) {
	src--
	dst--
	move := pop(&(*stack)[src])
	push(&(*stack)[dst], move)
}

func pop(s *string) rune {
	end := len(*s) - 1
	last := rune((*s)[end])
	*s = (*s)[:end]

	return last
}

func push(s *string, add rune) {
	*s += string(add)
}
