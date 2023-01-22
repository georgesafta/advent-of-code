package main

import (
	"fmt"
	"strings"
)

func main() {
	s := toIntArr("hxbxwxba")
	next := findNextSuitable(s)
	fmt.Println("next =", next)
	s = toIntArr(next)
	fmt.Println("next =", findNextSuitable(s))
}

func findNextSuitable(s []int) string {
	s = next(s)
	for !has3Consecutive(s) || !hasIdenticalPairs(s, 2) || containsIllegal(s) {
		s = next(s)
	}

	return toString(s)
}

func toString(arr []int) string {
	s := []string{}
	for _, v := range arr {
		r := rune(v + 'a')
		s = append(s, string(r))
	}

	return strings.Join(s, "")
}

func toIntArr(s string) []int {
	arr := []int{}
	for _, r := range s {
		arr = append(arr, int(r-'a'))
	}

	return arr
}

func next(input []int) []int {
	next := []int{}
	next = append(next, input...)
	mod := 26
	add := 1
	for i := len(next) - 1; i > 0; i-- {
		curr := next[i] + add
		add = curr / mod
		next[i] = curr % mod
	}
	next[0] = (next[0] + add) % mod

	return next
}

func has3Consecutive(arr []int) bool {
	for i := 0; i+2 < len(arr); i++ {
		if arr[i]+1 == arr[i+1] && arr[i+1]+1 == arr[i+2] {
			return true
		}
	}

	return false
}

func hasIdenticalPairs(arr []int, times int) bool {
	for i := 0; i+1 < len(arr); i++ {
		if arr[i] == arr[i+1] {
			times--
			i++
		}
		if times == 0 {
			return true
		}
	}

	return false
}

func containsIllegal(arr []int) bool {
	m := map[int]bool{'i' - 'a': true, 'o' - 'a': true, 'l' - 'a': true}
	for _, v := range arr {
		if _, exists := m[v]; exists {
			return true
		}
	}

	return false
}
