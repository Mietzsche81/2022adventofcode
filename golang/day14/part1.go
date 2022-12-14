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

type Board struct {
	square [][]Square
	bounds [2][2]int
	depth  *int
	source *Square
	active *Square
	count  int
}

type Square struct {
	value rune
	i     int
	j     int
}

var DIRECTION map[string][2]int = map[string][2]int{
	"l":         {-1, 0},
	"r":         {1, 0},
	"u":         {0, -1},
	"d":         {0, 1},
	"FallDown":  {0, 1},
	"FallLeft":  {-1, 1},
	"FallRight": {1, 1},
}

func IsDirection(x [2]int, s string) bool {
	for i, val := range x {
		if DIRECTION[s][i] != val {
			return false
		}
	}
	return true
}

func Normalize(in [2]int) (out [2]int) {
	for i, val := range in {
		if val >= 1 {
			out[i] = 1
		} else if val <= -1 {
			out[i] = -1
		}
	}
	return out
}

func CreateBoard(forms [][][2]int) Board {
	// Find the bounds of the board, and add a margin around it.
	min, max := MinMax(forms)
	min[1] = 0
	if min[0] > 0 {
		min[0] -= 1
	}
	max[0] += 2
	max[1] += 2
	// Create empty board
	b := Board{
		bounds: [2][2]int{min, max},
	}
	b.square = make([][]Square, max[0]+1)
	for i := range b.square {
		b.square[i] = make([]Square, max[1]+2)
		for j := range b.square[i] {
			b.square[i][j] = Square{
				value: '.',
				i:     i,
				j:     j,
			}
		}
	}
	// Set the bounds of the board just beyond the squares used
	min[1] = 0
	b.depth = &b.bounds[1][1]
	// Initialize source
	b.source = &b.square[500][0]
	b.source.value = '+'
	// Initialize formations
	for _, form := range forms {
		for p := range form[:len(form)-1] {
			// Find direction between current points
			dir := Normalize([2]int{
				form[p+1][0] - form[p][0],
				form[p+1][1] - form[p][1],
			})
			// Paint each
			i, j := form[p][0], form[p][1]
			for (i != form[p+1][0]) || (j != form[p+1][1]) {
				b.square[i][j].value = '#'
				i += dir[0]
				j += dir[1]
			}
			b.square[i][j].value = '#'
		}
	}

	return b
}

func (b *Board) Print() {
	for j := b.bounds[0][1]; j < b.bounds[1][1]; j++ {
		for i := b.bounds[0][0]; i < b.bounds[1][0]; i++ {
			fmt.Printf("%c", b.square[i][j].value)
		}
		fmt.Printf("\n")
	}
}

func (b *Board) Simulate() {
	for b.source != nil {
		b.CreateDrop()
	}
}

func (b *Board) CreateDrop() {
	if b.source == nil {
		return
	}
	b.active = b.source
	b.count += 1
	for b.active != nil {
		b.AdvanceDrop()
	}
	// Subtract the drop increment if it fell off
	if b.source == nil {
		b.count -= 1
	}
}

func (b *Board) AdvanceDrop() {
	if b.active == nil {
		// If no active drop, can't do anything
		return
	} else if (b.active.j + 1) >= *b.depth {
		// If falling off the edge, we are done
		b.active = nil
		b.source = nil
		return
	}
	// Advance drop
	i, j := b.active.i, b.active.j
	if b.square[i][j+1].value == '.' {
		b.active = &b.square[i][j+1]
	} else if b.square[i-1][j+1].value == '.' {
		b.active = &b.square[i-1][j+1]
	} else if b.square[i+1][j+1].value == '.' {
		b.active = &b.square[i+1][j+1]
	} else {
		b.active.value = 'O'
		b.active = nil
	}
}

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	formations := ParseInput(fileName)
	fmt.Println(formations)

	//
	// process
	//
	board := CreateBoard(formations)
	board.Simulate()

	//
	// report
	//

	board.Print()
	fmt.Println(board.count)
}

func ParseInput(fileName string) [][][2]int {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	pattern := regexp.MustCompile(`(\d+),(\d+)`)
	scanner := bufio.NewScanner(fin)
	data := make([][][2]int, 0)
	for scanner.Scan() {
		// Extract
		line := strings.TrimSpace(scanner.Text())
		match := pattern.FindAllStringSubmatch(line, -1)
		// Transform
		formation := make([][2]int, len(match))
		for i := range match {
			for j, x := range match[i][1:] {
				var err error
				formation[i][j], err = strconv.Atoi(x)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		// Load
		data = append(data, formation)
	}
	return data
}

func ProcessData(data []string) []string {
	out := make([]string, len(data))

	// Process

	return out
}

func MinMax(all [][][2]int) (min [2]int, max [2]int) {
	min = all[0][0]
	max = all[0][0]

	for _, form := range all {
		for _, point := range form {
			for k, value := range point {
				if min[k] > value {
					min[k] = value
				}
				if max[k] < value {
					max[k] = value
				}
			}
		}
	}
	return
}
