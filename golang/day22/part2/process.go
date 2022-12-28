package part2

import (
	"fmt"
	"log"
	"strings"
)

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
		Face: &b.Face[0],
	}
	s.y = strings.Index(b.Face[0].value[0], ".")
	return s
}

func (b *Board) Score(s State) int {
	iFace := 0
	for i := range b.Face {
		if &(b.Face[i]) == s.Face {
			iFace = i
			break
		}
	}
	var topleft [2]int
	for i := range b.Meta {
		for j := range b.Meta[i] {
			if b.Meta[i][j] == iFace {
				topleft = [2]int{50 * i, 50 * j}
				break
			}
		}
	}
	return (topleft[0]+s.x+1)*1000 + (topleft[1]+s.y+1)*4 + s.z
}

func EncodeDirection(dir int) string {
	switch dir {
	case 0:
		return ">"
	case 1:
		return "v"
	case 2:
		return "<"
	case 3:
		return "^"
	}
	log.Fatal(fmt.Errorf("unrecognized direction %d", dir))
	return "?"
}
