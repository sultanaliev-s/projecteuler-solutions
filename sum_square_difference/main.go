package main

import "fmt"

func main() {
	sumSquares := int64(0)
	for i := int64(1); i <= 100; i++ {
		sumSquares += i * i
	}

	sumHundred := int64((100 * 101) / 2)
	squareSum := sumHundred * sumHundred

	diff := squareSum - sumSquares
	fmt.Println(diff)
}
