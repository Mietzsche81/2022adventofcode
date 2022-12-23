package part1

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

func Process(data [][2]int, iterations int) [][2]int {
	state := make([][2]int, len(data))
	copy(state, data)

	for i := 0; i < iterations; i++ {
		fmt.Printf("ITERATION %d\n", i)
		fmt.Println(rules)
		Print(state)
		fmt.Printf("--------------\n")
		next := Propose(state)
		state = Negotiate(state, next)
		CycleDirections()
	}
	fmt.Printf("FINAL\n")
	Print(state)
	fmt.Printf("--------------\n")

	return state
}

func Propose(state [][2]int) [][2]int {
	proposed := make([][2]int, len(state))
	copy(proposed, state)

	for i, elf := range proposed {
		fmt.Printf("QUERY: %d %d\n", elf[0], elf[1])
		for _, move := range rules {
			if ValidMove(state, elf, move) {
				for axis := range elf {
					proposed[i][axis] += directions[move][axis]
				}
				fmt.Printf("Propose move %s to %d %d\n", move, proposed[i][0], proposed[i][1])
				break
			}
			fmt.Printf("Can't move %s\n", move)
		}
	}

	return proposed
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
			fmt.Printf("@[%d %d] Negotiate [%d %d] v [%d %d] ",
				current[i][0], current[i][1],
				query[0], query[1],
				other[0], other[1],
			)
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
			fmt.Printf("CONFLICT, stay at [%d %d]\n", final[i][0], final[i][1])
		} else {
			// If no other would occupy the same cell, take it
			for axis := range query {
				final[i][axis] = proposed[i][axis]
			}
			fmt.Printf(" .  FREE, move to [%d %d]\n", final[i][0], final[i][1])
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
	fmt.Println(imax-imin, jmax-jmin, size)
	return (imax-imin+1)*(jmax-jmin+1) - size
}
