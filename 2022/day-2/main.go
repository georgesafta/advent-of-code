package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type move struct {
	Win   byte
	Loss  byte
	Value int
}

func (m move) isWin(b byte) bool {
	return m.Win == b
}
func (m move) isLoss(b byte) bool {
	return m.Loss == b
}
func (m move) score(b byte) int {
	score := 3
	if m.isWin(b) {
		score = 6
	}
	if m.isLoss(b) {
		score = 0
	}

	return m.Value + score
}

func main() {
	m := make(map[byte]move)
	m['X'] = move{
		Win:   'C',
		Loss:  'B',
		Value: 1,
	}
	m['Y'] = move{
		Win:   'A',
		Loss:  'C',
		Value: 2,
	}
	m['Z'] = move{
		Win:   'B',
		Loss:  'A',
		Value: 3,
	}
	file, err := os.Open("./input/day-2.txt")

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()
	part2(file, m)
}

func part1(file *os.File, m map[byte]move) {
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		score += m[text[2]].score(text[0])
	}

	fmt.Println("score : ", score)
}

func part2(file *os.File, m map[byte]move) {
	// there are easier ways to to this, but just wanted to reuse the first part
	mapper := make(map[byte]map[byte]byte)
	mapper['X'] = winMapping()
	mapper['Y'] = drawMapping()
	mapper['Z'] = lossMapping()
	scanner := bufio.NewScanner(file)

	score := 0
	for scanner.Scan() {
		text := scanner.Text()
		b := mapper[text[2]][text[0]]
		score += m[b].score(text[0])
	}

	fmt.Println("score : ", score)
}

func lossMapping() map[byte]byte {
	loss := make(map[byte]byte)
	loss['A'] = 'Y'
	loss['B'] = 'Z'
	loss['C'] = 'X'
	return loss
}

func drawMapping() map[byte]byte {
	draw := make(map[byte]byte)
	draw['A'] = 'X'
	draw['B'] = 'Y'
	draw['C'] = 'Z'
	return draw
}

func winMapping() map[byte]byte {
	win := make(map[byte]byte)
	win['A'] = 'Z'
	win['B'] = 'X'
	win['C'] = 'Y'
	return win
}
