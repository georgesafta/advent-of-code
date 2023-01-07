package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Println("floor =", floor(lines))
	fmt.Println("basement position =", basementPosition(lines[0]))
}

func floor(lines []string) int {
	floor := 0
	for _, line := range lines {
		for _, r := range line {
			if r == '(' {
				floor++
			} else {
				floor--
			}
		}
	}

	return floor
}

func basementPosition(line string) int {
	floor := 0
	for i, r := range line {
		if r == '(' {
			floor++
		} else {
			floor--
		}
		if floor < 0 {
			return i + 1
		}
	}

	return -1
}
