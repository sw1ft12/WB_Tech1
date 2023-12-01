package main

import (
	"math/rand"
	"sort"

	"github.com/stretchr/testify/assert"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка
*/

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	l, r := 0, len(arr)-1
	pivot := rand.Intn(len(arr))
	arr[pivot], arr[r] = arr[r], arr[pivot]

	for i := range arr {
		if arr[i] < arr[r] {
			arr[i], arr[l] = arr[l], arr[i]
			l++
		}
	}
	arr[l], arr[r] = arr[r], arr[l]

	QuickSort(arr[:l])
	QuickSort(arr[l+1:])
}

func main() {

	{
		arr := []int{}
		copy := []int{}
		QuickSort(arr)
		sort.Ints(copy)
		assert.EqualValues(nil, arr, copy)
	}

	{
		arr := []int{4343, 2, 23, 8, 773, 0, -200}
		copy := []int{4343, 2, 23, 8, 773, 0, -200}
		QuickSort(arr)
		sort.Ints(copy)
		assert.EqualValues(nil, arr, copy)
	}

	{
		arr := []int{}
		copy := []int{}
		for i := 10000; i >= 1; i-- {
			arr = append(arr, i)
			copy = append(copy, i)
		}
		QuickSort(arr)
		sort.Ints(copy)
		assert.EqualValues(nil, arr, copy)
	}
}
