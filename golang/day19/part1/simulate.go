package part1

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func (x0 *State) OptimizeBFS(maxTime int) int {
	// Initialize
	x0.MaxTime = maxTime
	alive := []*State{x0}
	// Set up best for each step
	bestAtTime := make(map[int]State)
	for i := 0; i <= maxTime; i++ {
		bestAtTime[i] = x0.Copy()
	}
	// Use best at t=0 to store max number of resources needed
	for _, robot := range x0.Blueprint {
		for resource, cost := range robot {
			if cost > bestAtTime[0].Producers[resource] {
				bestAtTime[0].Producers[resource] = cost
			}
		}
	}

	// Optimize
	t := 0
	for len(alive) > 0 {
		// Advance
		tic := time.Now()
		next := []*State{}
		for _, s := range alive {
			next = append(next, s.PossibleAdvances()...)
		}
		// Cache best & Prune
		Prune(next, bestAtTime)
		// Iterate
		alive = next
		t++
		toc := time.Since(tic).Seconds()
		// Report
		fmt.Printf("t = %2d :: %5d NODES, %fs elapsed\n", t, len(alive), toc)
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
	}

	return bestAtTime[maxTime].Stockpile["geode"]
}

func (x *State) PossibleAdvances() (next []*State) {
	next = make([]*State, 0)
	if x.Time == x.MaxTime {
		// If at final time, done
		return
	}

	// Spend resources
	if x.CanBuild("geode") {
		// Priority #1: If can create geode producer, always do so
		geode := x.Build("geode")
		next = append(next, &geode)
	} else {
		// Otherwise explore other possibilities
		for _, robot := range []string{"obsidian", "clay", "ore", ""} {
			if x.CanBuild(robot) {
				y := x.Build(robot)
				next = append(next, &y)
			}
		}
	}

	return
}

func Prune(next []*State, bestAtTime map[int]State) {
	// Prune anything that is producing over the max needed
	// NOTE: we stored the max needed in bestAtTime[0].producers
	countPruned := 0
	for i, s := range next {
		for _, resource := range []string{"obsidian", "clay", "ore"} {
			if s.Producers[resource] > bestAtTime[0].Producers[resource] {
				next = append(next[:i], next[i+1:]...)
				countPruned++
			}
		}
	}

	// Cache the best
	for _, s := range next {
		if s.Stockpile["geode"] > bestAtTime[s.Time].Stockpile["geode"] {
			// If found new best, store as best
			bestAtTime[s.Time] = s.Copy()
		} else if s.Stockpile["geode"] == bestAtTime[s.Time].Stockpile["geode"] {
			// Tied for best
			for _, resource := range []string{"geode", "obsidian", "clay", "ore"} {
				// Break tie with the most productive
				if s.Producers[resource] > bestAtTime[s.Time].Producers[resource] {
					bestAtTime[s.Time] = s.Copy()
				}
			}
		}
	}

	// Prune anything that is falling behind.
	for i, s := range next {
		if s.Stockpile["geode"] < bestAtTime[s.Time].Stockpile["geode"] {
			// If behind in geode production, assume will never catch up
			next = append(next[:i], next[i+1:]...)
			countPruned++
		}
	}
	fmt.Printf("\t --- Pruned %d\n", countPruned)
}
