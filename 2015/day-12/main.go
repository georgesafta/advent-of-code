package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	total, _ := sum(lines[0], 0, func(s string) bool { return false })
	fmt.Println(total)
	total, _ = sum(lines[0], 0, func(s string) bool { return strings.Contains(s, "\"red\":") || strings.Contains(s, ":\"red\"") })
	fmt.Println(total)
}

func sum(s string, index int, isIgnored func(string) bool) (int, int) {
	total := 0
	isNegative := false
	curr := 0
	t := []byte{}
	for ; index < len(s) && s[index] != '}'; index++ {
		b := s[index]
		if b <= '9' && b >= '0' {
			curr *= 10
			curr += int(b - '0')
			t = append(t, b)
		} else {
			if isNegative {
				curr *= -1
			}
			total += curr
			curr = 0
			isNegative = false

			if b == '{' {
				subSum, end := sum(s, index+1, isIgnored)
				total += subSum
				index = end
			} else {
				if b == '-' {
					isNegative = true
				}
				t = append(t, b)
			}
		}
	}

	if isIgnored(string(t)) {
		return 0, index
	}

	if isNegative {
		curr *= -1
	}

	return total + curr, index
}
