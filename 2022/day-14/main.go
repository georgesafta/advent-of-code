package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

const SIZE = 1000

type point struct {
	x, y int
}

func main() {
	lines := util.ReadFile("./input/day-14.txt")
	m, minX, maxX, minY, maxY := parseInput(lines)

	fmt.Println("sand = ", addSand(m, minX, minY, maxX, maxY))

	m, minX, maxX, minY, maxY = parseInput(lines)
	maxY += 2
	for j := 0; j < SIZE; j++ {
		m[maxY][j] = 1
	}
	fmt.Println("sandTillTop = ", addSandTillTop(m, -1, minY, SIZE+1, maxY))
}

func addSand(m [][]int, minX, minY, maxX, maxY int) int {
	count := 0
	for i, j, _ := dfs(m, minX, minY, maxX, maxY, 500, 0); i >= minX && i <= maxX && j <= maxY; {
		m[j][i] = 1
		count++
		i, j, _ = dfs(m, minX, minY, maxX, maxY, 500, 0)
	}
	return count
}

func addSandTillTop(m [][]int, minX, minY, maxX, maxY int) int {
	count := 0
	for i, j, _ := dfs(m, minX, minY, maxX, maxY, 500, 0); i != 500 || j != 0; {
		m[j][i] = 1
		count++
		i, j, _ = dfs(m, minX, minY, maxX, maxY, 500, 0)
	}
	return count + 1
}

func dfs(m [][]int, minX, minY, maxX, maxY, x, y int) (int, int, bool) {
	if x < minX || x > maxX || y > maxY {
		return x, y, true
	}
	if m[y][x] == 1 {
		return -1, -1, false
	}

	i, j, success := dfs(m, minX, minY, maxX, maxY, x, y+1)
	if success {
		return i, j, true
	}
	i, j, success = dfs(m, minX, minY, maxX, maxY, x-1, y+1)
	if success {
		return i, j, true
	}
	i, j, success = dfs(m, minX, minY, maxX, maxY, x+1, y+1)
	if success {
		return i, j, true
	}

	return x, y, true
}

func parseInput(lines []string) ([][]int, int, int, int, int) {
	minX, minY, maxX, maxY := SIZE, SIZE, 0, 0
	m := make([][]int, SIZE)
	for i := 0; i < SIZE; i++ {
		m[i] = make([]int, SIZE)
	}

	for _, line := range lines {
		points := parseLine(line)
		x := points[0].x
		y := points[0].y
		minX = util.MinInt(minX, x)
		minY = util.MinInt(minY, y)
		maxX = util.MaxInt(maxX, x)
		maxY = util.MaxInt(maxY, y)
		for i := 1; i < len(points); i++ {
			curr := points[i]
			if curr.x == x {
				for i := util.MinInt(curr.y, y); i <= util.MaxInt(curr.y, y); i++ {
					m[i][x] = 1
				}
			} else {
				for j := util.MinInt(curr.x, x); j <= util.MaxInt(curr.x, x); j++ {
					m[y][j] = 1
				}
			}
			x = curr.x
			y = curr.y
			minX = util.MinInt(minX, x)
			minY = util.MinInt(minY, y)
			maxX = util.MaxInt(maxX, x)
			maxY = util.MaxInt(maxY, y)
		}
	}

	return m, minX, maxX, minY, maxY
}

func parseLine(line string) []point {
	points := []point{}

	pairs := strings.Fields(strings.ReplaceAll(line, "->", " "))
	for _, pair := range pairs {
		arr := strings.Split(pair, ",")
		x := util.Atoi(arr[0])
		y := util.Atoi(arr[1])
		points = append(points, point{x, y})
	}

	return points
}
