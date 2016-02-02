package main

// Shamelessly stolen from the wikipedia page on Project Euler
// http://en.wikipedia.org/wiki/Project_Euler
// NOTE: currently only works for 1 or 2 divisors.

import (
	"fmt"
	"math"
)

var up_to = 1000
var divisors = []int{3, 5}

func floor(n int) int {
	return int(math.Floor(float64(n)))
}

func sum_ki(k int, n int) int {
	m := floor((n - 1) / k)
	return k * m * (m + 1) / 2
}

func inclusion(in <-chan int, answer chan<- int, pipe chan<- int) {
	for d := range in {
		answer <- sum_ki(d, up_to)
		pipe <- d
	}
	close(answer)
	close(pipe)
}

func exclusion(in <-chan int, answer chan<- int) {
	defer func() {
		close(answer)
	}()
	for {
		d1, ok := <-in
		if !ok {
			return
		}
		d2 := <-in
		if !ok {
			return
		}
		d := d1 * d2
		answer <- sum_ki(d, up_to)
	}
}

func main() {
	sum := 0
	gen := make(chan int)
	go func() {
		for _, c := range divisors {
			gen <- c
		}
		close(gen)
	}()

	add := make(chan int)
	included := make(chan int)
	go inclusion(gen, add, included)

	subtract := make(chan int)
	if len(divisors) > 1 {
		go exclusion(included, subtract)
	}
	n := 2
	for {
		select {
		case p, ok := <-add:
			if !ok {
				n--
			} else {
				sum += p
			}
		case m, ok := <-subtract:
			if !ok {
				n--
			} else {
				sum -= m
			}
		}
		if n == 0 {
			break
		}
	}
	fmt.Println(sum)
}
