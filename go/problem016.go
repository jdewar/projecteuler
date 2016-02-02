package main

import (
	"fmt"
	"math/big"
)

func main() {

	o := big.NewInt(1)

	for i := 1; i <= 1000; i++ {
		o.Mul(o, big.NewInt(int64(2)))
	}
	s := big.NewInt(0)
	ten := big.NewInt(10)
	d := big.NewInt(0)

	for m := o; m.Cmp(big.NewInt(0)) != 0; {
		d.Mod(m, ten)
                s.Add(s, d)
		m.Div(m, ten)
	}
	fmt.Println(s)
}
