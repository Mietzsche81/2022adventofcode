package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Square struct {
	value       int
	neighbor    []*Square
	i           int
	j           int
	visited     bool
	visit_index int
}

type Board struct {
	square [][]Square
	m      int
	n      int
}

type Node struct {
	parent   *Node
	child    *Node
	location *Square
	distance int
}

func (b *Board) FindAllNeighbors() {
	for i := 0; i < b.m; i++ {
		for j := 0; j < b.n; j++ {
			b.FindNeighbor(i, j)
		}
	}
}

func (b *Board) FindNeighbor(i int, j int) {
	b.square[i][j].neighbor = make([]*Square, 0)
	for _, d := range DIRECTION {
		x, y := i+d[0], j+d[1]
		if x < 0 || x >= b.m || y < 0 || y >= b.n {
			continue
		}
		here, there := b.square[i][j].value, b.square[x][y].value
		if (there - here) <= 1 {
			b.square[i][j].neighbor = append(
				b.square[i][j].neighbor,
				&(b.square[x][y]),
			)
		}
	}
}

func (b *Board) Print(fileName string) {
	fout, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	for i := 0; i < b.m; i++ {
		for j := 0; j < b.n; j++ {
			fout.WriteString(fmt.Sprintf("%4d ", b.square[i][j].visit_index))
		}
		fout.WriteString("\n")
	}
}

func (here *Node) Advance() (there []Node) {
	there = make([]Node, 0)
	distance := here.distance + 1
	for _, neighbor := range here.location.neighbor {
		if !neighbor.visited {
			there = append(
				there,
				Node{
					parent:   here,
					location: neighbor,
					distance: distance,
				},
			)
		}
	}

	return
}

func (n *Node) IsAt(xf *Square) bool {
	return n.location == xf
}

var DIRECTION [4][2]int = [4][2]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func RtoI(r rune) int {
	if r == 'S' {
		return 0
	} else if r == 'E' {
		return 26
	} else {
		return int(r - 'a')
	}
}

func main() {

	//
	// Read input
	//

	fileName := strings.TrimSpace(os.Args[1])
	board, x0, xf := ParseInput(fileName)
	board.FindAllNeighbors()

	//
	// process
	//

	path := FindShortestPathBFS(&board, x0, xf)

	//
	// report
	//
	board.Print("result.csv")
	fmt.Printf("Started at [%d %d] seeking [%d %d]\n", x0.i, x0.j, xf.i, xf.j)
	fmt.Printf("Takes %d steps to reach.\n", path.distance)

}

func ParseInput(fileName string) (b Board, x0 *Square, xf *Square) {
	// Open File
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Scan to read line by line
	b.square = make([][]Square, 0)
	i := 0
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		// Extract
		line := strings.TrimSpace(scanner.Text())
		row := make([]Square, len(line))
		// Transform
		for j, char := range line {
			row[j].value = RtoI(char)
			row[j].i = i
			row[j].j = j
			row[j].visited = false
			if char == 'S' {
				x0 = &row[j]
			} else if char == 'E' {
				xf = &row[j]
			}
		}
		// Load
		b.square = append(b.square, row)
		i++
	}
	b.m = len(b.square)
	b.n = len(b.square[0])
	return
}

func FindShortestPathBFS(b *Board, x0 *Square, xf *Square) *Node {
	// Define removal logic
	RemoveDuplicates := func(in []Node) []Node {
		reduce := make(map[*Square]bool)
		out := []Node{}
		for _, item := range in {
			if _, value := reduce[item.location]; !value {
				reduce[item.location] = true
				out = append(out, item)
			}
		}
		return out
	}

	fmt.Printf("Traversing %3d x %3d board (size = %6d)\n", b.m, b.n, b.m*b.n)

	var path *Node = nil
	distance := 0
	here := make([]Node, 0)
	there := []Node{{location: x0}}
	for path == nil {
		// Iterate
		here = there
		there = make([]Node, 0)
		distance++
		// Mark all as visited
		for i := range here {
			here[i].location.visited = true
			here[i].location.visit_index = distance - 1
		}
		// Check each candidate location for destinations
		for i := range here {
			add := here[i].Advance()
			there = append(there, add...)
		}
		// Deduplicate
		there = RemoveDuplicates(there)
		// Report
		if distance%10 == 0 {
			fmt.Printf("Iteration %d: %5d candidate nodes found %5d children\n", distance, len(here), len(there))
		}
		// Check if any have finished
		for i := range there {
			if there[i].IsAt(xf) {
				path = &there[i]
				break
			}
		}
	}

	// Reverse link the path for children
	node := path
	for node.parent != nil {
		node.parent.child = node
		node = node.parent
	}

	return path
}
