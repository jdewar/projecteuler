package main

// Find diff between sum of squares and square of sum of 1..100
import (
	"fmt"
)

func main() {
	n := 100
	s := n * (n + 1) / 2
	fmt.Println(s*s - n*(n+1)*(2*n+1)/6)
}
