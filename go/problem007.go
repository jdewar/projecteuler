package main

import (
	"fmt"
)

func generate(c chan<- int) {
	for i := 2; ; i++ {
		c <- i
	}
}

func divisible_by(in <-chan int, out chan<- int, divisor int) {
	for i := range in {
		if i%divisor != 0 {
			out <- i
		}
	}
	close(out)
}

func main() {
	ch := make(chan int)
	go generate(ch)
	var n int
	for i := 1; i <= 10001; i++ {
		n = <-ch // n is a prime
		ch_next := make(chan int)
		go divisible_by(ch, ch_next, n)
		ch = ch_next
	}
	fmt.Println(n)
}
