package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	m := createGrid(lines, []func(int) int{on, off, flip})
	fmt.Println("lights on =", countOn(m))
	m = createGrid(lines, []func(int) int{intensityOn, intensityOff, intensityFlip})
	fmt.Println("total =", total(m))
}

func countOn(m [][]int) int {
	count := 0
	for _, row := range m {
		for _, v := range row {
			if v == 1 {
				count++
			}
		}
	}

	return count
}

func total(m [][]int) int {
	total := 0
	for _, row := range m {
		for _, v := range row {
			total += v
		}
	}

	return total
}

func createGrid(lines []string, mutators []func(int) int) [][]int {
	m := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = make([]int, 1000)
	}

	for _, line := range lines {
		mutator := mutators[0]
		s := line
		if strings.Contains(line, "turn on ") {
			mutator = mutators[0]
			s = s[8:]
		} else if strings.Contains(line, "turn off ") {
			mutator = mutators[1]
			s = s[9:]
		} else {
			mutator = mutators[2]
			s = s[7:]
		}
		var xStart, xEnd, yStart, yEnd int
		fmt.Sscanf(s, "%d,%d through %d,%d", &xStart, &yStart, &xEnd, &yEnd)
		changeValues(m, xStart, yStart, xEnd, yEnd, mutator)
	}

	return m
}

func changeValues(m [][]int, sRow, sCol, eRow, eCol int, mutator func(int) int) {
	for i := sRow; i <= eRow; i++ {
		for j := sCol; j <= eCol; j++ {
			m[i][j] = mutator(m[i][j])
		}
	}
}

func on(original int) int {
	return 1
}

func off(original int) int {
	return 0
}

func flip(original int) int {
	if original == 1 {
		return 0
	}
	return 1
}

func intensityOn(original int) int {
	return original + 1
}

func intensityOff(original int) int {
	return util.MaxInt(0, original-1)
}

func intensityFlip(original int) int {
	return original + 2
}
