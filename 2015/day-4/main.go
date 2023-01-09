package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var input string = "iwrupvqb"

func main() {
	fmt.Println("result =", findNumber("00000"))
	fmt.Println("result =", findNumber("000000"))
}

func findNumber(prefix string) int {
	for i := 0; ; i++ {
		s := fmt.Sprintf("%s%d", input, i)
		checksum := fmt.Sprintf("%x", md5.Sum([]byte(s)))
		if strings.HasPrefix(checksum, prefix) {
			return i
		}
	}
}
