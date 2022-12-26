package part2

import "strings"

func Process(board Board, steps []Instruction) State {
	state := board.Initialize()
	for _, step := range steps {
		state = state.Apply(step)
	}

	return state
}

func (b *Board) Initialize() State {
	s := State{
		z:    0,
		face: &b.face[0],
	}
	s.y = strings.Index(b.face[0].value[0], ".")
	return s
}

func (b *Board) Score(s State) int {
	iFace := 0
	for i := range b.face {
		if &(b.face[i]) == s.face {
			iFace = i
			break
		}
	}
	var topleft [2]int
	for i := range b.meta {
		for j := range b.meta[i] {
			if b.meta[i][j] == iFace {
				topleft = [2]int{50 * i, 50 * j}
				break
			}
		}
	}
	return (topleft[0]+s.x+1)*1000 + (topleft[1]+s.y+1)*4 + s.z
}
