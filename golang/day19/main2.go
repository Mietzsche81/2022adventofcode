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
	for _, bp := range data[:3] {
		s := part1.State{}
		s.Initialize()
		s.Blueprint = bp
		out := s.Optimize(maxTime)
		bpScore = append(bpScore, out)
	}

	//
	// report
	//
	totalScore := 1
	for i, score := range bpScore {
		totalScore *= score
		fmt.Println(i, score)
	}
	fmt.Println(totalScore)
}
