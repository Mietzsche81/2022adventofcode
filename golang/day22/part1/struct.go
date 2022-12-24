package part1

type Instruction struct {
	distance int
	turn     string
}

var directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func (step Instruction) apply(board []string, in [3]int) [3]int {
	out := [3]int{
		in[0],
		in[1],
		in[2],
	}
	// advance
	for i := 0; i < step.distance; i++ {
		next := Advance(board, out)
		if board[next[0]][next[1]] == ' ' {
			for board[next[0]][next[1]] == ' ' {
				next = Advance(board, next)
			}
		}
		if board[next[0]][next[1]] == '#' {
			break
		}
		out[0] = next[0]
		out[1] = next[1]
	}
	// turn
	if step.turn == "R" {
		out[2]++
	} else {
		out[2]--
	}
	if out[2] >= len(directions) {
		out[2] = 0
	} else if out[2] < 0 {
		out[2] = len(directions) - 1
	}

	return out
}

func Advance(board []string, state [3]int) [3]int {
	next := [3]int{
		state[0],
		state[1],
		state[2],
	}
	next[0] += directions[state[2]][0]
	if next[0] >= len(board) {
		next[0] = 0
	} else if next[0] < 0 {
		next[0] = len(board) - 1
	}
	next[1] += directions[state[2]][1]
	if next[1] >= len(board[next[0]]) {
		next[1] = 0
	} else if next[1] < 0 {
		next[1] = len(board[next[0]]) - 1
	}

	return next
}
