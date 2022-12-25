package main

import (
	"fmt"
	"strconv"
)

func main() {
	maxPalindrome := 0
	x := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			x = i * j
			if isPalindrome(x) {
				maxPalindrome = max(maxPalindrome, x)
			}
		}
	}

	fmt.Println(maxPalindrome)
}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
