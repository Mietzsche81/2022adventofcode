package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	//
	// Read input
	//

	groupSize, err := strconv.Atoi(strings.TrimSpace(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	fileName := strings.TrimSpace(os.Args[2])
	groups := ReadGroups(fileName, groupSize)

	//
	// Process data
	//
	badge := make([]rune, len(groups))
	score := 0
	for i, group := range groups {
		badge[i] = FindBadge(group)
		score += RtoI(badge[i])
	}

	//
	// Report Output
	//
	PrintGroups(groups)
	fmt.Printf("Total Priority: %d", score)
}

func ReadGroups(fileName string, groupSize int) [][]string {
	fin, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(fin), "\n")
	if len(text)%groupSize != 0 {
		log.Fatal(fmt.Errorf(
			"ReadGroups: %d line input. Must consist of groups of %d lines.\n",
			len(text),
			groupSize,
		))
	}
	groups := make([][]string, len(text)/groupSize)
	for i := range groups {
		group := make([]string, groupSize)
		for j := 0; j < groupSize; j++ {
			group[j] = text[3*i+j]
		}
		groups[i] = group
	}
	return groups
}

func PrintGroups(all [][]string) {
	fmt.Printf("%d groups of %d: \n", len(all), len(all[0]))
	for i, group := range all {
		for j, line := range group {
			fmt.Printf("GROUP %d LINE %d -- %s\n", i, j, line)
		}
	}
}

func RtoI(r rune) int {
	if (r-'a' < 26) && (r-'a' >= 0) {
		return int(r - 'a' + 1)
	} else {
		return int(r - 'A' + 27)
	}
}

func FindBadge(group []string) rune {
	for query, char := range group[0] {
		failed := false
		for _, line := range group[1:] {
			if strings.Index(line, group[0][query:query+1]) == -1 {
				failed = true
				break
			}
		}
		if !failed {
			return char
		}
	}
	return '?'
}
