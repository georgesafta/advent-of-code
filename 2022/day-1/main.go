package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input/day-1.txt")

	if err != nil {
		log.Fatal("could not open file", err)
	}

	defer file.Close()
	part2(file)
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)

	sum := 0
	max := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			v, _ := strconv.Atoi(text)
			sum += v
		}
	}

	fmt.Println("max : ", max)
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)

	sum := 0
	max1 := 0
	max2 := 0
	max3 := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			if sum > max1 {
				max2, max3 = max1, max2
				max1 = sum
			} else if sum > max2 {
				max3 = max2
				max2 = sum
			} else if sum > max3 {
				max3 = sum
			}
			sum = 0
		} else {
			v, _ := strconv.Atoi(text)
			sum += v
		}
	}

	fmt.Println("max : ", max1+max2+max3)
}
