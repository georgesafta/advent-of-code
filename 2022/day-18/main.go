package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type cube struct {
	x, y, z int
}

func (c cube) neighbors() []cube {
	return []cube{
		{c.x + 1, c.y, c.z},
		{c.x - 1, c.y, c.z},
		{c.x, c.y + 1, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y, c.z + 1},
		{c.x, c.y, c.z - 1},
	}
}

func (c cube) isOutside(limit cube) bool {
	return c.x < -1 || c.x > limit.x+1 || c.y < -1 || c.y > limit.y+1 || c.z < -1 || c.z > limit.z+1
}

func main() {
	lines := util.ReadFile("./input/day-18.txt")
	cubes := getCubes(lines)
	fmt.Println("surface = ", part1(cubes))
	fmt.Println("surface = ", part2(cubes))
}

func part1(cubes []cube) int {
	m := map[cube]bool{}
	for _, cube := range cubes {
		m[cube] = true
	}

	total := 0
	for _, cube := range cubes {
		surface := 0
		for _, n := range cube.neighbors() {
			if !m[n] {
				surface++
			}
		}

		total += surface
	}

	return total
}

func part2(cubes []cube) int {
	m := map[cube]bool{}
	limit := cube{}
	for _, cube := range cubes {
		m[cube] = true
		limit.x = util.MaxInt(limit.x, cube.x)
		limit.y = util.MaxInt(limit.y, cube.y)
		limit.z = util.MaxInt(limit.z, cube.z)
	}

	return countExterior(cube{0, 0, 0}, limit, m, map[cube]bool{})
}

func countExterior(c, limit cube, used, exterior map[cube]bool) int {
	if c.isOutside(limit) || exterior[c] {
		return 0
	}
	if used[c] {
		return 1
	}
	exterior[c] = true

	count := 0
	for _, n := range c.neighbors() {
		count += countExterior(n, limit, used, exterior)
	}

	return count
}

func getCubes(lines []string) []cube {
	cubes := []cube{}
	for _, line := range lines {
		cubes = append(cubes, processLine(line))
	}

	return cubes
}

func processLine(line string) cube {
	arr := strings.Split(line, ",")
	x, _ := strconv.Atoi(arr[0])
	y, _ := strconv.Atoi(arr[1])
	z, _ := strconv.Atoi(arr[2])

	return cube{x, y, z}
}
