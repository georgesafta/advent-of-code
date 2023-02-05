package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	m := parse(lines)
	fmt.Println(count(simulate(m, 100, next)))
	m[0][0] = 1
	m[0][len(m[0])-1] = 1
	m[len(m)-1][0] = 1
	m[len(m)-1][len(m[0])-1] = 1
	fmt.Println(count(simulate(m, 100, nextWithCornersOn)))
}

func count(m [][]int) int {
	sum := 0
	for _, row := range m {
		for _, v := range row {
			if v == 1 {
				sum++
			}
		}
	}

	return sum
}

func simulate(m [][]int, turns int, nextState func(int, int, [][]int) int) [][]int {
	for turn := 0; turn < turns; turn++ {
		nextTurn := make([][]int, len(m))
		for i, row := range m {
			for j := range row {
				nextTurn[i] = append(nextTurn[i], nextState(i, j, m))
			}
		}
		m = nextTurn
	}

	return m
}

func parse(lines []string) [][]int {
	m := make([][]int, len(lines))
	for row, line := range lines {
		for _, r := range line {
			val := 0
			if r == '#' {
				val = 1
			}
			m[row] = append(m[row], val)
		}
	}

	return m
}

func next(row, col int, m [][]int) int {
	isAlive := false
	if m[row][col] == 1 {
		isAlive = true
	}

	aliveNeighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			r := row + i
			c := col + j
			if (i == 0 && j == 0) || r < 0 || c < 0 || r >= len(m) || c >= len(m[0]) {
				continue
			}
			if m[r][c] == 1 {
				aliveNeighbours++
			}
		}
	}

	result := 0
	if (isAlive && aliveNeighbours >= 2 && aliveNeighbours <= 3) || (!isAlive && aliveNeighbours == 3) {
		result = 1
	}

	return result
}

func nextWithCornersOn(row, col int, m [][]int) int {
	val := next(row, col, m)
	if isCorner(row, col, m) {
		val = 1
	}

	return val
}

func isCorner(row, col int, m [][]int) bool {
	l := len(m)
	h := len(m[0])

	return (row == 0 && col == 0) || (row == 0 && col == h-1) || (row == l-1 && col == 0) || (row == l-1 && col == h-1)
}
