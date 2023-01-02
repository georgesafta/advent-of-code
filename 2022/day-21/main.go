package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type operation struct {
	x, operand, y string
}

type monkey struct {
	op     operation
	caller string
}

var ops map[string]func(int, int) int = map[string]func(int, int) int{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"*": func(x, y int) int { return x * y },
	"/": func(x, y int) int { return x / y },
}

// a op x = y => a = y reverseOp x
var reverseOps map[string]func(int, int, bool) int = map[string]func(int, int, bool) int{
	"+": func(x, y int, isLeftOperand bool) int { return y - x },
	"-": func(x, y int, isLeftOperand bool) int {
		// x - a = y
		if isLeftOperand {
			return x - y
		}
		// a - x = y
		return y + x
	},
	"*": func(x, y int, isLeftOperand bool) int { return y / x },
	"/": func(x, y int, isLeftOperand bool) int {
		// x / a = y
		if isLeftOperand {
			return x / y
		}
		// a / x = y
		return y * x
	},
}

func main() {
	lines := util.ReadFile("./input/day-21.txt")
	numbers, operations := parseLines(lines)
	fmt.Println("evaluate = ", evaluate("root", numbers, operations))
	fmt.Println("part2 = ", part2(numbers, operations))
}

func part2(numbers map[string]int, operations map[string]operation) int {
	monkeys := monkeys(numbers, operations)
	stack := []string{}
	curr := monkeys["humn"].caller
	for curr != "root" {
		stack = append(stack, curr)
		curr = monkeys[curr].caller
	}

	total := 0
	if m := operations["root"]; m.x == stack[len(stack)-1] {
		total = numbers[m.y]
	} else {
		total = numbers[m.x]
	}

	for len(stack) > 0 {
		m := monkeys[stack[len(stack)-1]]
		op := "humn"
		if len := len(stack); len > 1 {
			op = stack[len-2]
		}
		if m.op.x != op {
			total = reverseOps[m.op.operand](numbers[m.op.x], total, true)
		} else {
			total = reverseOps[m.op.operand](numbers[m.op.y], total, false)
		}
		stack = stack[:len(stack)-1]
	}

	return total
}

func evaluate(name string, numbers map[string]int, operations map[string]operation) int {
	if v, exists := numbers[name]; exists {
		return v
	}

	operation := operations[name]
	x := evaluate(operation.x, numbers, operations)
	y := evaluate(operation.y, numbers, operations)
	result := ops[operation.operand](x, y)
	numbers[name] = result
	return result
}

func parseLines(lines []string) (map[string]int, map[string]operation) {
	numbers := map[string]int{}
	operations := map[string]operation{}
	for _, line := range lines {
		arr := strings.Fields(line)
		key := arr[0][:len(arr[0])-1]
		if len(arr) == 2 {
			numbers[key] = util.Atoi(arr[1])
		} else {
			x := arr[1]
			operand := arr[2]
			y := arr[3]
			operations[key] = operation{x, operand, y}
		}
	}

	return numbers, operations
}

func monkeys(numbers map[string]int, operations map[string]operation) map[string]monkey {
	monkeys := map[string]monkey{}
	for k, op := range operations {
		monkeys[k] = monkey{op, ""}
	}

	for k := range numbers {
		if _, exists := monkeys[k]; !exists {
			monkeys[k] = monkey{}
		}
	}

	for k, m := range monkeys {
		if m.op.x != "" {
			key := m.op.x
			monkey := monkeys[key]
			monkey.caller = k
			monkeys[key] = monkey
		}
		if m.op.y != "" {
			key := m.op.y
			monkey := monkeys[key]
			monkey.caller = k
			monkeys[key] = monkey
		}
	}

	return monkeys
}
