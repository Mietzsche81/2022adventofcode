package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type State struct {
	head [2]int
	tail [2]int
}

func main() {

	//
	// Read input
	//
	fileName := strings.TrimSpace(os.Args[1])
	steps := ParseInput(fileName)

	//
	// process
	//
	x0 := State{
		head: [2]int{0, 0},
		tail: [2]int{0, 0},
	}
	x := Simulate(steps, x0)
	tail := BreadCrumbsTail(x)

	//
	// report
	//
	fmt.Printf("Tail visited %d tiles", len(tail))
}

func ParseInput(fileName string) [][2]int {
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	scanner := bufio.NewScanner(fin)

	steps := make([][2]int, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		n, err := strconv.Atoi(line[1])
		step := EncodeDirection(line[0])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < n; i++ {
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

func Simulate(steps [][2]int, x0 State) []State {
	states := make([]State, 0)
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

func SimulateStep(step [2]int, in State) State {
	out := State{
		head: in.head,
		tail: in.tail,
	}
	for i := range step {
		out.head[i] += step[i]
	}
	dx, dy := out.head[0]-out.tail[0], out.head[1]-out.tail[1]
	if (dx > 1) || (dx < -1) || (dy > 1) || (dy < -1) {
		out.tail = in.head
	}
	return out
}

func BreadCrumbsTail(states []State) map[[2]int]int {
	trail := make(map[[2]int]int)
	for i, state := range states {
		trail[state.tail] = i
	}
	return trail
}
