// 142913828922
// 30+ min of runtime!  concurrent prime sieve is slick but crap
package main

import (
	"fmt"
	"runtime"
)

func generate(c chan<- int) {
	for i := 2; i < 2000000; i++ {
		c <- i
	}
	close(c)
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
	runtime.GOMAXPROCS(3)
	ch := make(chan int, 10)
	go generate(ch)
	sum := 0
	count := 0
	for {
		n, ok := <-ch // n is a prime
		if !ok {
			break
		}
		ch_next := make(chan int, 10)
		go divisible_by(ch, ch_next, n)
		ch = ch_next
		sum += n
		count++
	}
	fmt.Println(sum, count)
}
