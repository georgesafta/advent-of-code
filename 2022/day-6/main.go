package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input/day-6.txt")

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()
	part2(file)
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text, " : ", marker(text, 4))
	}
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text, " : ", marker(text, 14))
	}
}

func marker(text string, size int) int {
	m := make(map[byte]int)
	for i, j := 0, 0; j < len(text); j++ {
		if j > size-1 {
			e := text[i]
			count, _ := m[e]
			if count == 1 {
				delete(m, e)
			} else {
				m[e] = count - 1
			}
			i++
		}

		e := text[j]
		count, exists := m[e]
		if exists {
			m[e] = count + 1
		} else {
			m[e] = 1
		}
		if len(m) == size {
			return j + 1
		}
	}

	return -1
}
