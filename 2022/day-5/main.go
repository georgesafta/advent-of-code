package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type stack []string

type move struct {
	Amount int
	From   int
	To     int
}

func (s stack) Push(e string) stack {
	return append(s, e)
}

func (s stack) IsEmpty() bool {
	return len(s) == 0
}

func (s stack) Pop() (stack, string) {
	if s.IsEmpty() {
		return s, ""
	}
	index := len(s) - 1

	return s[:index], s[index]
}

func (s stack) Add(e string) stack {
	s = append([]string{e}, s...)
	return s
}

func (s stack) Peek() string {
	if len(s) == 0 {
		return ""
	}
	return s[len(s)-1]
}

func main() {
	lines := util.ReadFile("./input/day-5.txt")
	m, index := parse(lines)
	moveCrates(lines, m, index, false)
	fmt.Println()

	m, index = parse(lines)
	moveCrates(lines, m, index, true)
}

func parse(lines []string) (map[int]stack, int) {
	m := map[int]stack{}
	index := 0
	for i, line := range lines {
		if line == "" {
			index = i + 1
			break
		}
		if !strings.Contains(line, "[") {
			continue
		}
		m = parseCrates(line, m)
	}

	return m, index
}

func moveCrates(lines []string, m map[int]stack, index int, withChunks bool) {
	parts := len(m)
	for i := index; i < len(lines); i++ {
		text := lines[i]
		move := parseMove(text)
		if withChunks {
			simulateChunk(move, m)
		} else {
			simulate(move, m)
		}
	}

	for i := 1; i <= parts; i++ {
		st, exists := m[i]
		if exists && !st.IsEmpty() {
			fmt.Print(st.Peek())
		}
	}
}

func parseCrates(text string, m map[int]stack) map[int]stack {
	for i, e := range text {
		if e != ' ' && e != '[' && e != ']' {
			st, exists := m[i/4+1]
			if exists {
				st = st.Add(string(e))
				m[i/4+1] = st
			} else {
				st = stack{}
				st = st.Add(string(e))
				m[i/4+1] = st
			}
		}
	}
	return m
}

func parseMove(text string) move {
	arr := strings.Split(text, " ")
	amount := util.Atoi(arr[1])
	from := util.Atoi(arr[3])
	to := util.Atoi(arr[5])
	return move{Amount: amount, From: from, To: to}
}

func simulate(move move, m map[int]stack) {
	st, exists := m[move.To]
	if !exists {
		st = make(stack, 0)
	}
	for i := 0; i < move.Amount; i++ {
		s, e := m[move.From].Pop()
		st = st.Push(e)
		m[move.To] = st
		m[move.From] = s
	}
}

func simulateChunk(move move, m map[int]stack) {
	st, exists := m[move.To]
	if !exists {
		st = make(stack, 0)
	}
	elems := []string{}
	for i := 0; i < move.Amount; i++ {
		s, e := m[move.From].Pop()
		elems = append(elems, e)
		m[move.From] = s
	}

	for i := len(elems) - 1; i >= 0; i-- {
		st = st.Push(elems[i])
	}
	m[move.To] = st
}
