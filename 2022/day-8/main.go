package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	m := fileToMatrix()
	part1(m)
	part2(m)
}

func part1(m [][]int) {
	rows := len(m)
	cols := len(m[0])
	left := buildMatrix(rows, cols)
	right := buildMatrix(rows, cols)
	up := buildMatrix(rows, cols)
	down := buildMatrix(rows, cols)

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			left[i][j] = max(m[i][j-1], left[i][j-1])
			up[i][j] = max(m[i-1][j], up[i-1][j])
			right[i][cols-1-j] = max(m[i][cols-j], right[i][cols-j])
			down[rows-1-i][cols-1-j] = max(m[rows-i][cols-1-j], down[rows-i][cols-1-j])
		}
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if right[i][j] < m[i][j] || left[i][j] < m[i][j] || up[i][j] < m[i][j] || down[i][j] < m[i][j] {
				count++
			}
		}
	}
	fmt.Println("seen trees = ", count)
}

func part2(m [][]int) {
	maximum := 0
	rows := len(m)
	cols := len(m[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			r := right(m, i, j)
			l := left(m, i, j)
			u := up(m, i, j)
			d := down(m, i, j)
			p := r * l * u * d
			maximum = max(p, maximum)
		}
	}

	fmt.Println("max = ", maximum)
}

func right(m [][]int, row, col int) int {
	count := 0
	for j := col + 1; j < len(m[0]); j++ {
		count++
		if m[row][j] >= m[row][col] {
			break
		}
	}

	return count
}

func left(m [][]int, row, col int) int {
	count := 0
	for j := col - 1; j >= 0; j-- {
		count++
		if m[row][j] >= m[row][col] {
			break
		}
	}

	return count
}

func up(m [][]int, row, col int) int {
	count := 0
	for i := row - 1; i >= 0; i-- {
		count++
		if m[i][col] >= m[row][col] {
			break
		}
	}

	return count
}

func down(m [][]int, row, col int) int {
	count := 0
	for i := row + 1; i < len(m); i++ {
		count++
		if m[i][col] >= m[row][col] {
			break
		}
	}

	return count
}

func buildMatrix(rows, cols int) [][]int {
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		m[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			m[i][j] = -1
		}
	}

	return m
}

func fileToMatrix() [][]int {
	file, err := os.Open("./input/day-8.txt")

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	rows := len(input)
	m := make([][]int, rows)
	for i := 0; i < rows; i++ {
		for _, e := range input[i] {
			m[i] = append(m[i], int(e-'0'))
		}
	}

	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
