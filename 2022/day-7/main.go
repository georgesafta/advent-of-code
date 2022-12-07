package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dir struct {
	Files   map[string]bool
	Folders map[string]*dir
	Size    int
	Parent  *dir
}

func (r *dir) totalSmallerThan(min int) int {
	if r == nil {
		return 0
	}
	total := 0
	if r.Size <= min {
		total += r.Size
	}
	for _, d := range r.Folders {
		total += d.totalSmallerThan(min)
	}

	return total
}

func (r *dir) largerThan(max int) []int {
	if r == nil {
		return []int{}
	}

	larger := []int{}
	if r.Size >= max {
		larger = append(larger, r.Size)
	}
	for _, d := range r.Folders {
		larger = append(larger, d.largerThan(max)...)
	}

	return larger
}

func main() {
	file, err := os.Open("./input/day-7.txt")

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()
	part2(file)
}

func part1(file *os.File) {
	root := parseFileTree(file)

	fmt.Println("total = ", root.totalSmallerThan(100000))
}

func bubbleUp(d *dir, size int) {
	for curr := d.Parent; curr != nil; {
		curr.Size += size
		curr = curr.Parent
	}
}

func part2(file *os.File) {
	root := parseFileTree(file)
	unused := 70000000 - root.Size
	if unused >= 30000000 {
		fmt.Println("min = 0")
	} else {
		max := 30000000 - unused
		larger := root.largerThan(max)
		fmt.Println("min = ", min(larger))
	}
}

func parseFileTree(file *os.File) *dir {
	scanner := bufio.NewScanner(file)
	root := &dir{Files: map[string]bool{}, Folders: map[string]*dir{}, Size: 0, Parent: nil}

	curr := root
	for scanner.Scan() {
		text := scanner.Text()
		arr := strings.Split(text, " ")
		if "$" == arr[0] {
			if "cd" != arr[1] {
				continue
			}
			if ".." == arr[2] {
				curr = curr.Parent
			} else if "/" == arr[2] {
				curr = root
			} else {
				d, exists := curr.Folders[arr[2]]
				if exists {
					curr = d
				} else {
					d := &dir{Files: map[string]bool{}, Folders: map[string]*dir{}, Size: 0, Parent: curr}
					curr.Folders[arr[2]] = d
					curr = d
				}
			}
		} else {
			if "dir" == arr[0] {
				continue
			}
			size, _ := strconv.Atoi(arr[0])
			_, exists := curr.Files[arr[1]]
			if !exists {
				curr.Files[arr[1]] = true
				curr.Size += size
				bubbleUp(curr, size)
			}
		}
	}

	return root
}

func min(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	min := arr[0]
	for _, e := range arr {
		if e < min {
			min = e
		}
	}

	return min
}
