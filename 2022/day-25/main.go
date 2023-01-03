package main

import (
	"fmt"

	util "github.com/georgesafta/advent-of-code"
)

type mapping struct {
	carry int
	val   byte
}

var mappings map[int]mapping = map[int]mapping{
	-5: {-1, '0'},
	-4: {-1, '1'},
	-3: {-1, '2'},
	-2: {0, '='},
	-1: {0, '-'},
	0:  {0, '0'},
	1:  {0, '1'},
	2:  {0, '2'},
	3:  {1, '='},
	4:  {1, '-'},
	5:  {1, '0'},
}

func main() {
	lines := util.ReadFile("./input/day-25.txt")
	fmt.Println("sum = ", sum(lines))
}

func sum(lines []string) string {
	sum := "0"
	for _, line := range lines {
		sum = add(sum, line)
	}

	return sum
}

func add(s, t string) string {
	r, v := byte('0'), byte('0')
	i, j := len(s)-1, len(t)-1
	carry := 0
	arr := []byte{}
	for i >= 0 || j >= 0 {
		if i >= 0 {
			r = s[i]
			i--
		} else {
			r = byte('0')
		}
		if j >= 0 {
			v = t[j]
			j--
		} else {
			v = byte('0')
		}
		b := byte('0')
		carry, b = addBytes(r, v, carry)
		arr = append([]byte{b}, arr...)
	}
	if carry == 1 {
		arr = append([]byte{byte('1')}, arr...)
	}

	return string(arr)
}

func addBytes(r, v byte, carry int) (int, byte) {
	x := toInt(r)
	y := toInt(v)
	mapping := mappings[x+y+carry]

	return mapping.carry, mapping.val
}

func toInt(r byte) int {
	if r == '=' {
		return -2
	}
	if r == '-' {
		return -1
	}

	return int(r - '0')
}
