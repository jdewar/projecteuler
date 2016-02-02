package main

import (
	"fmt"
	//	"math"
)

func main() {
	for a := 1; a < 333; a++ {
		for b := a + 1; b < 500; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				fmt.Println(a, b, c)
				break
			}
		}
	}
}
