package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

type listNode struct {
	value int
	prev  *listNode
	next  *listNode
}

func (node *listNode) shift(steps int) {
	if steps == 0 {
		return
	}

	currNext := node.next
	currPrev := node.prev
	currPrev.next = currNext
	currNext.prev = currPrev

	if steps < 0 {
		curr := currPrev
		for ; steps < -1; steps++ {
			curr = curr.prev
		}
		prev := curr.prev
		prev.next = node
		curr.prev = node
		node.next = curr
		node.prev = prev
	} else {
		curr := currNext
		for ; steps > 1; steps-- {
			curr = curr.next
		}
		next := curr.next
		next.prev = node
		curr.next = node
		node.next = next
		node.prev = curr
	}
}

func main() {
	lines := util.ReadFile("./input/day-20.txt")

	nodes, zero := parseLines(lines)
	fmt.Println("part1 = ", rotate(nodes, zero, 1, 1))

	nodes, zero = parseLines(lines)
	fmt.Println("part2 = ", rotate(nodes, zero, 10, 811589153))
}

func rotate(nodes []*listNode, zero *listNode, repetitions, multiply int) int {
	for ; repetitions > 0; repetitions-- {
		for _, node := range nodes {
			mod := len(nodes) - 1
			shift := ((node.value % mod) * (multiply % mod)) % mod
			node.shift(shift)
		}
	}

	sum := 0
	indexes := []int{1000, 2000, 3000}
	for _, index := range indexes {
		curr := zero
		for i := 0; i < index%len(nodes); i++ {
			curr = curr.next
		}
		sum += curr.value * multiply
	}

	return sum
}

func parseLines(lines []string) ([]*listNode, *listNode) {
	nodes := []*listNode{}
	var head *listNode
	var tail *listNode
	var zero *listNode

	for _, line := range lines {
		value := util.Atoi(line)
		node := &listNode{value, nil, nil}
		if head == nil {
			head = node
			node.next = node
		} else {
			node.next = head
		}
		if tail == nil {
			tail = node
			node.prev = node
		} else {
			node.prev = tail
			tail.next = node
			tail = node
		}
		head.prev = tail
		nodes = append(nodes, node)
		if node.value == 0 {
			zero = node
		}
	}

	return nodes, zero
}
