package part1

import (
	"fmt"
	"time"
)

var bestAtTime = map[int]State{}
var compressedCache = map[int]bool{}

func (x0 *State) Optimize(maxTime int) int {
	// Initialize
	x0.MaxTime = maxTime
	compressedCache = make(map[int]bool)
	alive := []*State{x0}
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
	n := 0
	for len(alive) > 0 {
		// Advance
		tic := time.Now()
		next := []*State{}
		for _, s := range alive {
			next = append(next, s.PossibleAdvances()...)
		}
		// Iterate
		alive = next
		n++
		// Score the best
		ScoreBest(alive)
		toc := time.Since(tic).Seconds()
		// Report
		fmt.Printf("cycle = %2d :: %5d NODES, %fs elapsed\t", n, len(alive), toc)
		fmt.Printf("BEST @ t=%2d:\tt\tgeodes = %2d\tgeodebots = %2d\n",
			maxTime, bestAtTime[maxTime].Stockpile["geode"], bestAtTime[maxTime].Producers["geode"])
	}

	return bestAtTime[maxTime].Stockpile["geode"]
}

func (x *State) PossibleAdvances() (next []*State) {
	// Step 0: initialize output
	next = make([]*State, 0)

	// Step 1: check for pruning
	if x.Time >= x.MaxTime {
		// If at final time, done
		return
	} else if CheckCache(x) {
		// Already searched, moot point
		return
	} else if Prune(x) {
		// Suboptimal, moot
		return
	}

	// Step 2: make a decision & advance
	if x.HaveResourcesToBuild("geode") {
		// Priority #1: If can create geode producer, always do so immediately
		geode := x.Build("geode")
		next = append(next, geode)
	} else {
		// Priority #2: build another producer
		for robot := range x.Blueprint {
			y := x.Build(robot)
			if y != nil {
				next = append(next, y)
			}
		}
	}
	return
}

func ScoreBest(next []*State) {
	for _, s := range next {
		if s.Stockpile["geode"] > bestAtTime[s.Time].Stockpile["geode"] {
			// If found new undisputed best, store as best
			bestAtTime[s.Time] = s.Copy()
		} else if s.Stockpile["geode"] == bestAtTime[s.Time].Stockpile["geode"] {
			// Break tie with the most productive (in order of degree)
			if s.Producers["geode"] > bestAtTime[s.Time].Producers["geode"] {
				bestAtTime[s.Time] = s.Copy()
			}
		}
	}
	// Propagate best forward
	for i := 1; i < bestAtTime[0].MaxTime; i++ {
		possible := bestAtTime[i].Copy()
		for possible.Time < possible.MaxTime {
			possible.Produce()
			if possible.Stockpile["geode"] > bestAtTime[possible.Time].Stockpile["geode"] {
				bestAtTime[possible.Time] = possible.Copy()
			}
		}
	}
}

func Prune(x *State) bool {
	// Prune anything that is producing over the max needed
	// NOTE: we stored the max needed in bestAtTime[0].producers
	for _, resource := range []string{"obsidian", "clay", "ore"} {
		if x.Producers[resource] > bestAtTime[0].Producers[resource] {
			return true
		}
	}

	// Prune anything that is falling behind & can't catch up even after adding an additional geodebot
	if x.Stockpile["geode"] < (bestAtTime[x.Time].Stockpile["geode"] - x.Producers["geode"] - 1) {
		return true
	}

	return false
}

func Encode(s *State) (encoding int) {
	// encode
	encoding = s.Time
	resource := []string{"geode", "obsidian", "clay", "ore"}
	for i, j := 0, 100; i < len(resource); i, j = i+1, j*10000 {
		encoding += j * s.Stockpile[resource[i]]
		encoding += j * s.Producers[resource[i]] * 100
	}
	return
}

func CheckCache(s *State) bool {
	encoding := Encode(s)
	_, alreadyCached := compressedCache[encoding]
	compressedCache[encoding] = true
	return alreadyCached
}

func Remove(arr []*State, i int) []*State {
	if i < len(arr)-1 {
		return append(arr[:i], arr[i+1:]...)
	} else {
		return arr[:i]
	}
}
