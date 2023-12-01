package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20
*/

func main() {

	a := new(big.Int)
	a.SetString("10000000000000000000000000000", 10)
	b := new(big.Int)
	b.SetString("99999999999999999999999999999999999999", 10)

	sum := big.NewInt(0)
	sum.Add(a, b)
	fmt.Printf("%v\n", sum) // 100000000009999999999999999999999999999

	sub := big.NewInt(0)
	sub.Sub(b, a)
	fmt.Printf("%v\n", sub) // 99999999989999999999999999999999999999

	prod := big.NewInt(0)
	prod.Mul(a, b)
	fmt.Printf("%v\n", prod) // 999999999999999999999999999999999999990000000000000000000000000000

	div := big.NewInt(0)
	div.Div(b, a)
	fmt.Printf("%v\n", div) // 9999999999
}
