package part1

import (
	"fmt"
)

type ComponentMap map[string]int
type BlueprintMap map[string]ComponentMap

type State struct {
	Time      int
	MaxTime   int
	Stockpile ComponentMap
	Producers ComponentMap
	Blueprint BlueprintMap
}

func (s *State) Initialize() {
	s.Time = 0
	s.Stockpile = initializeComponent()
	s.Producers = initializeComponent()
	s.Blueprint = make(BlueprintMap)
	for key := range s.Producers {
		s.Blueprint[key] = initializeComponent()
	}
	// default single ore miner
	s.Producers["ore"] = 1
	//
}

func (src State) Copy() State {
	dst := State{
		Time:      src.Time,
		MaxTime:   src.MaxTime,
		Blueprint: src.Blueprint,
	}
	dst.Stockpile = copyComponent(src.Stockpile)
	dst.Producers = copyComponent(src.Producers)

	return dst
}

func initializeComponent() ComponentMap {
	return ComponentMap{
		"ore":      0,
		"clay":     0,
		"obsidian": 0,
		"geode":    0,
	}
}

func copyComponent(src ComponentMap) ComponentMap {
	dst := make(ComponentMap)

	for key, val := range src {
		dst[key] = val
	}

	return dst
}

func PrintBluePrint(bp BlueprintMap) {
	for robot, costs := range bp {
		fmt.Printf("%s robot costs:\n", robot)
		for resource, cost := range costs {
			fmt.Printf("\t%3d %s\n", cost, resource)
		}
	}
}

func (s State) HaveResourcesToBuild(robot string) bool {
	for resource, cost := range s.Blueprint[robot] {
		if s.Stockpile[resource] < cost {
			return false
		}
	}
	return true
}

func (x State) Build(robot string) *State {
	// For each resource required to build
	for resource, cost := range x.Blueprint[robot] {
		// If cannot produce the required resource
		if x.Producers[resource] == 0 && cost > 0 {
			// Cannot build the requested, cannot advance state
			return nil
		}
	}

	// Initialize Output from current state
	y := x.Copy()
	// If don't yet have the resources
	for !y.HaveResourcesToBuild(robot) {
		// if out of time, failed to produce
		if y.Time >= y.MaxTime {
			return &y
		}
		// produce and advance time
		y.Produce()
	}
	// We now have the resources, must advance time before building
	y.Produce()
	// Now add the robot
	y.Producers[robot]++
	// Subtract the cost of the robot
	for resource, cost := range y.Blueprint[robot] {
		y.Stockpile[resource] -= cost
	}
	return &y
}

func (x *State) Produce() {
	x.Time++
	for resource, produced := range x.Producers {
		x.Stockpile[resource] += produced
	}
}
