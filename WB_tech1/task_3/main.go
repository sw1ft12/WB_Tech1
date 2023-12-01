package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
Дана последовательность чисел: 2,4,6,8,10.
Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений
*/

func main() {
	arr := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	wg.Add(len(arr))
	var sum int32
	for _, el := range arr {
		go func(n int) {
			atomic.AddInt32(&sum, int32(n*n))
			wg.Done()
		}(el)
	}
	wg.Wait()
	fmt.Print(sum)
}
