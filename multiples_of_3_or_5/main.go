package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 1000; i += 3 {
		sum += i
	}

	for i, j := 0, 0; i < 1000; i, j = i+5, j+1 {
		if j == 3 {
			j = 0
			continue
		}
		sum += i
	}

	fmt.Println(sum)
}
