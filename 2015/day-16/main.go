package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	target := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
	aunts := parse(lines)
	fmt.Println("Sue =", find(target, aunts, equals))
	fmt.Println("Sue =", find(target, aunts, intervalMatch))
}

func parse(lines []string) map[string]map[string]int {
	aunts := map[string]map[string]int{}
	for _, line := range lines {
		arr := strings.Fields(line)
		m := map[string]int{}
		for i := 2; i < len(arr); i += 2 {
			key := arr[i][:len(arr[i])-1]
			value := arr[i+1]
			if i+1 != len(arr)-1 {
				value = value[:len(value)-1]
			}
			m[key] = util.Atoi(value)
		}
		key := arr[1][:len(arr[1])-1]
		aunts[key] = m
	}

	return aunts
}

func find(target map[string]int, aunts map[string]map[string]int, isOk func(string, int, int) bool) string {
	for k, v := range aunts {
		if isMatch(target, v, isOk) {
			return k
		}
	}

	return ""
}

func isMatch(target, aunt map[string]int, isOk func(string, int, int) bool) bool {
	for k, v := range target {
		if value, exists := aunt[k]; exists && !isOk(k, v, value) {
			return false
		}
	}

	return true
}

func equals(key string, expected, actual int) bool {
	return expected == actual
}

func intervalMatch(key string, expected, actual int) bool {
	if key == "cats" || key == "trees" {
		return actual > expected
	}

	if key == "pomeranians" || key == "goldfish" {
		return actual < expected
	}

	return expected == actual
}
