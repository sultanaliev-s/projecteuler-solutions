package main

import (
	"fmt"
	"math"
)

func main() {
	found := false
	for b := 1; b < 1000 && !found; b++ {
		for a := 1; a < 1000 && !found; a++ {
			cSq := a*a + b*b
			c := int(math.Pow(float64(cSq), 0.5))
			if c*c == cSq {
				if a+b+c == 1000 {
					fmt.Println(a * b * c)
					found = true
				}
			}
		}
	}
}
