package main

// Find largest number that is a multiple of 2 three-digit numbers that is
// a palindrome
// HACK!  I only run to 800 because it's 'certainly' got to be above that

import (
	"fmt"
)

var start = 999

func is_palindrome(n int) bool {
	a := make([]int, 0)
	for m := n; m > 0; m /= 10 {
		a = append(a, m%10)
	}
	l := len(a)
	for i := 0; i < l/2; i++ {
		if a[i] != a[l-i-1] {
			return false
		}
	}
	return true
}

func main() {
	for i := start; i > 0; i-- {
		for j := i; j > 800; j-- {
			ij := i * j
			if is_palindrome(ij) {
				fmt.Printf("%d * %d = %d\n", i, j, ij)
				return
			}
		}
	}
}
