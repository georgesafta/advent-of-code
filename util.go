package util

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) []string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func MinInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Atoi(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		log.Panic()
	}

	return value
}
