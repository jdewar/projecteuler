package main

// Find largest prime factor of test_num

import (
	"fmt"
	"math"
)

var test_num = 600851475143
var factors = make([]int, 2)
var result = 1

// Maximum of a slice of int arguments
func max(v []int) (m int) {
	m = math.MinInt32
	for _, a := range v {
		if a > m {
			m = a
		}
	}
	return
}

func generate(c chan<- int, max float64) {
	m := int(max)
	for i := 2; i <= m; i++ {
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
	ch := make(chan int)
	go generate(ch, math.Sqrt(float64(test_num)))
	test := test_num
	for result != test_num {
		n := <-ch // n is a prime

		for ; test%n == 0; test = int(test / n) {
			factors = append(factors, n)
			result *= n
		}
		ch_next := make(chan int)
		go divisible_by(ch, ch_next, n)
		ch = ch_next
	}
	fmt.Println(max(factors))
}
