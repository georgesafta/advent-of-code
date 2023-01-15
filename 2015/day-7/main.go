package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

type op struct {
	value    int16
	hasValue bool
	a, b     string
	action   string
}

func main() {
	lines := util.ReadFile("input.txt")
	m := parse(lines)
	part1 := dfs("a", m)
	fmt.Println("a =", part1)
	m = parse(lines)
	m["b"] = op{part1, true, "", "", "SET"}
	fmt.Println("a =", dfs("a", m))
}

func dfs(target string, m map[string]op) int16 {
	elem := m[target]
	if elem.hasValue {
		return elem.value
	}

	a := int16(0)
	b := int16(0)

	if elem.a[0] <= '9' && elem.a[0] >= '0' {
		a = int16(util.Atoi(elem.a))
	} else {
		a = dfs(elem.a, m)
	}
	if elem.b != "" {
		if elem.b[0] <= '9' && elem.b[0] >= '0' {
			b = int16(util.Atoi(elem.b))
		} else {
			b = dfs(elem.b, m)
		}
	}

	value := int16(0)
	if elem.action == "AND" {
		value = a & b
	} else if elem.action == "OR" {
		value = a | b
	} else if elem.action == "LSHIFT" {
		value = a << b
	} else if elem.action == "RSHIFT" {
		value = a >> b
	} else if elem.action == "NOT" {
		value = ^a
	} else {
		value = a
	}

	m[target] = op{value, true, elem.a, elem.b, elem.action}

	return value
}

func parse(lines []string) map[string]op {
	m := map[string]op{}
	for _, line := range lines {
		arr := strings.Fields(line)
		value := int16(0)
		hasValue := false
		if strings.Contains(line, "AND") {
			a := arr[0]
			b := arr[2]
			if a[0] <= '9' && a[0] >= '0' && b[0] <= '9' && b[0] >= '0' {
				hasValue = true
				value = int16(util.Atoi(a)) & int16(util.Atoi(b))
			}
			m[arr[4]] = op{value, hasValue, a, b, "AND"}
		} else if strings.Contains(line, "OR") {
			a := arr[0]
			b := arr[2]
			if a[0] <= '9' && a[0] >= '0' && b[0] <= '9' && b[0] >= '0' {
				hasValue = true
				value = int16(util.Atoi(a)) | int16(util.Atoi(b))
			}
			m[arr[4]] = op{value, hasValue, a, b, "OR"}
		} else if strings.Contains(line, "LSHIFT") {
			a := arr[0]
			b := arr[2]
			if a[0] <= '9' && a[0] >= '0' && b[0] <= '9' && b[0] >= '0' {
				hasValue = true
				value = int16(util.Atoi(a)) << int16(util.Atoi(b))
			}
			m[arr[4]] = op{value, hasValue, a, b, "LSHIFT"}
		} else if strings.Contains(line, "RSHIFT") {
			a := arr[0]
			b := arr[2]
			if a[0] <= '9' && a[0] >= '0' && b[0] <= '9' && b[0] >= '0' {
				hasValue = true
				value = int16(util.Atoi(a)) >> int16(util.Atoi(b))
			}
			m[arr[4]] = op{value, hasValue, a, b, "RSHIFT"}
		} else if strings.Contains(line, "NOT") {
			a := arr[1]
			if a[0] <= '9' && a[0] >= '0' {
				hasValue = true
				value = ^int16(util.Atoi(a))
			}
			m[arr[3]] = op{value, hasValue, a, "", "NOT"}
		} else {
			a := arr[0]
			if a[0] <= '9' && a[0] >= '0' {
				hasValue = true
				value = int16(util.Atoi(a))
			}
			m[arr[2]] = op{value, hasValue, a, "", "SET"}
		}
	}

	return m
}
