package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type reindeer struct {
	name   string
	speed  int
	run    int
	rest   int
	dist   int
	points int
}

func (r *reindeer) distance(seconds int) int {
	cycleTime := r.run + r.rest
	cycles := seconds / cycleTime

	remaining := seconds - cycles*cycleTime

	return cycles*r.run*r.speed + util.MinInt(remaining, r.run)*r.speed
}

func (r *reindeer) move(second int) {
	if second%(r.run+r.rest) < r.run {
		r.dist += r.speed
	}
}

func main() {
	lines := util.ReadFile("input.txt")
	reindeers := parse(lines)
	max := 0
	for _, r := range reindeers {
		max = util.MaxInt(max, r.distance(2503))
	}
	fmt.Println("part1 =", max)

	max = simulate(reindeers, 2503)
	fmt.Println("part2 =", max)
}

func parse(lines []string) []*reindeer {
	r := []*reindeer{}
	for _, line := range lines {
		arr := strings.Fields(line)
		r = append(r, &reindeer{arr[0], util.Atoi(arr[3]), util.Atoi(arr[6]), util.Atoi(arr[13]), 0, 0})
	}

	return r
}

func simulate(reindeers []*reindeer, turns int) int {
	for i := 0; i < turns; i++ {
		max := 0
		l := []*reindeer{}
		for _, r := range reindeers {
			r.move(i)
			if r.dist > max {
				max = r.dist
				l = []*reindeer{r}
			} else if r.dist == max {
				l = append(l, r)
			}
		}

		for _, r := range l {
			r.points++
		}
	}

	max := 0
	for _, r := range reindeers {
		max = util.MaxInt(max, r.points)
	}

	return max
}
