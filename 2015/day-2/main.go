package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Println("total paper =", totalPaper(lines))
	fmt.Println("total ribbon =", totalRibbon(lines))
}

func totalPaper(lines []string) int {
	total := 0
	for _, line := range lines {
		arr := strings.Split(line, "x")
		l, w, h := util.Atoi(arr[0]), util.Atoi(arr[1]), util.Atoi(arr[2])
		min := util.MinInt(h*l, util.MinInt(l*w, w*h))
		total += 2*l*w + 2*w*h + 2*h*l + min
	}

	return total
}

func totalRibbon(lines []string) int {
	total := 0
	for _, line := range lines {
		arr := strings.Split(line, "x")
		l, w, h := util.Atoi(arr[0]), util.Atoi(arr[1]), util.Atoi(arr[2])
		min1, min2 := l, w
		tmp := h
		if min1 > tmp {
			tmp, min1 = min1, tmp
		}
		if min2 > tmp {
			tmp, min2 = min2, tmp
		}
		total += 2*min1 + 2*min2 + l*w*h
	}

	return total
}
