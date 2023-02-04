package main

import (
	"fmt"
	"sort"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	s := parse(lines)
	sort.Ints(s)
	m := map[int]int{}
	fmt.Println(numWays(150, 0, s, m))
	minParts := 150
	v := 0
	for k, val := range m {
		if k < minParts {
			minParts = k
			v = val
		}
	}
	fmt.Println("min parts count =", v)
}

func parse(lines []string) []int {
	s := []int{}
	for _, line := range lines {
		s = append(s, util.Atoi(line))
	}

	return s
}

func numWays(liters, used int, containers []int, seen map[int]int) int {
	if liters == 0 {
		occ := 0
		if v, exists := seen[used]; exists {
			occ = v
		}
		occ++
		seen[used] = occ
		return 1
	}
	if liters < 0 || len(containers) == 0 {
		return 0
	}

	sum := 0
	for i := 0; i < len(containers); i++ {
		if containers[i] > liters {
			break
		}
		sum += numWays(liters-containers[i], used+1, containers[i+1:], seen)
	}

	return sum
}
