package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type valve struct {
	flowRate    int
	connections []string
}

type position struct {
	time int
	name string
}

type state struct {
	time   int
	name   string
	total  int
	opened map[string]bool
}

func main() {
	lines := util.ReadFile("./input/day-16.txt")

	valves := parseLines(lines)
	fmt.Println("max = ", bfs(valves))
}

func parseLines(lines []string) map[string]valve {
	m := map[string]valve{}
	for _, line := range lines {
		arr := strings.Fields(line)
		name := arr[1]
		flowRate := util.Atoi(arr[4][5 : len(arr[4])-1])
		connections := arr[9:]
		for i, e := range connections {
			if e[len(e)-1] == ',' {
				connections[i] = e[:len(e)-1]
			}
		}
		m[name] = valve{flowRate, connections}
	}

	return m
}

func bfs(m map[string]valve) int {
	max := 0
	seen := map[position]int{}
	states := []state{}
	states = append(states, state{1, "AA", 0, map[string]bool{}})

	for len(states) > 0 {
		curr := states[0]

		if v, exists := seen[position{curr.time, curr.name}]; exists && v >= curr.total {
			states = states[1:]
			continue
		}
		seen[position{curr.time, curr.name}] = curr.total

		if curr.time == 30 {
			max = util.MaxInt(max, curr.total)
			states = states[1:]
			continue
		}

		if m[curr.name].flowRate > 0 && !curr.opened[curr.name] {
			mapCopy := copyMap(curr.opened)
			mapCopy[curr.name] = true
			total := curr.total + sum(m, mapCopy)
			newState := state{curr.time + 1, curr.name, total, mapCopy}
			states = append(states, newState)
		}

		total := curr.total + sum(m, curr.opened)
		for _, v := range m[curr.name].connections {
			newState := state{curr.time + 1, v, total, copyMap(curr.opened)}
			states = append(states, newState)
		}

		states = states[1:]
	}

	return max
}

func copyMap(m map[string]bool) map[string]bool {
	copy := map[string]bool{}
	for k, v := range m {
		copy[k] = v
	}

	return copy
}

func sum(m map[string]valve, opened map[string]bool) int {
	sum := 0
	for k := range opened {
		sum += m[k].flowRate
	}

	return sum
}
