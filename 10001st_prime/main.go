package main

import (
	"fmt"
	"math"
)

func main() {
	prime := 2
	for i := 1; i < 10_001; i++ {
		prime = nextPrime(prime)
	}

	fmt.Println(prime)
}

func nextPrime(x int) int {
	for x += 1; ; x++ {
		if isPrime(x) {
			return x
		}
	}

}

func isPrime(x int) bool {
	limit := int(math.Sqrt(float64(x)))
	for i := 2; i <= limit; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
