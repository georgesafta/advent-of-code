package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	m := createGrid(lines)
	fmt.Println("lights on =", countOn(m))
	grid := createIntensityGrid(lines)
	fmt.Println("total =", total(grid))
}

func countOn(m [][]bool) int {
	count := 0
	for _, row := range m {
		for _, v := range row {
			if v {
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

func createGrid(lines []string) [][]bool {
	m := make([][]bool, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = make([]bool, 1000)
	}

	for _, line := range lines {
		mutator := on
		s := line
		if strings.Contains(line, "turn on ") {
			mutator = on
			s = s[8:]
		} else if strings.Contains(line, "turn off ") {
			mutator = off
			s = s[9:]
		} else {
			mutator = flip
			s = s[7:]
		}
		var xStart, xEnd, yStart, yEnd int
		fmt.Sscanf(s, "%d,%d through %d,%d", &xStart, &yStart, &xEnd, &yEnd)
		changeValues(m, xStart, yStart, xEnd, yEnd, mutator)
	}

	return m
}

func createIntensityGrid(lines []string) [][]int {
	m := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		m[i] = make([]int, 1000)
	}

	for _, line := range lines {
		mutator := intensityOn
		s := line
		if strings.Contains(line, "turn on ") {
			mutator = intensityOn
			s = s[8:]
		} else if strings.Contains(line, "turn off ") {
			mutator = intensityOff
			s = s[9:]
		} else {
			mutator = insensityFlip
			s = s[7:]
		}
		var xStart, xEnd, yStart, yEnd int
		fmt.Sscanf(s, "%d,%d through %d,%d", &xStart, &yStart, &xEnd, &yEnd)
		changeIntensity(m, xStart, yStart, xEnd, yEnd, mutator)
	}

	return m
}

func changeValues(m [][]bool, sRow, sCol, eRow, eCol int, mutator func(bool) bool) {
	for i := sRow; i <= eRow; i++ {
		for j := sCol; j <= eCol; j++ {
			m[i][j] = mutator(m[i][j])
		}
	}
}

func changeIntensity(m [][]int, sRow, sCol, eRow, eCol int, mutator func(int) int) {
	for i := sRow; i <= eRow; i++ {
		for j := sCol; j <= eCol; j++ {
			m[i][j] = mutator(m[i][j])
		}
	}
}

func on(original bool) bool {
	return true
}

func off(original bool) bool {
	return false
}

func flip(original bool) bool {
	return !original
}

func intensityOn(original int) int {
	return original + 1
}

func intensityOff(original int) int {
	return util.MaxInt(0, original-1)
}

func insensityFlip(original int) int {
	return original + 2
}
