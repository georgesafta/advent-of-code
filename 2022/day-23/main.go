package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

type cell struct {
	pos  point
	next point
}

type point struct {
	row int
	col int
}

func main() {
	lines := util.ReadFile("./input/day-23.txt")
	cells, grid := parse(lines)
	fmt.Println("simulation = ", simulate(cells, grid, 10))
	cells, grid = parse(lines)
	fmt.Println("part2 = ", part2(cells, grid))
}

func simulate(cells []cell, grid map[point]bool, times int) int {
	moves := moves()
	for i := 0; i < times; i++ {
		set := map[point]int{}
		for j, c := range cells {
			nextPos := c.pos
			if hasNeighbour(c.pos, grid) {
				nextPos = next(c.pos, grid, moves)
			}
			c.next = nextPos
			cells[j] = c
			set[nextPos] = set[nextPos] + 1
		}

		newGrid := map[point]bool{}
		for j, c := range cells {
			if set[c.next] == 1 {
				c.pos = c.next
			}
			c.next = point{}
			newGrid[c.pos] = true
			cells[j] = c
		}

		grid = newGrid

		first := moves[0]
		moves = moves[1:]
		moves = append(moves, first)
	}

	xMin := 1_000_000
	xMax := -1_000_000
	yMin := 1_000_000
	yMax := -1_000_000
	for p := range grid {
		xMin = util.MinInt(xMin, p.row)
		xMax = util.MaxInt(xMax, p.row)
		yMin = util.MinInt(yMin, p.col)
		yMax = util.MaxInt(yMax, p.col)
	}

	return (util.AbsInt(xMax-xMin)+1)*(util.AbsInt(yMax-yMin)+1) - len(cells)
}

func part2(cells []cell, grid map[point]bool) int {
	moves := moves()
	iteratios := 0
	for {
		changes := 0
		set := map[point]int{}
		for j, c := range cells {
			nextPos := c.pos
			if hasNeighbour(c.pos, grid) {
				nextPos = next(c.pos, grid, moves)
				if c.pos.row != nextPos.row || c.pos.col != nextPos.col {
					changes++
				}
			}
			c.next = nextPos
			cells[j] = c
			set[nextPos] = set[nextPos] + 1
		}

		newGrid := map[point]bool{}
		for j, c := range cells {
			if set[c.next] == 1 {
				c.pos = c.next
			}
			c.next = point{}
			newGrid[c.pos] = true
			cells[j] = c
		}

		iteratios++
		if changes == 0 {
			break
		}

		grid = newGrid

		first := moves[0]
		moves = moves[1:]
		moves = append(moves, first)
	}

	return iteratios
}

func hasNeighbour(pos point, grid map[point]bool) bool {
	for row := pos.row - 1; row <= pos.row+1; row++ {
		for col := pos.col - 1; col <= pos.col+1; col++ {
			if row == pos.row && col == pos.col {
				continue
			}
			p := point{row, col}
			if grid[p] {
				return true
			}
		}
	}

	return false
}

func parse(lines []string) ([]cell, map[point]bool) {
	cells := []cell{}
	grid := map[point]bool{}
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				pos := point{i, j}
				cell := cell{pos, point{}}
				cells = append(cells, cell)
				grid[pos] = true
			}
		}
	}

	return cells, grid
}

func moves() []func(point, map[point]bool) (point, bool) {
	moves := []func(point, map[point]bool) (point, bool){}
	moves = append(moves, func(curr point, grid map[point]bool) (point, bool) {
		N := point{curr.row - 1, curr.col}
		NW := point{curr.row - 1, curr.col - 1}
		NE := point{curr.row - 1, curr.col + 1}
		if !grid[N] && !grid[NW] && !grid[NE] {
			return N, true
		}

		return curr, false
	})

	moves = append(moves, func(curr point, grid map[point]bool) (point, bool) {
		S := point{curr.row + 1, curr.col}
		SW := point{curr.row + 1, curr.col - 1}
		SE := point{curr.row + 1, curr.col + 1}
		if !grid[S] && !grid[SW] && !grid[SE] {
			return S, true
		}

		return curr, false
	})

	moves = append(moves, func(curr point, grid map[point]bool) (point, bool) {
		W := point{curr.row, curr.col - 1}
		NW := point{curr.row - 1, curr.col - 1}
		SW := point{curr.row + 1, curr.col - 1}
		if !grid[W] && !grid[NW] && !grid[SW] {
			return W, true
		}

		return curr, false
	})

	moves = append(moves, func(curr point, grid map[point]bool) (point, bool) {
		E := point{curr.row, curr.col + 1}
		NE := point{curr.row - 1, curr.col + 1}
		SE := point{curr.row + 1, curr.col + 1}
		if !grid[E] && !grid[NE] && !grid[SE] {
			return E, true
		}

		return curr, false
	})

	return moves
}

func next(curr point, grid map[point]bool, moves []func(point, map[point]bool) (point, bool)) point {
	for _, m := range moves {
		if p, success := m(curr, grid); success {
			return p
		}
	}

	return curr
}
