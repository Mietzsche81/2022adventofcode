package part1

import (
	"strings"
)

func Process(board []string, steps []Instruction) [3]int {
	state := Initialize(board)
	for _, step := range steps {
		state = step.apply(board, state)
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
