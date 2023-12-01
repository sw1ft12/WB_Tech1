package main

import (
	"time"

	"github.com/stretchr/testify/assert"
)

/*
Реализовать собственную функцию sleep
*/

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	{
		s := time.Now()
		Sleep(5 * time.Second)
		assert.LessOrEqual(nil, 5*time.Second, time.Since(s))
	}

	{
		s := time.Now()
		Sleep(9 * time.Millisecond)
		assert.LessOrEqual(nil, 9*time.Millisecond, time.Since(s))
	}

	{
		s := time.Now()
		Sleep(125 * time.Nanosecond)
		assert.LessOrEqual(nil, 125*time.Nanosecond, time.Since(s))
	}
}
