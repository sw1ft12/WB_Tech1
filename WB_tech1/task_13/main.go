package main

import "fmt"

/*
Поменять местами два числа без создания временной переменной
*/

func main() {
	a, b := 5, 10
	a, b = b, a
	fmt.Printf("%d %d", a, b)
}
