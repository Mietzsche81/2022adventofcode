package part1

import (
	"fmt"
	"strings"
)

func Process(board []string, steps []Instruction) [3]int {
	state := Initialize(board)
	Print(board, state)
	fmt.Println("----------------------")
	for _, step := range steps {
		state = step.apply(board, state)
		// Print(board, state)
		fmt.Println("----------------------")
		//scanner := bufio.NewScanner(os.Stdin)
		//scanner.Scan()
	}

	return state
}

func Initialize(board []string) [3]int {
	col := strings.Index(board[0], ".")
	return [3]int{0, col, 0}
}

func Score(state [3]int) int {
	return 1000*(state[0]+1) + 4*(state[1]+1) + state[2]
}
