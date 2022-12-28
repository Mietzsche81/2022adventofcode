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

var topLeft = [][2]int{
	{0, 0},
	{0, 49},
	{49, 49},
	{49, 0},
}

var locationNeighbors = map[int][4]int{
	1:  {2, -3, -2, 3},
	-1: {-2, -3, 2, 3},
	2:  {-1, -3, 1, 3},
	-2: {1, -3, -1, 3},
	3:  {2, 1, -2, -1},
	-3: {2, -1, -2, 1},
}

func (in *State) Copy() State {
	return State{
		Face: in.Face,
		x:    in.x,
		y:    in.y,
		z:    in.z,
	}
}
