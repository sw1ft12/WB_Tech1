package main

import (
	"sync"
	"sync/atomic"

	"github.com/stretchr/testify/assert"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счётчика
*/

// Разбиваем счётчик на несколько частей так, чтобы разные части попадали в разные кэш-линии
// Тем самым уменьшая коммуникацию между процессорами
type Shard struct {
	value   int32
	padding [64]byte
}

type Counter struct {
	Shards [8]Shard
}

// Увеличить n-ую часть счётчика на 1
func (c *Counter) Increment(n int) {
	atomic.AddInt32(&c.Shards[n].value, 1)
}

// Найти общую сумму по всем частям счётчика
func (c *Counter) GetValue() int32 {
	var sum int32
	for _, i := range c.Shards {
		sum += i.value
	}
	return sum
}

func main() {

	{
		c := Counter{}
		var wg sync.WaitGroup
		wg.Add(8)
		for i := 0; i < 8; i++ {
			go func(n int) {
				defer wg.Done()
				for j := 0; j < 50; j++ {
					c.Increment(n)
				}
			}(i)
		}
		wg.Wait()
		assert.EqualValues(nil, c.GetValue(), 400)
	}

	{
		c := Counter{}
		var wg sync.WaitGroup
		wg.Add(8)
		for i := 0; i < 8; i++ {
			go func(n int) {
				defer wg.Done()
				for j := 0; j < 100000000; j++ {
					c.Increment(n)
				}
			}(i)
		}
		wg.Wait()
		assert.EqualValues(nil, c.GetValue(), 800000000)
	}
}
