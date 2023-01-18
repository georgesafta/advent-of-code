package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part1, part2 := total()
	fmt.Println("sum =", part1, part2)
}

func total() (int, int) {
	total := 0
	escapedTotal := 0
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		total += lineSum(line)
		escapedTotal += escapeLen(line) - len(line)
	}

	return total, escapedTotal
}

func lineSum(line string) int {
	sum := 0
	length := len(line)
	for i := 0; i < length; {
		r := line[i]
		i++
		if r == '"' {
			continue
		}
		if r == '\\' && i < length {
			c := line[i]
			if c == '\\' || c == '"' {
				i++
			} else if c == 'x' {
				i += 3
			}
		}
		sum++
	}

	return length - sum
}

func escapeLen(s string) int {
	length := len(s) + 2
	for _, r := range s {
		if r == '"' || r == '\'' || r == '\\' {
			length++
		}
	}

	return length
}
