package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	for i := 2; i < 1000000; i++ {
		count := 0
		for n := i; n != 1; {
			prev_count, ok := m[n]
			if (!ok) {
				if n % 2 == 0 {
					n = n/2
				} else {
					n = 3*n + 1
				}
				count++
			} else {
				count += prev_count
				n = 1
			}
			if n == 1 {
				m[i] = count
			}
		}
	}
	longest := 0
	for k, v := range m {
		if v > longest {
			longest = v
			fmt.Printf("%d: %d\n", k, v)
		}
	}
}
