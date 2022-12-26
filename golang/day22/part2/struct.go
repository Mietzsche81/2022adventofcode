package part2

import "fmt"

type Instruction struct {
	distance int
	turn     string
}

type Board struct {
	face [6]Face
	meta [][]int
}

type Face struct {
	value [50]string
	edge  [4]Edge
}

type State struct {
	face *Face
	x    int
	y    int
	z    int
}

type Edge struct {
	newface *Face
	newz    int
}

var directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func (in State) Apply(step Instruction) State {
	// initialize
	out := in.Copy()

	// advance
	for i := 0; i < step.distance; i++ {
		next := out.Advance()
		if next.face.value[next.x][next.y] == ' ' {
			for next.face.value[next.x][next.y] == ' ' {
				next = next.Advance()
			}
		}
		if next.face.value[next.x][next.y] == '#' {
			break
		}
		out = next.Copy()
	}

	// turn
	out.Turn(step.turn)

	//output
	return out
}

func (in State) Advance() State {
	next := State{
		face: in.face,
		x:    in.x,
		y:    in.y,
		z:    in.z,
	}
	next.x += directions[next.z][0]
	if next.x >= len(next.face.value) {
		// TODO transition right
	} else if next.x < 0 {
		// TODO transition left
	}
	next.y += directions[next.z][1]
	if next.y >= len(next.face.value[next.x]) {
		// TODO transition down
	} else if next.y < 0 {
		// TODO transition up
	}

	return next
}

func (s *State) Turn(dir string) {
	if dir == "R" {
		s.z++
	} else {
		s.z--
	}
	if s.z >= len(directions) {
		s.z = 0
	} else if s.z < 0 {
		s.z = len(directions) - 1
	}
}

func (in *State) Copy() State {
	return State{
		face: in.face,
		x:    in.x,
		y:    in.y,
		z:    in.z,
	}
}

func (s State) Print() {
	fmt.Sprintf("%v [%d %d %s]\n",
		s.face, s.x, s.y, EncodeDirection(s.z),
	)
}
