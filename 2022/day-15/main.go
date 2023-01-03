package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type point struct {
	x, y int
}

type scan struct {
	sensor, beacon point
	distance       int
}

func (p point) isEqual(v point) bool {
	return p.x == v.x && p.y == v.y
}

func main() {
	lines := util.ReadFile("./input/day-15.txt")
	fmt.Println("occupied = ", part1(lines, 2000000))
	fmt.Println("frequency = ", part2(lines))
}

func part1(lines []string, y int) int {
	scans := []scan{}
	max := 0
	min := 9999999
	for _, line := range lines {
		s, b := parseLine(line)
		d := dist(s, b)
		scans = append(scans, scan{s, b, d})
		if s.x-d < min {
			min = s.x - d
		}
		if s.x+d > max {
			max = s.x + d
		}
	}

	occupied := 0
	p := point{min, y}
	for x := min; x <= max; x++ {
		p.x = x
		for _, s := range scans {
			if s.distance >= dist(s.sensor, p) && !p.isEqual(s.sensor) && !p.isEqual(s.beacon) {
				occupied++
				break
			}
		}
	}

	return occupied
}

func parseLine(line string) (point, point) {
	arr := strings.Fields(line)
	x := util.Atoi(arr[2][2 : len(arr[2])-1])
	y := util.Atoi(arr[3][2 : len(arr[3])-1])
	sensor := point{x, y}

	x = util.Atoi(arr[8][2 : len(arr[8])-1])
	y = util.Atoi(arr[9][2:])
	beacon := point{x, y}

	return sensor, beacon
}

func dist(a, b point) int {
	return util.AbsInt(a.x-b.x) + util.AbsInt(a.y-b.y)
}

func part2(lines []string) int {
	scans := []scan{}
	for _, line := range lines {
		s, b := parseLine(line)
		d := dist(s, b)
		scans = append(scans, scan{s, b, d})
	}

	for _, scan := range scans {
		distPlusOne := scan.distance + 1

		for r := -distPlusOne; r <= distPlusOne; r++ {
			x := scan.sensor.x + r

			if x < 0 {
				continue
			}
			if x > 4000000 {
				break
			}

			offset := distPlusOne - util.AbsInt(r)
			leftY := scan.sensor.y - offset
			rightY := scan.sensor.y + offset

			if isValidCoord(leftY) && !isReachable(scans, point{leftY, x}) {
				return leftY*4000000 + x
			}
			if isValidCoord(rightY) && !isReachable(scans, point{rightY, x}) {
				return rightY*4000000 + x
			}
		}
	}
	panic("unreachable")
}

func isValidCoord(n int) bool {
	return n >= 0 && n <= 4000000
}

func isReachable(scans []scan, p point) bool {
	for _, scan := range scans {
		if scan.distance >= dist(p, scan.sensor) {
			return true
		}
	}

	return false
}
