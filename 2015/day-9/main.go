package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	graph, places := generateGraph(lines)

	min := []int{10000000}
	max := []int{-1}
	backtrack("", graph, places, map[string]bool{}, min, max, 0)
	fmt.Println("min =", min[0], "max =", max[0])
}

func generateGraph(lines []string) (map[string]int, map[string]bool) {
	m := map[string]int{}
	places := map[string]bool{}
	for _, line := range lines {
		arr := strings.Fields(line)
		src := arr[0]
		dest := arr[2]
		dist := util.Atoi(arr[4])

		m[key(src, dest)] = dist
		m[key(dest, src)] = dist
		places[src] = true
		places[dest] = true
	}

	return m, places
}

func backtrack(src string, graph map[string]int, places, visited map[string]bool, min, max []int, cost int) {
	if len(places) == len(visited) {
		min[0] = util.MinInt(min[0], cost)
		max[0] = util.MaxInt(max[0], cost)
		return
	}

	for place := range places {
		if _, exists := visited[place]; exists || place == src {
			continue
		}
		visited[place] = true
		if src != "" {
			k := key(src, place)
			backtrack(place, graph, places, visited, min, max, cost+graph[k])
		} else {
			backtrack(place, graph, places, visited, min, max, cost)
		}
		delete(visited, place)
	}
}

func key(src, dest string) string {
	return fmt.Sprintf("%s,%s", src, dest)
}
