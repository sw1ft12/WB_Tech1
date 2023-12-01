package main

import "github.com/stretchr/testify/assert"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func BinarySearch(arr []int, val int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := int(uint(left+right) >> 1)
		if arr[mid] < val {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	{
		arr := []int{1, 2, 3, 4, 5}
		assert.Equal(nil, BinarySearch(arr, 2), 1)
		assert.Equal(nil, BinarySearch(arr, 5), 4)
		assert.Equal(nil, BinarySearch(arr, -1), 0)
		assert.Equal(nil, BinarySearch(arr, 100), 5)
	}

	{
		arr := []int{1, 3, 4, 10, 17, 21}
		assert.Equal(nil, BinarySearch(arr, 7), 3)
		assert.Equal(nil, BinarySearch(arr, 19), 5)
		assert.Equal(nil, BinarySearch(arr, 2), 1)
	}

}
