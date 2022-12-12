package util

import (
	"bufio"
	"log"
	"os"
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
