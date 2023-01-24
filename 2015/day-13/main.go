package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	m := hapinessChange(lines)
	var max *int = new(int)
	*max = -10000000
	backtrack(m, map[string]bool{}, []string{}, max)
	fmt.Println(*max)
	*max = -10000000
	addMe(m)
	backtrack(m, map[string]bool{}, []string{}, max)
	fmt.Println(*max)
}

func backtrack(hapiness map[string]map[string]int, visited map[string]bool, arr []string, max *int) {
	if len(hapiness) == len(visited) {
		curr := computeHapiness(hapiness, arr)
		if *max < curr {
			*max = curr
		}
		return
	}

	for k := range hapiness {
		if _, exists := visited[k]; !exists {
			visited[k] = true
			arr = append(arr, k)
			backtrack(hapiness, visited, arr, max)
			arr = arr[0 : len(arr)-1]
			delete(visited, k)
		}
	}
}

func computeHapiness(hapiness map[string]map[string]int, arr []string) int {
	total := 0
	len := len(arr)
	for i, v := range arr {
		next := i + 1
		prev := i - 1
		if next == len {
			next = 0
		}
		if prev == -1 {
			prev = len - 1
		}
		total += hapiness[v][arr[prev]]
		total += hapiness[v][arr[next]]
	}

	return total
}

func hapinessChange(lines []string) map[string]map[string]int {
	m := map[string]map[string]int{}
	for _, line := range lines {
		arr := strings.Fields(line)
		hapiness := util.Atoi(arr[3])
		if arr[2] == "lose" {
			hapiness = -hapiness
		}
		a := arr[0]
		b := arr[10][:len(arr[10])-1]
		if _, exists := m[a]; !exists {
			m[a] = map[string]int{}
		}
		m[a][b] = hapiness
	}

	return m
}

func addMe(happiness map[string]map[string]int) {
	me := "Me"
	happiness[me] = map[string]int{}
	for k, m := range happiness {
		happiness[me][k] = 0
		m[me] = 0
	}
}
