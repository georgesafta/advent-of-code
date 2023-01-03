package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func main() {
	lines := util.ReadFile("./input/day-22.txt")
	m := matrix(lines)
	fmt.Println("simulation = ", simulate(m, lines[len(lines)-1], 0))
}

func simulate(m [][]int, line string, dir int) int {
	curr := 0
	row, col := getStart(m)
	for _, v := range line {
		if v == 'R' || v == 'L' {
			row, col = move(row, col, curr, m, dirs[dir])
			curr = 0
			if v == 'R' {
				dir++
			}
			if v == 'L' {
				dir--
			}
			dir += len(dirs)
			dir %= len(dirs)
		} else {
			curr *= 10
			curr += int(v - '0')
		}
	}
	if curr > 0 {
		row, col = move(row, col, curr, m, dirs[dir])
	}

	return (row+1)*1000 + (col+1)*4 + dir
}

func getStart(m [][]int) (int, int) {
	for j := 0; j < len(m[0]); j++ {
		if m[0][j] == 0 {
			return 0, j
		}
	}

	return -1, -1
}

func move(row, col, moves int, m [][]int, dir []int) (int, int) {
	length := len(m)
	width := len(m[0])
	nextRow, nextCol := nextIndex(row, col, length, width, dir)
	for moves > 0 {
		for m[nextRow][nextCol] == -1 {
			nextRow, nextCol = nextIndex(nextRow, nextCol, length, width, dir)
		}
		if m[nextRow][nextCol] == 1 {
			break
		}
		row, col = nextRow, nextCol
		nextRow, nextCol = nextIndex(nextRow, nextCol, length, width, dir)
		moves--
	}

	return row, col
}

func nextIndex(row, col, length, width int, dir []int) (int, int) {
	nextRow, nextCol := (row+dir[0])%length, (col+dir[1])%width
	if nextRow < 0 {
		nextRow += length
	}
	if nextCol < 0 {
		nextCol += width
	}

	return nextRow, nextCol
}

func matrix(lines []string) [][]int {
	width := width(lines)
	length := len(lines) - 2
	m := make([][]int, length)
	for i := 0; i < length; i++ {
		m[i] = make([]int, width)
		for j := range m[i] {
			m[i][j] = -1
		}
		for j, r := range lines[i] {
			if r == '.' {
				m[i][j] = 0
			}
			if r == '#' {
				m[i][j] = 1
			}
		}
	}

	return m
}

func width(lines []string) int {
	width := 0
	for i := 0; i < len(lines)-2; i++ {
		width = util.MaxInt(width, len(lines[i]))
	}

	return width
}
