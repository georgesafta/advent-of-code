package main

import (
	"fmt"
)

func main() {
	arr := encode([]int{1, 1, 1, 3, 2, 2, 2, 1, 1, 3}, 40)
	fmt.Println("size =", len(arr))
	arr = encode(arr, 10)
	fmt.Println("size =", len(arr))
}

func encode(arr []int, times int) []int {
	for i := 0; i < times; i++ {
		arr = runLenghtEncode(arr)
	}

	return arr
}

func runLenghtEncode(arr []int) []int {
	encode := []int{}
	for i := 0; i < len(arr); {
		size := 0
		value := arr[i]
		for i < len(arr) && arr[i] == value {
			size++
			i++
		}
		encode = append(encode, size, value)
	}

	return encode
}
