package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input/day-3.txt")

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()
	part2(file)
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		score += value(commonLetter(text))
	}

	fmt.Println("score : ", score)
}

func part2(file *os.File) {
	m := map[rune]int{}
	scanner := bufio.NewScanner(file)

	score := 0
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		for k := range letters(text) {
			v, ok := m[k]
			if ok {
				m[k] = v + 1
			} else {
				m[k] = 1
			}
		}
		if i%3 == 2 {
			score += value(common(m))
			m = map[rune]int{}
		}
		i++
	}

	fmt.Println("score : ", score)
}

func commonLetter(s string) rune {
	len := len(s)
	m := map[rune]bool{}
	for i, v := range s {
		if i < len/2 {
			m[v] = true
		} else if _, ok := m[v]; ok {
			return v
		}
	}
	return '0'
}

func value(r rune) int {
	if r <= 'Z' {
		return 27 + int(r) - int('A')
	}
	return int(r) - int('a') + 1
}

func letters(s string) map[rune]bool {
	m := map[rune]bool{}
	for _, v := range s {
		m[v] = true
	}
	return m
}

func common(m map[rune]int) rune {
	for k, v := range m {
		if v == 3 {
			return k
		}
	}
	return '0'
}
