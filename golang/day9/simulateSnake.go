package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Snake struct {
	length int
	link   [][2]int
}

func NewSnake(length int) Snake {
	link := make([][2]int, length)
	return Snake{
		length: length,
		link:   link,
	}
}

func PrintSnake(s Snake) {
	for _, link := range s.link {
		fmt.Println(link)
	}
}

func main() {

	//
	// Read input
	//

	snakeLength, err := strconv.Atoi(strings.TrimSpace(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	fileName := strings.TrimSpace(os.Args[2])
	steps := ParseInput(fileName)

	//
	// process
	//

	x0 := NewSnake(snakeLength)
	x := Simulate(steps, x0)
	tail := BreadCrumbsTail(x)

	//
	// report
	//

	PrintSnake(x[len(x)-1])
	fmt.Printf("Tail visited %d tiles", len(tail))
}

func ParseInput(fileName string) [][2]int {
	// Open file
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	scanner := bufio.NewScanner(fin)

	// Read each line
	steps := make([][2]int, 0)
	for scanner.Scan() {
		// Grab values
		line := strings.Split(scanner.Text(), " ")
		repititions, err := strconv.Atoi(line[1])
		// Convert step directions to dx
		step := EncodeDirection(line[0])
		if err != nil {
			log.Fatal(err)
		}
		// Repeat steps for specified number of times
		for i := 0; i < repititions; i++ {
			steps = append(steps, step)
		}
	}
	return steps
}

func EncodeDirection(d string) [2]int {
	var pair [2]int
	switch d {
	case "U":
		pair = [2]int{-1, 0}
	case "D":
		pair = [2]int{1, 0}
	case "L":
		pair = [2]int{0, -1}
	case "R":
		pair = [2]int{0, 1}
	default:
		log.Fatal(fmt.Errorf("EncodeDirection: unrecognized '%s'", d))
	}
	return pair
}

func Simulate(steps [][2]int, x0 Snake) []Snake {
	states := make([]Snake, 0)
	states = append(states, x0)

	for _, step := range steps {
		states = append(states,
			SimulateStep(
				step,
				states[len(states)-1],
			),
		)
	}

	return states
}

func SimulateStep(step [2]int, in Snake) Snake {
	// Initialize output
	out := NewSnake(in.length)
	for i := range in.link {
		for dim := range in.link[0] {
			out.link[i][dim] = in.link[i][dim]
		}
	}
	// Move head by applying step
	for dim := range step {
		out.link[0][dim] += step[dim]
	}
	// Propagate
	for i := 1; i < out.length; i++ {
		// Find distance between points, and how to move
		dx := make([]int, len(out.link[i]))
		for dim := range dx {
			dx[dim] = out.link[i-1][dim] - out.link[i][dim]
		}
		// If stretched, need to move in the direction of the stretch
		direction, stretched := Normalize(dx)
		if stretched {
			for dim := range direction {
				out.link[i][dim] += direction[dim]
			}
		}
	}
	return out
}

// Force any vector to have a maximum component magnitude of 1.
// Returns the normalized vector followed by a bool indicating whether
// normalization was performed.
func Normalize(vector []int) ([]int, bool) {
	normal := make([]int, len(vector))
	scaled := false
	for i := range vector {
		if vector[i] < -1 {
			normal[i] = -1
			scaled = true
		} else if vector[i] > 1 {
			normal[i] = 1
			scaled = true
		} else {
			normal[i] = vector[i]
		}
	}
	return normal, scaled
}

func BreadCrumbsTail(states []Snake) map[[2]int]int {
	trail := make(map[[2]int]int)
	for i, state := range states {
		trail[state.link[state.length-1]] = i
	}
	return trail
}
