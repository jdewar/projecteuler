package main

// Sum even terms of Fibonacci sequence up to gen_max
import (
	"fmt"
)

var gen_max = 4000000

func generate(c chan<- int) {
	last := 1
	for i := 1; i <= gen_max; i += last {
		c <- i
		last = i
	}
	close(c)
}

func divisible_by(in <-chan int, out chan<- int, divisor int) {
	for i := range in {
		if i%divisor == 0 {
			out <- i
		}
	}
	close(out)
}

func main() {
	sum := 0
	gen := make(chan int)
	go generate(gen)

	allow := make(chan int)
	go divisible_by(gen, allow, 2)
	for n := range allow {
		sum += n
	}
	fmt.Println(sum)
}
