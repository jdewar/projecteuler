package main

import (
	"fmt"
	"math"
	"sort"
)


func main() {


	num := 0
	for e := 1; ; e++ {
		num = num + e
		ds := []int{1, num}
		for d := 2; d <= int(math.Sqrt(float64(num))); d++ {
			if num % d == 0 {
				if num/d == d {
					ds = append(ds, d)
				} else {
					ds = append(ds, d, num/d)
				}
			}
		}
		sort.Ints(ds)
		//fmt.Printf("%d: %v\n", num, ds)
		if len(ds) > 500 {
			fmt.Printf("%d: %d\n",num,len(ds))
			break
		}
	}
}
