package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

type point struct {
	x, y int
}

func (p point) add(dir []int) point {
	return point{p.x + dir[0], p.y + dir[1]}
}

var dirs map[rune][]int = map[rune][]int{'<': {0, -1}, '>': {0, 1}, '^': {-1, 0}, 'v': {1, 0}}

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Println("visited houses =", deliverPresents(lines[0], 1))
	fmt.Println("visited houses =", deliverPresents(lines[0], 2))
}

func deliverPresents(line string, workers int) int {
	if workers <= 0 {
		return 0
	}

	m := map[point]bool{}
	s := []point{}
	for i := 0; i < workers; i++ {
		s = append(s, point{0, 0})
	}
	m[s[0]] = true
	for i, r := range line {
		p := s[i%workers]
		p = p.add(dirs[r])
		s[i%workers] = p
		m[p] = true
	}

	return len(m)
}
