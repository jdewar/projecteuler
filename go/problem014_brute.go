package main

import (
	"fmt"
)

func main() {
	longest := 0
	for i := 2; i < 1000000; i++ {
		seq := []int{}
		for n := i; n != 1; {
			if n % 2 == 0 {
				n = n/2
			} else {
				n = 3*n + 1
			}
			seq = append(seq, n)
		}
		if len(seq) > longest {
			longest = len(seq)
			fmt.Printf("%d: %d\n", i, len(seq))
		}
	}
}
