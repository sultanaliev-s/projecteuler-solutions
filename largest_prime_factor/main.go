package main

import (
	"fmt"
	"math"
)

func main() {
	target := int64(600_851_475_143)
	primes := []int64{2}
	maxFactor := int64(1)

	i := 0
	for target != 1 {
		if i == len(primes) {
			primes = nextPrime(primes)
		}
		if target%primes[i] == 0 {
			target /= primes[i]
			maxFactor = max(primes[i], maxFactor)
		} else {
			i++
		}
	}

	fmt.Println(maxFactor)
}

func nextPrime(primes []int64) []int64 {
	x := primes[len(primes)-1] + 1
	for {
		if isPrime(x) {
			return append(primes, x)
		}
		x++
	}

}

func isPrime(x int64) bool {
	limit := int64(math.Sqrt(float64(x)))
	for i := int64(2); i <= limit; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func max(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
