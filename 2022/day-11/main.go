package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type monkey struct {
	items     []int
	divideBy  int
	op        func(int) int
	goodPass  int
	badPass   int
	processed int
}

func (m *monkey) String() string {
	return fmt.Sprintf("items: %d, divide: %d, true: %d, false: %d, processed: %d", m.items, m.divideBy, m.goodPass, m.badPass, m.processed)
}

func main() {
	input := util.ReadFile("./input/day-11.txt")
	m := parseInput(input)
	lcm := 1
	for _, m := range m {
		lcm *= m.divideBy
	}

	//solve(m, 20, func(n int) int { return n / 3 })
	solve(m, 10000, func(n int) int { return n % lcm })
}

func solve(monkeys []*monkey, rounds int, reduce func(int) int) {
	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			processMonkeyItems(m, monkeys, reduce)
		}
	}

	processed := []int{}
	for _, m := range monkeys {
		processed = append(processed, m.processed)
	}
	fmt.Println(processed)

	sort.Sort(sort.Reverse(sort.IntSlice(processed)))
	fmt.Println("total = ", (processed[0] * processed[1]))
}

func processMonkeyItems(m *monkey, monkeys []*monkey, reduce func(int) int) {
	m.processed += len(m.items)
	for _, v := range m.items {
		worry := reduce(m.op(v))
		if worry%m.divideBy == 0 {
			monkey := monkeys[m.goodPass]
			monkey.items = append(monkey.items, worry)
		} else {
			monkey := monkeys[m.badPass]
			monkey.items = append(monkey.items, worry)
		}
	}
	m.items = []int{}
}

func parseInput(lines []string) []*monkey {
	m := []*monkey{}
	for i := 0; i < len(lines); i += 7 {
		items := toItems(lines[i+1])
		op := parseOperation(lines[i+2])
		divideBy, _ := strconv.Atoi(strings.Fields(lines[i+3])[3])
		goodPass, _ := strconv.Atoi(strings.Fields(lines[i+4])[5])
		badPass, _ := strconv.Atoi(strings.Fields(lines[i+5])[5])
		monkey := &monkey{items: items, divideBy: divideBy, op: op, goodPass: goodPass, badPass: badPass, processed: 0}
		fmt.Println(monkey)
		m = append(m, monkey)
	}

	return m
}

func toItems(line string) []int {
	arr := strings.Fields(line)[2:]
	items := []int{}
	for _, v := range arr {
		if v[len(v)-1] == ',' {
			i, _ := strconv.Atoi(v[:len(v)-1])
			items = append(items, i)
		} else {
			i, _ := strconv.Atoi(v)
			items = append(items, i)
		}
	}

	return items
}

func parseOperation(line string) func(int) int {
	arr := strings.Fields(line)
	f := func(old int) int {
		x := old
		y := old
		if arr[3] != "old" {
			x, _ = strconv.Atoi(arr[3])
		}
		if arr[5] != "old" {
			y, _ = strconv.Atoi(arr[5])
		}

		if arr[4] == "*" {
			return x * y
		}
		return x + y
	}

	return f
}
