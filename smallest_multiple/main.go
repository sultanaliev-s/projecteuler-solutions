package main

import "fmt"

func main() {
	for i := 20; ; i++ {
		isDivisible := true
		for j := 11; j <= 20; j++ {
			if i%j != 0 {
				isDivisible = false
				break
			}
		}
		if isDivisible {
			fmt.Println(i)
			break
		}
	}
}
