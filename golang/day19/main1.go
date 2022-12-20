package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"day19/part1"
)

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	maxTime, err := strconv.Atoi(strings.TrimSpace(os.Args[2]))
	if err != nil {
		log.Fatal(err)
	}
	data := part1.ParseInput(fileName)

	//
	// process
	//
	bpScore := []int{}
	for i, bp := range data {
		s := part1.State{}
		s.Initialize()
		s.Blueprint = bp
		out := s.OptimizeBFS(maxTime)
		fmt.Println(i, out)
		bpScore = append(bpScore, out)
	}

	//
	// report
	//
	totalScore := 0
	for i, score := range bpScore {
		totalScore += (i + 1) * score
	}
	fmt.Println(totalScore)
}
