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

type Record struct {
	lo int
	hi int
}

type Pair struct {
	a Record
	b Record
}

func main() {

	//
	// Read input
	//
	fileName := strings.TrimSpace(os.Args[1])
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	//
	// Process
	//
	found := 0
	scanner := bufio.NewScanner(fin)
	pattern := regexp.MustCompile(`,|-`)
	data := make([]Pair, 0)
	for scanner.Scan() {
		line := scanner.Text()
		match := pattern.Split(line, 4)
		ints := make([]int, 4)
		for i := range ints {
			ints[i], err = strconv.Atoi(match[i])
			if err != nil {
				log.Fatal(err)
			}
		}
		data = append(data, Pair{
			a: Record{
				lo: ints[0],
				hi: ints[1],
			},
			b: Record{
				lo: ints[2],
				hi: ints[3],
			},
		},
		)
		record := data[len(data)-1]
		if isIntersect(record) {
			found++
		}
	}

	//
	// Report
	//
	fmt.Println(found)
}

func IsSubset(rec Pair) bool {
	if (rec.a.lo <= rec.b.lo) && (rec.a.hi >= rec.b.hi) || (rec.b.lo <= rec.a.lo) && (rec.b.hi >= rec.a.hi) {
		return true
	}

	return false
}

func isIntersect(rec Pair) bool {
	if (rec.a.lo <= rec.b.hi) && (rec.b.hi <= rec.a.hi) {
		return true
	} else if (rec.a.lo <= rec.b.lo) && (rec.b.lo <= rec.a.hi) {
		return true
	} else if (rec.b.lo <= rec.a.hi) && (rec.a.hi <= rec.b.hi) {
		return true
	} else if (rec.b.lo <= rec.a.lo) && (rec.a.lo <= rec.b.hi) {
		return true
	}

	return false

}
