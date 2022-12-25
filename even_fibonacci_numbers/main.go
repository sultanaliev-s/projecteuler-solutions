package main

import "fmt"

func main() {
	sum := 0
	a := 0
	b := 1
	t := b
	i := 0
	for a <= 4_000_000 {
		if i == 0 {
			i = 3
			sum += a
		}
		t = b
		b = a + b
		a = t
		i--
	}

	fmt.Println(sum)
}
