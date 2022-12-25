package main

import "fmt"

func main() {
	var sum int64
	var sieve [2_000_000]bool
	sieve[0] = true
	sieve[1] = true
	for i := range sieve {
		if !sieve[i] {
			sum += int64(i)
			for j := i + i; j < len(sieve); j += i {
				sieve[j] = true
			}
		}
	}

	fmt.Println(sum)
}
