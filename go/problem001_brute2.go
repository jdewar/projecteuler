package main

import (
	"fmt"
)

var generate_max = 1000000
var divisors = []int{3, 5}

func generate(c chan<- int) {
	for i := 1; i < generate_max; i++ {
		c <- i
	}
	close(c)
}

func divisible_by(in <-chan int, allow chan<- int, next chan<- int, d int) {
	for i := range in {
		if i%d == 0 {
			allow <- i
		} else {
			select {
			case next <- i:
			default:
			}
		}
	}
	if next == nil {
		close(allow)
	} else {
		close(next)
	}
}

func main() {
	sum := 0

	// Generate numbers up to the given max
	in := make(chan int)
	go generate(in)

	allow := make(chan int)
	for i, d := range divisors {
		next := make(chan int)
		if i+1 == len(divisors) {
			next = nil
		}
		go divisible_by(in, allow, next, d)
		in = next
	}

	// Pull numbers to sum from channel until closed
	for n := range allow {
		sum += n
	}
	fmt.Println(sum)
}
