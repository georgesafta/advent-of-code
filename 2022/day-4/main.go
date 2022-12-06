package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	Start int
	End   int
}

func (i interval) contains(ii interval) bool {
	return i.Start <= ii.Start && i.End >= ii.End
}
func (i interval) overlaps(ii interval) bool {
	return (i.Start <= ii.Start && i.End >= ii.Start) || (i.Start <= ii.End && i.End >= ii.End)
}

func main() {
	file, err := os.Open("./input/day-4.txt")

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()
	part2(file)
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	overlap := 0
	for scanner.Scan() {
		text := scanner.Text()
		i1, i2 := toIntevals(text)
		if i1.contains(i2) || i2.contains(i1) {
			overlap++
		}
	}

	fmt.Println("overlap : ", overlap)
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)

	overlap := 0
	for scanner.Scan() {
		text := scanner.Text()
		i1, i2 := toIntevals(text)
		if i1.overlaps(i2) || i2.overlaps(i1) {
			overlap++
		}
	}

	fmt.Println("overlap : ", overlap)
}

func toIntevals(text string) (interval, interval) {
	arr := strings.Split(text, ",")
	return toInteval(arr[0]), toInteval(arr[1])
}

func toInteval(text string) interval {
	arr := strings.Split(text, "-")
	start, _ := strconv.Atoi(arr[0])
	end, _ := strconv.Atoi(arr[1])
	return interval{Start: start, End: end}
}
