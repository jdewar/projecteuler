package main

import (
	"fmt"
	"math/big"
)

var sizeX int64 = 20
var sizeY int64 = 20

func main() {

	halfPerim := sizeX + sizeY

	n := big.NewInt(1)
	d := big.NewInt(1)

	for i := int64(2); i <= halfPerim; i++ {
		n.Mul(n,big.NewInt(i))
	}
	for i := int64(2); i <= halfPerim - sizeX; i++ {
		d.Mul(d,big.NewInt(i))
	}
	for i := int64(2); i <= sizeX; i++ {
		d.Mul(d,big.NewInt(i))
	}
	
	o := big.NewInt(0)
	o.Div(n,d)

	fmt.Println(o)
}
