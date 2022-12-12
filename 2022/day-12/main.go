package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

type cell struct {
	x, y int
}

var dirs [][]int = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func main() {
	m, s, e := readMatrix()
	min, _ := bfs(m, s, e)
	fmt.Println("part1 = ", min)

	for i, row := range m {
		for j, v := range row {
			if v == 0 {
				curr, err := bfs(m, cell{x: i, y: j}, e)
				if err == nil {
					min = util.MinInt(min, curr)
				}
			}
		}
	}

	fmt.Println("part2 = ", min)
}

func readMatrix() ([][]int, cell, cell) {
	lines := util.ReadFile("./input/day-12.txt")
	matrix := make([][]int, len(lines))
	start := cell{}
	end := cell{}
	for i, line := range lines {
		for j, e := range line {
			if e == 'E' {
				end.x = i
				end.y = j
				e = 'z'
			} else if e == 'S' {
				start.x = i
				start.y = j
				e = 'a'
			}
			matrix[i] = append(matrix[i], int(e-'a'))
		}
	}

	return matrix, start, end
}

func bfs(m [][]int, s, e cell) (int, error) {
	rows := len(m)
	cols := len(m[0])
	visited := make([][]bool, len(m))
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			visited[i] = append(visited[i], false)
		}
	}

	queue := []cell{}
	queue = append(queue, s)
	visited[s.x][s.y] = true
	count := 0
	for len(queue) > 0 {
		for size := len(queue); size > 0; size-- {
			curr := queue[0]
			if e.x == curr.x && e.y == curr.y {
				return count, nil
			}

			value := m[curr.x][curr.y]
			for _, dir := range dirs {
				row := curr.x + dir[0]
				col := curr.y + dir[1]

				if row < 0 || row >= rows || col < 0 || col >= cols || visited[row][col] || m[row][col] > value+1 {
					continue
				}

				queue = append(queue, cell{x: row, y: col})
				visited[row][col] = true
			}
			queue = queue[1:]
		}
		count++
	}

	return -1, fmt.Errorf("Can't reach")
}
