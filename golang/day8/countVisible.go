package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	//
	// Read input
	//
	fileName := strings.TrimSpace(os.Args[1])
	grid := ParseGrid(fileName)

	//
	// Process input
	//

	hidden := FindEclipsed(&grid)
	score := ScoreForest(&grid)

	//
	// Report out
	//

	PrintReport(&grid, &hidden, &score)
	fmt.Printf("\n%d Trees Visible\n", CountVisible(&hidden))
	fmt.Printf("\nMost scenic: %d\n", FindLargest(&score))
}

func PrintReport(grid *[][]int, result *[][]int, score *[][]int) {
	for i := range *grid {
		fmt.Printf("%v\t%v\t%v\n", (*grid)[i], (*result)[i], (*score)[i])
	}
}

func ParseGrid(fileName string) [][]int {
	// Open file
	fin, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()
	linestream := bufio.NewScanner(fin)
	linestream.Split(bufio.ScanLines)

	// Create grid
	grid := make([][]int, 0)
	for linestream.Scan() {
		line := strings.TrimSpace(linestream.Text())
		row := make([]int, len(line))

		// Read each rune from line, convert to int
		for i, height := range line {
			row[i] = int(height - '0')
			if row[i] < 0 || row[i] > 9 {
				err := fmt.Errorf("ParseGrid: Tree height '%c' not understood", height)
				log.Fatal(err)
			}
		}
		grid = append(grid, row)
	}
	err = linestream.Err()
	if err != nil {
		log.Fatal(err)
	}
	return grid
}

func FindEclipsed(grid *[][]int) [][]int {
	// Create empty results
	result := make([][]int, len(*grid))
	for i := range result {
		result[i] = make([]int, len((*grid)[i]))
	}

	// Check each tree, skipping edges
	for i := 1; i < (len(result) - 1); i++ {
		for j := 1; j < (len(result[i]) - 1); j++ {
			result[i][j] = IsEclipsed(grid, i, j)
		}
	}
	return result
}

func IsEclipsed(grid *[][]int, i int, j int) int {
	// value of point of interest
	poi := (*grid)[i][j]
	// Directions to scan
	directions := [][]int{
		{-1, 0}, // north
		{1, 0},  // south
		{0, 1},  // east
		{0, -1}, // west
	}

	// Check in each direction
	rows, cols := len(*grid), len((*grid)[i])
	for _, dir := range directions {
		// initialize direction
		k, l := i+dir[0], j+dir[1]
		tallest := (*grid)[k][l]
		// find tallest tree in direction
		for (k >= 0) && (l >= 0) && (k < rows) && (l < cols) {
			// if tallest tree so far, capture
			if tallest < (*grid)[k][l] {
				tallest = (*grid)[k][l]
			}
			// iterate
			k += dir[0]
			l += dir[1]
		}
		// return not eclipsed if tallest in direction doesn't eclipse
		if tallest < poi {
			return 0
		}
	}

	return 1
}

func CountVisible(results *[][]int) int {
	total := 0
	for _, row := range *results {
		for _, cell := range row {
			if cell == 0 {
				total++
			}
		}
	}
	return total
}

func ScoreForest(grid *[][]int) [][]int {
	// Create empty results
	result := make([][]int, len(*grid))
	for i := range result {
		result[i] = make([]int, len((*grid)[i]))
	}

	// Check each tree
	for i := range result {
		for j := range result[i] {
			result[i][j] = ScoreTree(grid, i, j)
		}
	}
	return result
}

func ScoreTree(grid *[][]int, i int, j int) int {
	// value of point of interest
	poi := (*grid)[i][j]
	// Directions to scan
	directions := [][]int{
		{-1, 0}, // north
		{1, 0},  // south
		{0, 1},  // east
		{0, -1}, // west
	}

	// Check in each direction
	rows, cols := len(*grid), len((*grid)[i])
	dir_score := make([]int, len(directions))
	for d, dir := range directions {
		// initialize direction
		k, l := i+dir[0], j+dir[1]
		// find tallest tree in direction
		for (k >= 0) && (l >= 0) && (k < rows) && (l < cols) {
			// increment direction score
			dir_score[d]++
			// if tree is blocking, break
			if poi <= (*grid)[k][l] {
				break
			}
			// iterate
			k += dir[0]
			l += dir[1]
		}
	}

	// Tally score
	score := 1
	for _, multiplier := range dir_score {
		score *= multiplier
	}
	// fmt.Printf("CHECKING [%d][%d]: %v %d\n", i, j, dir_score, score)
	return score
}

func FindLargest(grid *[][]int) int {
	largest := 0
	for _, row := range *grid {
		for _, cell := range row {
			if cell > largest {
				largest = cell
			}
		}
	}
	return largest
}
