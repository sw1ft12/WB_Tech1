package main

import "github.com/stretchr/testify/assert"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

func setBit(n int64, bit int, i uint) int64 {
	if bit == 1 {
		n |= (1 << i)
	} else if bit == 0 {
		n &= ^int64(0) - (1 << i)
	}
	return n
}

func main() {
	assert.Equal(nil, setBit(0, 1, 0), int64(1))
	assert.Equal(nil, setBit(1, 0, 0), int64(0))
	assert.Equal(nil, setBit(16, 0, 4), int64(0))
	assert.Equal(nil, setBit(24, 1, 2), int64(28))
	assert.Equal(nil, setBit(^int64(0), 0, 62), ^int64(0)-(1<<62))
}
