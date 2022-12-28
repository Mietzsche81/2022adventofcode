package part2

type Instruction struct {
	distance int
	turn     string
}

type Board struct {
	Face [6]Face
	Meta [][]int
}

type Face struct {
	Id          int
	value       [50]string
	Edge        [4]Edge
	Location    int
	Orientation int
}

type State struct {
	Face *Face
	x    int
	y    int
	z    int
}

type Edge struct {
	Newface *Face
	Newedge int
}

var directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

var locationNeighbors = map[int][4]int{
	1:  {2, -3, -2, 3},
	-1: {-2, -3, 2, 3},
	2:  {-1, -3, 1, 3},
	-2: {1, -3, -1, 3},
	3:  {2, 1, -2, -1},
	-3: {2, -1, -2, 1},
}

func (in State) Apply(step Instruction) State {
	// initialize
	out := in.Copy()

	// advance
	for i := 0; i < step.distance; i++ {
		next := out.Advance()
		if next.Face.value[next.x][next.y] == ' ' {
			for next.Face.value[next.x][next.y] == ' ' {
				next = next.Advance()
			}
		}
		if next.Face.value[next.x][next.y] == '#' {
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
		Face: in.Face,
		x:    in.x,
		y:    in.y,
		z:    in.z,
	}
	next.x += directions[next.z][0]
	if next.x >= len(next.Face.value) {
		// TODO transition right
		return next
	} else if next.x < 0 {
		// TODO transition left
		return next
	}
	next.y += directions[next.z][1]
	if next.y >= len(next.Face.value[next.x]) {
		// TODO transition down
		return next
	} else if next.y < 0 {
		// TODO transition up
		return next
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
		Face: in.Face,
		x:    in.x,
		y:    in.y,
		z:    in.z,
	}
}
