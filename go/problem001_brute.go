package main

import (
	"fmt"
)

var generate_max = 1000
var divisors = []int{3, 5}

func generate(c chan<- int) {
	for i := 1; i < generate_max; i++ {
		c <- i
	}
	close(c)
}

func divisible_by_any(in <-chan int, out chan<- int, divisors []int) {
	for i := range in {
		for _, d := range divisors {
			if i%d == 0 {
				out <- i
				break
			}
		}
	}
	close(out)
}

func main() {
	sum := 0

	// Generate numbers up to the given max
	gen := make(chan int)
	go generate(gen)

	allow := make(chan int)
	go divisible_by_any(gen, allow, divisors)

	// Pull numbers to sum from channel until closed
	for n := range allow {
		sum += n
	}
	fmt.Println(sum)
}
