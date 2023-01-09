package main

import (
	"fmt"
	"strings"

	util "github.com/georgesafta/advent-of-code"
)

func main() {
	lines := util.ReadFile("input.txt")
	fmt.Println("nice strings =", niceStrings(lines, isNice))
	fmt.Println("nice strings =", niceStrings(lines, isNicer))
}

func niceStrings(lines []string, isNice func(string) bool) int {
	count := 0
	for _, line := range lines {
		if isNice(line) {
			count++
		}
	}

	return count
}

func isNice(s string) bool {
	return !hasDisallowedSequence(s) && hasAtLeast3Vowels(s) && hasDoubleConsecutiveLetter(s)
}

func isNicer(s string) bool {
	return hasNonOverlappingDuplicates(s) && hasDuplicateWithSpace(s)
}

func hasNonOverlappingDuplicates(s string) bool {
	m := map[string]int{}
	for i := 1; i < len(s); i++ {
		seq := s[i-1 : i+1]
		index, exists := m[seq]
		if exists && i-index-1 > 1 {
			return true
		}
		if !exists {
			m[seq] = i - 1
		}
	}

	return false
}

func hasDuplicateWithSpace(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}

	return false
}

func hasDisallowedSequence(s string) bool {
	disallowed := []string{"ab", "cd", "pq", "xy"}
	for _, seq := range disallowed {
		if strings.Contains(s, seq) {
			return true
		}
	}

	return false
}

func hasAtLeast3Vowels(s string) bool {
	count := 0
	for _, r := range s {
		if isVowel(r) {
			count++
		}
		if count >= 3 {
			return true
		}
	}

	return false
}

func hasDoubleConsecutiveLetter(s string) bool {
	prev := ' '
	for _, r := range s {
		if r == prev {
			return true
		}
		prev = r
	}

	return false
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
