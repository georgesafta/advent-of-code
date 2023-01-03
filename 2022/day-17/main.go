package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

const (
	WIDTH  = 7
	HEIGHT = 9000
)

type point struct {
	x, y int
}

type shape struct {
	points []point
}

func (s shape) up(v int) shape {
	p := []point{}
	for _, pp := range s.points {
		p = append(p, point{pp.x, pp.y + v})
	}

	return shape{p}
}

func (s shape) left(grid [][]bool) shape {
	p := []point{}
	for _, pp := range s.points {
		x := pp.x - 1
		if x < 0 || grid[pp.y][x] {
			return s
		}
		p = append(p, point{x, pp.y})
	}

	return shape{p}
}

func (s shape) right(grid [][]bool) shape {
	p := []point{}
	for _, pp := range s.points {
		x := pp.x + 1
		if x >= WIDTH || grid[pp.y][x] {
			return s
		}
		p = append(p, point{x, pp.y})
	}

	return shape{p}
}

func (s shape) down(grid [][]bool) (shape, bool) {
	p := []point{}
	for _, pp := range s.points {
		y := pp.y - 1
		if y < 0 || grid[y][pp.x] {
			return s, false
		}
		p = append(p, point{pp.x, y})
	}

	return shape{p}, true
}

func main() {
	lines := util.ReadFile("./input/day-17.txt")
	fmt.Println("height = ", tetrisHeight(lines[0], 2022))
}

func tetrisHeight(wind string, iterations int) int {
	height := 0
	shapes := generateShapes()
	grid := make([][]bool, HEIGHT)
	for i := 0; i < HEIGHT; i++ {
		grid[i] = make([]bool, WIDTH)
	}

	j := 0
	for i := 0; i < iterations; i++ {
		s := shapes[i%len(shapes)]
		s = s.up(height + 3)
		canContinue := true
		for canContinue {
			w := wind[j%len(wind)]
			s = simulateWind(s, w, grid)
			ss, success := s.down(grid)
			if success {
				s = ss
				j++
			}
			canContinue = success
		}

		for _, p := range s.points {
			height = util.MaxInt(height, p.y+1)
			grid[p.y][p.x] = true
		}

		j++
	}

	return height
}

func simulateWind(s shape, direction byte, grid [][]bool) shape {
	if direction == '>' {
		return s.right(grid)
	}

	return s.left(grid)
}

func generateShapes() []shape {
	shapes := []shape{}
	shapes = append(shapes, shape{[]point{{2, 0}, {3, 0}, {4, 0}, {5, 0}}})
	shapes = append(shapes, shape{[]point{{2, 1}, {3, 0}, {3, 1}, {4, 1}, {3, 2}}})
	shapes = append(shapes, shape{[]point{{2, 0}, {3, 0}, {4, 0}, {4, 1}, {4, 2}}})
	shapes = append(shapes, shape{[]point{{2, 0}, {2, 1}, {2, 2}, {2, 3}}})
	shapes = append(shapes, shape{[]point{{2, 0}, {3, 0}, {2, 1}, {3, 1}}})

	return shapes
}
