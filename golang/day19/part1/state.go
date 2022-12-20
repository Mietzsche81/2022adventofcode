package part1

import (
	"fmt"
	"log"
	"strings"
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

func (src *State) Copy() State {
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

func (s *State) CanBuild(robot string) bool {
	if len(strings.TrimSpace(robot)) == 0 {
		return true
	}
	for resource, cost := range s.Blueprint[robot] {
		if s.Stockpile[resource] < cost {
			return false
		}
	}
	return true
}

func (x *State) Build(robot string) State {
	// If invalid input, fail
	if !x.CanBuild(robot) {
		log.Fatal(fmt.Errorf(
			"cannot build '%s'\n have: %v\n need: %v",
			robot, x.Stockpile, x.Blueprint[robot],
		))
	}
	// Initialize Output from current state
	y := x.Copy()
	// Current robots produce as time advances
	y.Produce()
	if len(strings.TrimSpace(robot)) > 0 {
		// Add valid robots if selected
		y.Producers[robot]++
		// Subtract the cost of the robot
		for resource, cost := range y.Blueprint[robot] {
			y.Stockpile[resource] -= cost
		}
	}
	return y
}

func (x *State) Produce() {
	x.Time++
	for resource, produced := range x.Producers {
		x.Stockpile[resource] += produced
	}
}
