package main

import "fmt"

/*
Разработать конвейер чисел.
Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout
*/

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := make(chan int)
	squares := make(chan int)
	go func() {
		defer close(in)
		for _, el := range x {
			in <- el
		}
	}()

	go func() {
		defer close(squares)
		for el := range in {
			squares <- el * el
		}
	}()

	for el := range squares {
		fmt.Printf("%d ", el)
	}
}
