package main

import (
	"fmt"
	"strconv"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	input := util.ReadFile("./input/day-10.txt")
	part1(input)
	part2(input)
}

func part1(lines []string) {
	total := 0
	cycle := 1
	registry := 1

	for _, line := range lines {
		arr := strings.Fields(line)
		if arr[0] == "noop" {
			cycle += 1
		} else {
			v, _ := strconv.Atoi(arr[1])
			cycle += 1
			if (cycle-20)%40 == 0 && cycle <= 220 {
				total += cycle * registry
			}
			cycle += 1
			registry += v
		}
		if (cycle-20)%40 == 0 && cycle <= 220 {
			total += cycle * registry
		}
	}

	fmt.Println("total = ", total)
}

func part2(lines []string) {
	cycle := 0
	registry := 1

	for _, line := range lines {
		arr := strings.Fields(line)
		cycle += 1
		printCRT(cycle, registry%40)

		if arr[0] == "noop" {
			continue
		}

		v, _ := strconv.Atoi(arr[1])
		cycle += 1
		printCRT(cycle, registry%40)
		registry += v
	}
}

func printCRT(cycle, pos int) {
	prev := (cycle - 1) % 40
	if prev == 0 {
		fmt.Println()
	}

	if pos >= prev-1 && pos <= prev+1 {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}
