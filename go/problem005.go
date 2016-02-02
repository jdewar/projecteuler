package main

// Find smallest positive number evenly divisible by 1..20
import (
	"fmt"
)

var up_to = 20

func generate(c chan<- int, max int) {
	m := max
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

func generate_factors(test_num int) map[int]int {
	result := 1
	factors := make(map[int]int)
	ch := make(chan int)
	go generate(ch, test_num)
	test := test_num
	for result != test_num {
		n, ok := <-ch // n is a prime
		if !ok {
			break
		}
		for ; test%n == 0; test = int(test / n) {
			factors[n]++
			result *= n
		}
		ch_next := make(chan int)
		go divisible_by(ch, ch_next, n)
		ch = ch_next
	}
	if result != test_num {
		factors[test_num]++
	}
	return factors
}

func main() {
	tot_factors := make(map[int]int)
	for i := 2; i <= up_to; i++ {
		a := generate_factors(i)
		for k := range a {
			if a[k] > tot_factors[k] {
				tot_factors[k] = a[k]
			}
		}
	}
	full := 1
	for k := range tot_factors {
		for i := 0; i < tot_factors[k]; i++ {
			full *= k
		}
	}
	fmt.Println(full, tot_factors)
}
