package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type point struct {
	x, y int
}

func (p *point) Add(w point) {
	p.x += w.x
	p.y += w.y
}

func (p *point) isTouching(w point) bool {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if p.x+dx == w.x && p.y+dy == w.y {
				return true
			}
		}
	}

	return false
}

func (p *point) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p *point) follow(w point) {
	if p.x == w.x {
		if p.y > w.y {
			p.y -= 1
		} else {
			p.y += 1
		}
	} else if p.y == w.y {
		if p.x > w.x {
			p.x -= 1
		} else {
			p.x += 1
		}
	} else {
		if p.x > w.x {
			p.x -= 1
		} else {
			p.x += 1
		}
		if p.y > w.y {
			p.y -= 1
		} else {
			p.y += 1
		}
	}
}

var dirs = map[string]point{
	"R": {x: 0, y: 1},
	"L": {x: 0, y: -1},
	"U": {x: -1, y: 0},
	"D": {x: 1, y: 0},
}

func main() {
	lines := util.ReadFile("./input/day-9.txt")

	fmt.Println("visited = ", simulate(lines, 1))
	fmt.Println("visited = ", simulate(lines, 9))
}

func simulate(lines []string, knots int) int {
	head := &point{x: 0, y: 0}
	followers := []*point{}
	for i := 0; i < knots; i++ {
		followers = append(followers, &point{x: 0, y: 0})
	}
	visited := map[string]bool{"0,0": true}
	for _, line := range lines {
		arr := strings.Fields(line)
		dir, _ := dirs[arr[0]]
		times := util.Atoi(arr[1])
		for i := 0; i < times; i++ {
			head.Add(dir)
			prev := head
			for _, tail := range followers {
				if !tail.isTouching(*prev) {
					tail.follow(*prev)
				}
				prev = tail
			}
			visited[prev.String()] = true
		}
	}

	return len(visited)
}
