package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

var dirs [][]int = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {0, 0}}
var moves map[rune][]int = map[rune][]int{'>': {0, 1}, '<': {0, -1}, '^': {-1, 0}, 'v': {1, 0}}

type point struct {
	x, y int
}

type cell struct {
	pos    point
	symbol rune
}

func main() {
	lines := util.ReadFile("./input/day-24.txt")
	m, cells, start, end := parse(lines)
	stepsToEnd := bfs(m, cells, start, end, 0)
	fmt.Println("bfs = ", stepsToEnd)
	//stepsToStart := bfs(m, cells, end, start, stepsToEnd)
	//toEnd := bfs(m, cells, start, end, stepsToStart)
	//fmt.Println("part2 = ", toEnd)
}

func bfs(m [][]int, cells []cell, start, end point, steps int) int {
	queue := []point{}
	queue = append(queue, start)
	visited := map[string]bool{}
	visited["["+toString(start.x)+","+toString(start.y)+"]"+state(m, cells)] = true
	for len(queue) > 0 {
		m, cells = nextState(m, cells)
		state := state(m, cells)
		size := len(queue)
		for i := 0; i < size; i++ {
			p := queue[i]
			if p.x == end.x && p.y == end.y {
				return steps
			}
			for _, dir := range dirs {
				x := p.x + dir[0]
				y := p.y + dir[1]
				currState := "[" + toString(x) + "," + toString(y) + "]" + state
				if x >= 0 && x < len(m) && y >= 0 && y < len(m[0]) && m[x][y] == 0 && !visited[currState] {
					queue = append(queue, point{x, y})
					visited[currState] = true
				}
			}
		}
		steps++
		queue = queue[size:]
	}

	panic("Should return in for")
}

func state(m [][]int, cells []cell) string {
	s := ""
	set := map[point][]int{}
	for _, cell := range cells {
		if _, exists := set[cell.pos]; !exists {
			set[cell.pos] = make([]int, 4)
		}
		set[cell.pos][toInt(cell.symbol)]++
	}

	for i, line := range m {
		for j, v := range line {
			if v > 0 {
				p := point{i, j}
				arr := set[p]
				s += "[" + toString(i) + "," + toString(j) + "," + toString(arr[0]) + "," + toString(arr[1]) + "," + toString(arr[2]) + "," + toString(arr[3]) + "]"
			}
		}
	}

	return s
}

func nextState(m [][]int, cells []cell) ([][]int, []cell) {
	length := len(m)
	width := len(m[0])
	for i, cell := range cells {
		x := cell.pos.x
		y := cell.pos.y
		m[x][y] -= 1
		move := moves[cell.symbol]
		nextX := x + move[0]
		nextY := y + move[1]
		if nextX < 0 {
			nextX = length - 1
		}
		if nextY < 0 {
			nextY = width - 1
		}
		nextX %= length
		nextY %= width
		for m[nextX][nextY] < 0 {
			nextX += move[0]
			nextY += move[1]
			if nextX < 0 {
				nextX = length - 1
			}
			if nextY < 0 {
				nextY = width - 1
			}
			nextX %= length
			nextY %= width
		}
		m[nextX][nextY]++
		cell.pos = point{nextX, nextY}
		cells[i] = cell
	}

	return m, cells
}

func parse(lines []string) ([][]int, []cell, point, point) {
	m := make([][]int, len(lines))
	cells := []cell{}
	var start point
	var end point
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				m[i] = append(m[i], -1)
			} else if r == '.' {
				m[i] = append(m[i], 0)
				if i == 0 {
					start = point{i, j}
				}
				if i == len(lines)-1 {
					end = point{i, j}
				}
			} else {
				m[i] = append(m[i], 1)
				cells = append(cells, cell{point{i, j}, r})
			}
		}
	}

	return m, cells, start, end
}

func toInt(r rune) int {
	result := -1
	switch r {
	case '>':
		result = 0
		break
	case '<':
		result = 1
		break
	case '^':
		result = 2
		break
	case 'v':
		result = 3
		break
	}

	return result
}

func toString(n int) string {
	return fmt.Sprint(n)
}
