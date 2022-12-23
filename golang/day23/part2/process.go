package part2

import "fmt"

var directions = map[string][2]int{
	"N":  {-1, 0},
	"NE": {-1, 1},
	"E":  {0, 1},
	"SE": {1, 1},
	"S":  {1, 0},
	"SW": {1, -1},
	"W":  {0, -1},
	"NW": {-1, -1},
	"-":  {0, 0},
}

var rules = []string{"-", "N", "S", "W", "E"}

var corners = map[string][]string{
	"N": {"N", "NE", "NW"},
	"S": {"S", "SE", "SW"},
	"W": {"W", "NW", "SW"},
	"E": {"E", "NE", "SE"},
	"-": {"N", "NE", "E", "SE", "S", "SW", "W", "NW"},
}

func CycleDirections() {
	newrules := []string{rules[0]}
	newrules = append(newrules, rules[2:]...)
	rules = append(newrules, rules[1])
}

func Process(data [][2]int) ([][2]int, int) {
	state := make([][2]int, len(data))
	copy(state, data)
	i := 0
	for {
		i++
		next, steady := Propose(state)
		if steady {
			break
		}
		state = Negotiate(state, next)
		CycleDirections()
		if i%100 == 0 {
			fmt.Println(i)
		}
	}

	return state, i
}

func Propose(state [][2]int) ([][2]int, bool) {
	proposed := make([][2]int, len(state))
	steady := true
	copy(proposed, state)

	for i, elf := range proposed {
		for _, move := range rules {
			if ValidMove(state, elf, move) {
				if move != "-" {
					steady = false
				}
				for axis := range elf {
					proposed[i][axis] += directions[move][axis]
				}
				break
			}
		}
	}

	return proposed, steady
}

func ValidMove(statemap [][2]int, poi [2]int, rule string) bool {
	for _, corner := range corners[rule] {
		// Pick out neighbor direction to query
		query := directions[corner]
		// construct query by adding direction & point of interest
		for axis := range query {
			query[axis] += poi[axis]
		}
		// See if an elf occupies the queried location
		for _, elf := range statemap {
			occupied := true
			for axis := range query {
				if query[axis] != elf[axis] {
					// Mismatched axis, no elf at cell
					occupied = false
					break
				}
			}
			if occupied {
				// If cell is occupied, do not execute rule
				return false
			}
		}

	}
	return true
}

func Negotiate(current [][2]int, proposed [][2]int) [][2]int {
	final := make([][2]int, len(current))
	for i, query := range proposed {
		// Check if another proposed cell occupies the query
		occupied := false
		for j, other := range proposed {
			if j == i {
				// Don't check cell against itself
				continue
			}
			same := true
			for axis := range query {
				if query[axis] != other[axis] {
					// Mismatched axis, no elf at cell
					same = false
					break
				}
			}
			if same {
				occupied = true
				break
			}
		}
		if occupied {
			// If cells would occupy same cell, don't move
			for axis := range query {
				final[i][axis] = current[i][axis]
			}
		} else {
			// If no other would occupy the same cell, take it
			for axis := range query {
				final[i][axis] = proposed[i][axis]
			}
		}
	}
	return final
}

func Score(state [][2]int) int {
	size := len(state)
	imin, imax := state[0][0], state[0][0]
	jmin, jmax := state[0][1], state[0][1]
	for _, elf := range state {
		if elf[0] < imin {
			imin = elf[0]
		} else if elf[0] > imax {
			imax = elf[0]
		}
		if elf[1] < jmin {
			jmin = elf[1]
		} else if elf[1] > jmax {
			jmax = elf[1]
		}
	}
	return (imax-imin+1)*(jmax-jmin+1) - size
}
