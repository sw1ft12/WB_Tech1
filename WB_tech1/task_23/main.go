package main

import "fmt"

/*
Удалить i-ый элемент из слайса.
*/

func erase(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = erase(arr, 1)
	fmt.Print(arr)
}
